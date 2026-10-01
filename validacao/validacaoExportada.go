package validacao

import "github.com/gin-gonic/gin"

func ValidarProduto(nome, categoria, valorStr string) (float64, error) {
	return validarProduto(nome, categoria, valorStr)
}

func ValidarImagem(c *gin.Context) (string, string, error) {
	return validarImagem(c)
}
