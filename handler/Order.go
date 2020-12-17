package handler

import (
	"strconv"
	"go-practice/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler interface {
	OrderProduct(*gin.Context)
}

type orderHandler struct{
	repo repository.OrderRepository
}

func NewOrderHandler() OrderHandler{
	return &orderHandler{
		repo: repository.NewOrderRepository
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context){
	productIdString := ctx.Param("productId")
	if productId, err:= strconv.Atoi(productIdString); err!=nil{
		ctx.JSON(http.StatusBadRequest, 
			gin.H{
				"error": err.Error()
			})
	}else {
		quantityString := ctx.Param("quantity")
		if quantityId, err:= strconv.Atoi(quantityString); err!=nil{
			ctx.JSON(http.StatusBadRequest, 
				gin.H{
					"error": err.Error()
				})
			}else{
				userId := ctx.GetFloat64("userId")
				if err:= h.repo.OrderProduct(int(userId), productId, quantityId ); err!=nil{
					ctx.JSON(http.StatusBadRequest, 
						gin.H{
							"error": err.Error()
						})
				}else{
					ctx.String(http.StatusOK, 
						"Product Ordered Successfully"
				}
			}
	}
}