package router

import (
	"go-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/generate-cnpj", handler.GenerateCNPJ)
		v1.GET("/generate-cpf", handler.GenerateCPF)
	}

	return r
}
