package handler

import (
	"api_produtos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarTodosClientes(c *gin.Context){

	c.JSON(http.StatusOK, service.ListarTodosClientes())
}

func CadastrarCliente(c *gin.Context){
	novoCliente, status, err := service.CadastrarCliente(c)

	if err != nil {
		c.JSON(status, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(status, gin.H{
		"mensagem": "Cliente cadastrado com sucesso!",
		"cliente":  novoCliente,
	})
}
func AlterarCliente(c *gin.Context) {
	codigo := c.Param("codigo")
	clienteAtualizado, status, err := service.AlterarCliente(codigo, c)
	if err != nil {
		c.JSON(status, gin.H{"mensagem": err.Error()})
		return
	}
	c.JSON(status, gin.H{
		"mensagem": "Cliente atualizado com sucesso!",
		"cliente":  clienteAtualizado,
	})
}
func RemoverCliente(c *gin.Context) {
	codigo := c.Param("codigo")
	status, err := service.RemoverCliente(codigo)
	if err != nil {
		c.JSON(status, gin.H{"mensagem": err.Error()})
		return
	}
	c.JSON(status, gin.H{"mensagem": "Cliente removido com sucesso!"})
}