package users

import (
	"fmt"
	"github.com/obadalov/bookstore-users-api/utils/errors"
)

var userDB = make(map[int64]*User)

func (u *User) Get() *errors.RestErr {
	result := userDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}
	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return nil
}

func (u *User) Save() *errors.RestErr {
	current := userDB[u.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.Id))
	}
	userDB[u.Id] = u
	return nil
}
