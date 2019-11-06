package role

import (
		"fmt"
		"gopkg.in/go-playground/validator.v9"
)

type Role struct {
	Id     int `validate:"required,gte=1"`
	Name   string `validate:"required"`
	Parent int `validate:"gte=0"`
}

var roleList []Role

//roleMap stores roleID and its subordinates map[0:[1] 1:[2] 2:[3] 3:[4 5]]
var roleMap map[int][]int

var aclTree map[int][]int

//keep the role json in memory (later on can we move to redis or some storage for scalibilty)
func SetRoles(rolesObjList []Role) error {
	roleList = make([]Role, 0)
	roleList = append(roleList, rolesObjList...)
	roleMap = make(map[int][]int, 0)
	if err := hasDuplicate(roleList); err != nil {
		return err
	}
    v := validator.New()
	for _, role := range rolesObjList {
		err := v.Struct(role)
        
        if err != nil {
		   return err
		}
		// make sure parent exists except 0
		if role.Parent != 0 {
			if _, err := getRoleInfo(role.Parent); err != nil {
				return err
			}
		}
		roleMap[role.Parent] = append(roleMap[role.Parent], role.Id)
	}

	if err := buildAclTree(); err != nil {
		return err
	}
	return nil
}

//get role info for given id
func getRoleInfo(roleID int) (Role, error) {
	if roleID > len(roleList) {
		return Role{}, fmt.Errorf("not an valid id %v", roleID)
	}
	return roleList[roleID-1], nil
}


func hasDuplicate(rolesObjList []Role) error {
    var list []int
    for _, role := range rolesObjList {
        list = append(list, role.Id)
    }

    unique := removeDuplicates(list)
    if(len(unique) != len(list)) {
    	return fmt.Errorf("duplicate found")
    }

    return nil
}

func removeDuplicates(s []int) []int {
      m := make(map[int]bool)
      for _, item := range s {
              if _, ok := m[item]; ok {
                      // duplicate item
                      fmt.Println(item, "is a duplicate")
              } else {
                      m[item] = true
              }
      }

      var result []int
      for item, _ := range m {
              result = append(result, item)
      }
      return result
}

// create relationship of parent id and role id map[1:[2 3 4 5] 2:[3 4 5] 3:[4 5] 4:[] 5:[]]
func buildAclTree() error {
	aclTree = make(map[int][]int, 0)
	for i := 1; i <= len(roleList); i++ {
		visited := map[int]bool{}
		results := []int{}
		queue := []int{}
		queue = append(queue, i)
		for len(queue) != 0 {
			head := queue[0]
			if visited[head] {
				return fmt.Errorf("please provide valid role json")
			}
			results = append(results, roleMap[head]...)
			queue = append(queue, roleMap[head]...)
			visited[head] = true
			queue = queue[1:]
		}
		aclTree[i] = results
	}
	return nil
}

func GetAcl(parent int) []int {
	return aclTree[parent]
}
