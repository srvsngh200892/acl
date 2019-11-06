package user

import (
	"fmt"
	"github.com/srvsngh200892/acl/src/role"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Id  int `validate:"required,gte=1"`
	Name string `validate:"required"`
	Role int `validate:"required,gte=1"`
}

var userList []User
var roleUserMap map[int][]int

//keep the user details json in memory (later on can we move to redis or some storage for scalibilty)
func SetUsers(userObjList []User) error {
	userList = make([]User, 0)
	userList = append(userList, userObjList...)
	roleUserMap = make(map[int][]int, 0)
	if err := hasDuplicate(userList); err != nil {
		fmt.Println(err)
		return err
	}
	v := validator.New()
	for _, user := range userObjList {
		err := v.Struct(user)

        if err != nil {
		   return err
		}
		roleUserMap[user.Role] = append(roleUserMap[user.Role], user.Id)
	}

	return nil
}

//get user info by passing user id
func getUserInfo(userID int) (User, error) {
	if userID > len(userList) {
		return User{}, fmt.Errorf("not an valid id %v", userID)
	}
	return userList[userID-1], nil
}

//get list of users by particular role id
func getUserInfosByRoleID(roleID int) ([]User, error) {
	userIDs := roleUserMap[roleID]
	userList := []User{}
	for _, userID := range userIDs {
		user, err := getUserInfo(userID)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}

//get list of subordinate for given user by id
func GetSubOrdinates(userID int) ([]User, error) {
	user, err := getUserInfo(userID)
	if err != nil {
		return nil, err
	}

	subordinateList := []User{}
	subOrdinatesIDs := role.GetAcl(user.Role)
	for _, roleID := range subOrdinatesIDs {
		subordinates, err := getUserInfosByRoleID(roleID)
		if err != nil {
			return nil, err
		}
		subordinateList = append(subordinateList, subordinates...)
	}
	return subordinateList, nil
}

func hasDuplicate(userObjList []User) error {
    var list []int
    for _, user := range userObjList {
        list = append(list, user.Id)
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