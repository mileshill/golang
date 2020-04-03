package domain

import (
	"fmt"
	"github.com/miles1618/udemy/rest/mvc/utils"
	"net/http"
)

var(
	users = map[int64]*User{
		123: &User{
			Id: 123,
			FirstName: "Miles",
			LastName: "Hill",
			Email: "miles1618@gmail.com",
			},
	}
)



func GetUser(userId int64) (*User, *utils.ApplicationError){
	// This is where the database call goes
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code: http.StatusText(http.StatusNotFound),
		}
	}
	return user, nil
}

