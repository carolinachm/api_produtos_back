package router

import (
	"api_produtos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProdutoRotas — registra todas as rotas de produto
func ProdutoRotas(router *gin.Engine) {
	router.GET("/produtos", listarTodosProdutos)
	router.POST("/produtos", cadastrarProduto)
	router.PUT("/produtos/:codigo", alterarProduto)
	router.DELETE("/produtos/:codigo", removerProduto)
}

func listarTodosProdutos(c *gin.Context) {
	c.JSON(http.StatusOK, service.ListarTodosProdutos())
}

func cadastrarProduto(c *gin.Context) {
	novoProduto, status, err := service.CadastrarProduto(c)
	if err != nil {
		c.JSON(status, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(status, gin.H{
		"mensagem": "Produto cadastrado com sucesso!",
		"produto":  novoProduto,
	})
}

func alterarProduto(c *gin.Context) {
	codigo := c.Param("codigo")
	produtoAtualizado, status, err := service.AlterarProduto(codigo, c)
	if err != nil {
		c.JSON(status, gin.H{"mensagem": err.Error()})
		return
	}
	c.JSON(status, gin.H{
		"mensagem": "Produto atualizado com sucesso!",
		"produto":  produtoAtualizado,
	})
}

func removerProduto(c *gin.Context) {
	codigo := c.Param("codigo")
	status, err := service.RemoverProduto(codigo)
	if err != nil {
		c.JSON(status, gin.H{"mensagem": err.Error()})
		return
	}
	c.JSON(status, gin.H{"mensagem": "Produto removido com sucesso!"})
}