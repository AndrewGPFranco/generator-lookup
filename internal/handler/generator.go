package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/internal/service"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func GeneratePassword(c *gin.Context) {
	lengthParam := c.Param("length")

	length, err := strconv.Atoi(lengthParam)
	if err != nil || length <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid length",
		})
		return
	}

	password := service.GeneratorPassword(length)
	c.JSON(http.StatusOK, gin.H{
		"password": password,
	})
}

func GenerateCNPJ(c *gin.Context) {
	cnpj := service.GeneratorValidCNPJ()
	c.JSON(http.StatusOK, gin.H{
		"lookup":    cnpj,
		"formatted": service.FormatCNPJ(cnpj),
	})
}

func GenerateCPF(c *gin.Context) {
	cpf := service.GenerateValidCPF()
	c.JSON(http.StatusOK, gin.H{
		"lookup":    cpf,
		"formatted": service.FormatCPF(cpf),
	})
}
