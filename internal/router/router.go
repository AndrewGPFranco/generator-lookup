package router

import (
	"go-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/generate-cnpj", handler.GenerateCNPJ)
		v1.POST("/generate-cpf", handler.GenerateCPF)
	}

	return r
}
