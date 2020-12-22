package handler

import (
	"fmt"
	"go-practice/api/repository"
	"go-practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	CreateUser(*gin.Context)
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetProductsOrdered(*gin.Context)
	GetAllUser(*gin.Context)
	SignIn(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repository.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePassword := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	fmt.Println(hPass)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
func (h *userHandler) GetAllUser(ctx *gin.Context) {
	fmt.Println(ctx.Get("userID"))
	user, err := h.repo.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}
func (h *userHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword(&user.Password)
	user, err := h.repo.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)

}
func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.repo.GetUserByID(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}
func (h *userHandler) SignIn(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbUser, err := h.repo.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No Such User Found"})
		return

	}
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignIN", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})
	return

}
func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user.ID = uint(intID)
	user, err = h.repo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("user")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	user, err := h.repo.DeleteUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}
func (h *userHandler) GetProductsOrdered(ctx *gin.Context) {

	userStr := ctx.Param("user")
	userID, _ := strconv.Atoi(userStr)
	if products, err := h.repo.GetProductOrdered(userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, products)
	}
}
