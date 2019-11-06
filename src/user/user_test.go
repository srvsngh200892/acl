package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetInvalidUser(t *testing.T) {
	err := SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   1,
			Name: "Emily Employee",
			Role: 4,
		},
	})
	assert.Contains(t, err.Error(), "duplicate found")

	err = SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   2,
			Role: 4,
		},
	})
	assert.Contains(t, err.Error(), "validation")

	err = SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   0,
			Name: "Emily Employee",
			Role: 4,
		},
	})
	assert.Contains(t, err.Error(), "validation")

	err = SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   1,
			Name: "Emily Employee",
			Role: 0,
		},
	})

	assert.Contains(t, err.Error(), "duplicate found")

	err = SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   3,
			Name: "Emily Employee",
			Role: 0,
		},
	})

	assert.Contains(t, err.Error(), "Error:Field validation")
}

func TestSetUsers(t *testing.T) {
	err := SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   3,
			Name: "Emily Employee",
			Role: 2,
		},
	})
	if(err==nil) {
	 	_, err = getUserInfo(3)
	}
	assert.Equal(t, "not an valid id 3", err.Error())
}


func TestGetUserInfo(t *testing.T) {
	err := SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   2,
			Name: "Emily Employee",
			Role: 2,
		},
	})

	if(err==nil) {
		data, _ := getUserInfo(2)
		assert.Equal(t, userList[1], data)
	}
}

func TestHasDuplicate(t *testing.T) {
	users := []User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   1,
			Name: "Emily Employee",
			Role: 2,
		},
	}
    err := hasDuplicate(users)
    assert.Contains(t, err.Error(), "duplicate found")
}

func TestGetUserInfosByRoleID(t *testing.T) {
	err := SetUsers([]User{
		User{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			Id:   2,
			Name: "Emily Employee",
			Role: 2,
		},
	})
    if(err==nil) {
		data, _ := getUserInfosByRoleID(1)
		assert.Equal(t, 1, len(data))
	}
}