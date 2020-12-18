package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileUploadHandler interface {
	SingleFile(*gin.Context)
	MultipleFile(*gin.Context)
}

func SingleFile(ctx *gin.Context) {
	file, err := ctx.FormFile("profile")
	if err != nil {
		log.Fatal(err)
	}
	err = ctx.SaveUploadedFile(file, "files/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func MultipleFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["profile"]
	for _, file := range files {
		fmt.Println(file.Filename)
		err := ctx.SaveUploadedFile(file, "files/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%d' uploaded!", len(files)))
}
