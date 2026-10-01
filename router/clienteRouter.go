package router

import (
	"api_produtos/handler"

	"github.com/gin-gonic/gin"
)

// ClienteRouter — registra todas as rotas de cliente
func ClienteRouter(router *gin.Engine) {
	router.GET("/clientes", handler.ListarTodosClientes)
	router.POST("/clientes", handler.CadastrarCliente)
	router.PUT("/clientes/:codigo", handler.AlterarCliente)
	router.DELETE("/clientes/:codigo", handler.RemoverCliente)
}