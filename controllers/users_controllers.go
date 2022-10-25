package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/bookstore-users-api/domain/users"
	"github.com/samarec1812/bookstore-users-api/services"
	"github.com/samarec1812/bookstore-users-api/utils/errors"
	"net/http"
	"strconv"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
