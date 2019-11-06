package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetInvalidRoles(t *testing.T) {
	err := SetRoles([]Role{
		Role{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			Id:     2,
			Name:   "Location Manager",
			Parent: 2,
		},
	})
	assert.Contains(t, err.Error(), "please provide valid role json")

	err = SetRoles([]Role{
		Role{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			Id:     1,
			Name:   "Location Manager",
			Parent: 2,
		},
	})
	assert.Contains(t, err.Error(), "duplicate found")

	err = SetRoles([]Role{
		Role{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			Id:     2,
			Name:   "Location Manager",
			Parent: -1,
		},
	})
	assert.Contains(t, err.Error(), "validation")
}

func TestSetRoles(t *testing.T) {
	err := SetRoles([]Role{
		Role{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			Id:     3,
			Name:   "Location Manager",
			Parent: 3,
		},
	})
	assert.Equal(t, "not an valid id 3", err.Error())
}


func TestGetRoleInfo(t *testing.T) {
	err := SetRoles([]Role{
		Role{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			Id:     2,
			Name:   "Location Manager",
			Parent: 3,
		},
	})

	if(err==nil) {
		data, _ := getRoleInfo(2)
		assert.Equal(t, roleList[1], data)
	}
}

func TestHasDuplicate(t *testing.T) {
	roles := []Role{
		Role{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			Id:     1,
			Name:   "Location Manager",
			Parent: 3,
		},
	}
    err := hasDuplicate(roles)
    assert.Contains(t, err.Error(), "duplicate found")
}

