package users

import (
	"fmt"
	"github.com/obadalov/bookstore-users-api/datasources/mysql/users_db"
	"github.com/obadalov/bookstore-users-api/utils/date"
	"github.com/obadalov/bookstore-users-api/utils/errors"
)

var userDB = make(map[int64]*User)

func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	u.DateCreated = date.GetNowAsString()
	userDB[u.Id] = u
	return nil
}
