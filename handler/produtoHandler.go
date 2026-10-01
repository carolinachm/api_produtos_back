package handler
import (
	"api_produtos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarTodosProdutos(c *gin.Context) {
	c.JSON(http.StatusOK, service.ListarTodosProdutos())
}

func CadastrarProduto(c *gin.Context) {
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

func AlterarProduto(c *gin.Context) {
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

func RemoverProduto(c *gin.Context) {
	codigo := c.Param("codigo")
	status, err := service.RemoverProduto(codigo)
	if err != nil {
		c.JSON(status, gin.H{"mensagem": err.Error()})
		return
	}
	c.JSON(status, gin.H{"mensagem": "Produto removido com sucesso!"})
}
