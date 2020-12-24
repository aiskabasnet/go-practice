package controllers

import (
	"fmt"
	fbservice "go-practice/api/service"
	service "go-practice/api/service/user"
	"go-practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type firebaseUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
	UID      string `json:"uid"`
	SNS      bool   `json:"sns"`
}
type userController struct {
	fbService   fbservice.FirebaseService
	userService service.UserService
}

func NewUserController(s service.UserService, f fbservice.FirebaseService) UserInterface {
	return &userController{
		fbService:   f,
		userService: s,
	}
}

//GetUsers ... Get all users
func (u *userController) GetUsers(c *gin.Context) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in fetching data"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

//CreateUser ... Create User
func (u *userController) CreateUser(c *gin.Context) {
	var user models.User
	var firebaseUser firebaseUser
	if err := c.ShouldBindJSON(&firebaseUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var fbID string
	var err error
	if !firebaseUser.SNS {
		fbID, err = u.fbService.CreateUser(firebaseUser.Email, firebaseUser.Password)
		//if already registered
		if err != nil {
			u.fbService.DeleteUser(fbID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "This User is already Registered"})
			return
		}
	} else {
		fbID, err = u.fbService.UpdateUser(fbUser.UID, true)
	}
	randUserName := utils.GenerateRandomInvitationCode(12)

	user.UserName = randUserName
	user.Email = fbUser.Email
	user.UserType = fbUser.UserType
	user.ID = fbID
	if fbUser.SNS {
		user.ID = fbUser.UID
	}

	_, err = u.userService.Save(user, fbUser.SNS)

	err = models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
