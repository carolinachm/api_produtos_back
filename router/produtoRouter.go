package router

import (
	"api_produtos/handler"

	"github.com/gin-gonic/gin"
)

// ProdutoRotas — registra todas as rotas de produto
func ProdutoRouter(router *gin.Engine) {
	router.GET("/produtos", handler.ListarTodosProdutos)
	router.POST("/produtos", handler.CadastrarProduto)
	router.PUT("/produtos/:codigo", handler.AlterarProduto)
	router.DELETE("/produtos/:codigo", handler.RemoverProduto)
}