package handler

import (
	"go-practice/api/repository"
	"go-practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProducts(*gin.Context)
	GetProductById(*gin.Context)
	CreateProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

func (h *productHandler) GetProducts(ctx *gin.Context) {
	product, err := h.repo.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}

func (h *productHandler) GetProductById(ctx *gin.Context) {
	prodStr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := h.repo.GetProductById(prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}

func (h *productHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := h.repo.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}
func (h *productHandler) UpdateProduct(ctx *gin.Context) {

	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prodStr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(prodID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}
func (h *productHandler) DeleteProduct(ctx *gin.Context) {

	var product models.Product
	prodStr := ctx.Param("product")
	prodID, _ := strconv.Atoi(prodStr)
	product.ID = uint(prodID)
	product, err := h.repo.DeleteProduct(prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}
