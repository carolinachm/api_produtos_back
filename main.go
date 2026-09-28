package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"uuid"
)

// struct produtos
type Produto struct {
	Codigo    string  `json:"codigo" form:"codigo"`
	Nome      string  `json:"nome" form:"nome"`
	Categoria string  `json:"categoria" form:"categoria"`
	Valor     float64 `json:"valor" form:"valor"`
	Imagem    string  `json:"imagem" form:"imagem"`
}

// Lista (slice)
var produtos = []Produto{}

// função principal
func main() {
	// roteador
	router := gin.Default()

	// rotas
	router.GET("/", helloWorld)
	router.GET("/produtos", listarTodosProdutos)
	router.POST("/produtos", cadastrarProdutos)
	router.PUT("/produtos/:codigo", alterarProduto)
	router.DELETE("/produtos/:codigo", removerProduto)

	// configurar o servidor
	router.Run(":8080")
}

// Função para retornar um Hello World
func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mensagem": "Hello World"})
}

// Função para listar todos os produtos
func listarTodosProdutos(c *gin.Context) {
	c.JSON(http.StatusOK, produtos)
}

// Função para cadastrar produtos
func cadastrarProdutos(c *gin.Context) {
	// obter nome, categoria, valor e imagem
	nome := c.PostForm("nome")
	categoria := c.PostForm("categoria")
	valor := c.PostForm("valor")
	imagem, _ := c.FormFile("imagem")

	// converter o valor
	valorConvertido, _ := strconv.ParseFloat(valor, 64)

	// gerar o código
	codigo := uuid.New().String()

	// extrair a extensão da imagem
	extensao := filepath.Ext(imagem.Filename)

	// gerar o novo nome da imagem
	nomeImagem := codigo + extensao

	// informar o local onde será armazenada a imagem
	caminho := filepath.Join("uploads", nomeImagem)

	// realizar o upload
	c.SaveUploadedFile(imagem, caminho)

	// criar struct produto
	novoProduto := Produto{
		Codigo:    codigo,
		Nome:      nome,
		Categoria: categoria,
		Valor:     valorConvertido,
		Imagem:    "uploads/" + nomeImagem,
	}

	// adicionar o novo Produto
	produtos = append(produtos, novoProduto)

	// retorno do endpoint ao cadastrar produto
	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Produto cadastrado com sucesso",
		"produto":  novoProduto,
	})
}

// função para alterar cadastro de produto
func alterarProduto(c *gin.Context) {
}

// função para remover produto
func removerProduto(c *gin.Context) {
	// extrair o código do produto via URL
	codigo := c.Param("codigo")

	// laço de repetição para localizar o produto na lista e remover
	for indice, produto := range produtos {
		// condicional — comparação correta
		if produto.Codigo == codigo {
			produtos = append(produtos[:indice], produtos[indice+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensagem": "Produto removido"})
			return
		}
	}

	// retorno caso não seja encontrado
	c.JSON(http.StatusNotFound, gin.H{"mensagem": "Produto não encontrado"})
}