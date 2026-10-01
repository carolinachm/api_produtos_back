package main

import (
	"api_produtos/router"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	
	// Limpar e recriar pasta uploads
	os.RemoveAll("./uploads")
	os.MkdirAll("./uploads", 0755)

	// Roteador
	r := gin.Default()

	// Arquivos estáticos
	r.Static("/uploads", "./uploads")

	// CORS
	r.Use(corsMiddleware())

	// Rota raiz
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Hello World"})
	})

	// Registrar rotas de produto
	router.ProdutoRotas(r)
	router.ClienteRouter(r)

	// Iniciar servidor
	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}