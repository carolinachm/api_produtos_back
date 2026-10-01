package service

import (
	"api_produtos/erro"
	"api_produtos/models"
	"api_produtos/validacao"
	"net/http"
	"path/filepath"
	"strconv"
	"uuid"

	"github.com/gin-gonic/gin"
)

// ========== BANCO EM MEMÓRIA ==========
var produtos = []models.ProdutoModel{}

// ========== SERVIÇOS ==========
func ListarTodosProdutos() []models.ProdutoModel {
	return produtos
}

func CadastrarProduto(c *gin.Context) (models.ProdutoModel, int, error) {
	nome := c.PostForm("nome")
	categoria := c.PostForm("categoria")
	valorStr := c.PostForm("valor")

	valorConvertido, err := validacao.ValidarProduto(nome, categoria, valorStr)
	if err != nil {
		return models.ProdutoModel{}, http.StatusBadRequest, err
	}

	nomeImagem, _, err := validacao.ValidarImagem(c)
	if err != nil {
		return models.ProdutoModel{}, http.StatusBadRequest, err
	}

	codigo := uuid.New().String()
	novoProduto := models.ProdutoModel{
		Codigo:    codigo,
		Nome:      nome,
		Categoria: categoria,
		Valor:     valorConvertido,
		Imagem:    "uploads/" + nomeImagem,
	}

	produtos = append(produtos, novoProduto)
	return novoProduto, http.StatusCreated, nil
}

func AlterarProduto(codigo string, c *gin.Context) (*models.ProdutoModel, int, error) {
	for indice, produto := range produtos {
		if produto.Codigo == codigo {
			nome := c.PostForm("nome")
			categoria := c.PostForm("categoria")
			valorStr := c.PostForm("valor")

			if nome != "" {
				if len(nome) < 2 {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "O 'nome' deve ter pelo menos 2 caracteres."}
				}
				produtos[indice].Nome = nome
			}

			if categoria != "" {
				if len(categoria) < 2 {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "A 'categoria' deve ter pelo menos 2 caracteres."}
				}
				produtos[indice].Categoria = categoria
			}

			if valorStr != "" {
				valorConvertido, err := strconv.ParseFloat(valorStr, 64)
				if err != nil {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "Valor inválido. Informe um número válido."}
				}
				if valorConvertido <= 0 {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "O valor deve ser maior que zero."}
				}
				produtos[indice].Valor = valorConvertido
			}

			imagem, err := c.FormFile("imagem")
			if err == nil {
				extensao := filepath.Ext(imagem.Filename)
				extensoesPermitidas := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
				if !extensoesPermitidas[extensao] {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "Formato de imagem inválido. Use: .jpg, .jpeg, .png ou .gif."}
				}

				tamanhoMaximo := int64(5 * 1024 * 1024)
				if imagem.Size > tamanhoMaximo {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "A imagem é muito grande. Tamanho máximo: 5MB."}
				}

				nomeImagem := codigo + extensao
				caminho := filepath.Join("uploads", nomeImagem)
				if err := c.SaveUploadedFile(imagem, caminho); err != nil {
					return nil, http.StatusInternalServerError, &erro.ErroValidacao{Mensagem: "Não foi possível salvar a nova imagem."}
				}
				produtos[indice].Imagem = "uploads/" + nomeImagem
			}

			return &produtos[indice], http.StatusOK, nil
		}
	}
	return nil, http.StatusNotFound, &erro.ErroValidacao{Mensagem: "Produto não encontrado."}
}

func RemoverProduto(codigo string) (int, error) {
	for indice, produto := range produtos {
		if produto.Codigo == codigo {
			produtos = append(produtos[:indice], produtos[indice+1:]...)
			return http.StatusOK, nil
		}
	}
	return http.StatusNotFound, &erro.ErroValidacao{Mensagem: "Produto não encontrado."}
}