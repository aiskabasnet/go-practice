package controllers

import (
	"go-practice/api/responses"
	fbservice "go-practice/api/service"
	service "go-practice/api/service/user"
	"go-practice/models"
	"go-practice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
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
	user.Email = firebaseUser.Email
	user.UserType = firebaseUser.UserType
	user.ID = fbID
	if firebaseUser.SNS {
		user.ID = firebaseUser.UID
	}

	_, err = u.userService.AddUser(user)
	if err != nil {
		if err.Error() == "This is Login" {
			responses.SuccessJSON(c, http.StatusOK, "This is Login")
			return

		}
		u.fbService.DeleteUser(user.ID)

		responses.ErrorJSON(c, http.StatusBadRequest, "This User is already Registered")
		return
	}
	//Add auth claims to the user

	claims := gin.H{"role": user.UserType}
	err = u.fbService.SetClaim(fbID, claims)

	responses.SuccessJSON(c, http.StatusOK, "SuccessFully Added User")

}

//GetUserByID ... Get the user by id
func (u *userController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := u.userService.GetUserByID(id)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadGateway, err.Error())
		return
	} else {
		responses.SuccessJSON(c, http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func (u *userController) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	// if not user
	err := u.userService.GetUserByID(id)
	if err != nil {
		responses.ErrorJSON(e, http.StatusNotFound, err.Error())
	}
	user.ID = id
	// update claim
	if user.UserType != "" {
		//Add auth claims to the user
		claims := gin.H{"role": user.UserType}
		u.fbService.SetClaim(id, claims)

	}
	updatedUser, err := u.userService.UpdateUser(user)

	if err != nil {
		responses.ErrorJSON(e, http.StatusBadGateway, err.Error())
	}
	responses.SuccessJSON(c, http.StatusOK, updatedUser)
}
