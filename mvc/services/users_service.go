package services

import (
	"github.com/miles1618/udemy/rest/mvc/domain"
	"github.com/miles1618/udemy/rest/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
