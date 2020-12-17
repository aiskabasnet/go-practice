package handler

import (
	"go-practice/Models"
	"strconv"
	"go-practice/repository"
	"net/http"
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

func NewProductHandler() ProductHandler{
	return &productHandler{
		repo: repository.NewProductRepository()
	}
}

func (h *productHandler)GetProducts(ctx *gin.Context){
	products, err := h.repo.GetAllProducts()
	if(err!=nil){
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error", err.Error()})
		return;
	}
	ctx.JSON(http.StatusOK, products)

}
func (h *productHandler)GetProductById(ctx *gin.Context){
	idString := ctx.Param("id")
	id,err := strconv.Atoi(idString)
	if(err!=nil){
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	product,err := h.repo.GetProductById(id)
	if(err!=nil){
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error", err.Error()})
		return;
	}
	ctx.JSON(http.StatusOK, product)

}
func (h *productHandler)CreateProduct(ctx *gin.Context){
	var product Models.Product
	
	product := h.repo.CreateProduct()
}
func (h *productHandler)UpdateProduct(ctx *gin.Context){

}
func (h *productHandler)DeleteProduct(ctx *gin.Context){

}