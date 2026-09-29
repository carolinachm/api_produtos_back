package service

import (
	"api_produtos/models"
	"net/http"
	"path/filepath"
	"strconv"

	"uuid"

	"github.com/gin-gonic/gin"
)

// ========== BANCO EM MEMÓRIA ==========
var produtos = []models.Produto{}

// ========== VALIDAÇÕES ==========
func validarDados(nome, categoria, valorStr string) (float64, error) {
	if nome == "" {
		return 0, &ErroValidacao{Campo: "nome", Mensagem: "O campo 'nome' é obrigatório."}
	}
	if len(nome) < 2 {
		return 0, &ErroValidacao{Campo: "nome", Mensagem: "O 'nome' deve ter pelo menos 2 caracteres."}
	}
	if categoria == "" {
		return 0, &ErroValidacao{Campo: "categoria", Mensagem: "O campo 'categoria' é obrigatório."}
	}
	if len(categoria) < 2 {
		return 0, &ErroValidacao{Campo: "categoria", Mensagem: "A 'categoria' deve ter pelo menos 2 caracteres."}
	}
	if valorStr == "" {
		return 0, &ErroValidacao{Campo: "valor", Mensagem: "O campo 'valor' é obrigatório."}
	}
	valorConvertido, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		return 0, &ErroValidacao{Campo: "valor", Mensagem: "Valor inválido. Informe um número válido (ex: 5500 ou 199.90)."}
	}
	if valorConvertido <= 0 {
		return 0, &ErroValidacao{Campo: "valor", Mensagem: "O valor deve ser maior que zero."}
	}
	return valorConvertido, nil
}

func validarImagem(c *gin.Context) (string, string, error) {
	imagem, err := c.FormFile("imagem")
	if err != nil {
		return "", "", &ErroValidacao{Campo: "imagem", Mensagem: "A imagem é obrigatória. Selecione um arquivo de imagem."}
	}

	extensao := filepath.Ext(imagem.Filename)
	extensoesPermitidas := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !extensoesPermitidas[extensao] {
		return "", "", &ErroValidacao{Campo: "imagem", Mensagem: "Formato de imagem inválido. Use: .jpg, .jpeg, .png ou .gif."}
	}

	tamanhoMaximo := int64(5 * 1024 * 1024)
	if imagem.Size > tamanhoMaximo {
		return "", "", &ErroValidacao{Campo: "imagem", Mensagem: "A imagem é muito grande. Tamanho máximo: 5MB."}
	}

	codigo := uuid.New().String()
	nomeImagem := codigo + extensao
	caminho := filepath.Join("uploads", nomeImagem)

	if err := c.SaveUploadedFile(imagem, caminho); err != nil {
		return "", "", &ErroValidacao{Campo: "imagem", Mensagem: "Não foi possível salvar a imagem. Verifique a pasta 'uploads'."}
	}

	return nomeImagem, caminho, nil
}

// ========== SERVIÇOS ==========
func ListarTodosProdutos() []models.Produto {
	return produtos
}

func CadastrarProduto(c *gin.Context) (models.Produto, int, error) {
	nome := c.PostForm("nome")
	categoria := c.PostForm("categoria")
	valorStr := c.PostForm("valor")

	valorConvertido, err := validarDados(nome, categoria, valorStr)
	if err != nil {
		return models.Produto{}, http.StatusBadRequest, err
	}

	nomeImagem, _, err := validarImagem(c)
	if err != nil {
		return models.Produto{}, http.StatusBadRequest, err
	}

	codigo := uuid.New().String()
	novoProduto := models.Produto{
		Codigo:    codigo,
		Nome:      nome,
		Categoria: categoria,
		Valor:     valorConvertido,
		Imagem:    "uploads/" + nomeImagem,
	}

	produtos = append(produtos, novoProduto)
	return novoProduto, http.StatusCreated, nil
}

func AlterarProduto(codigo string, c *gin.Context) (*models.Produto, int, error) {
	for indice, produto := range produtos {
		if produto.Codigo == codigo {
			nome := c.PostForm("nome")
			categoria := c.PostForm("categoria")
			valorStr := c.PostForm("valor")

			if nome != "" {
				if len(nome) < 2 {
					return nil, http.StatusBadRequest, &ErroValidacao{Mensagem: "O 'nome' deve ter pelo menos 2 caracteres."}
				}
				produtos[indice].Nome = nome
			}

			if categoria != "" {
				if len(categoria) < 2 {
					return nil, http.StatusBadRequest, &ErroValidacao{Mensagem: "A 'categoria' deve ter pelo menos 2 caracteres."}
				}
				produtos[indice].Categoria = categoria
			}

			if valorStr != "" {
				valorConvertido, err := strconv.ParseFloat(valorStr, 64)
				if err != nil {
					return nil, http.StatusBadRequest, &ErroValidacao{Mensagem: "Valor inválido. Informe um número válido."}
				}
				if valorConvertido <= 0 {
					return nil, http.StatusBadRequest, &ErroValidacao{Mensagem: "O valor deve ser maior que zero."}
				}
				produtos[indice].Valor = valorConvertido
			}

			imagem, err := c.FormFile("imagem")
			if err == nil {
				extensao := filepath.Ext(imagem.Filename)
				extensoesPermitidas := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
				if !extensoesPermitidas[extensao] {
					return nil, http.StatusBadRequest, &ErroValidacao{Mensagem: "Formato de imagem inválido. Use: .jpg, .jpeg, .png ou .gif."}
				}

				tamanhoMaximo := int64(5 * 1024 * 1024)
				if imagem.Size > tamanhoMaximo {
					return nil, http.StatusBadRequest, &ErroValidacao{Mensagem: "A imagem é muito grande. Tamanho máximo: 5MB."}
				}

				nomeImagem := codigo + extensao
				caminho := filepath.Join("uploads", nomeImagem)
				if err := c.SaveUploadedFile(imagem, caminho); err != nil {
					return nil, http.StatusInternalServerError, &ErroValidacao{Mensagem: "Não foi possível salvar a nova imagem."}
				}
				produtos[indice].Imagem = "uploads/" + nomeImagem
			}

			return &produtos[indice], http.StatusOK, nil
		}
	}
	return nil, http.StatusNotFound, &ErroValidacao{Mensagem: "Produto não encontrado."}
}

func RemoverProduto(codigo string) (int, error) {
	for indice, produto := range produtos {
		if produto.Codigo == codigo {
			produtos = append(produtos[:indice], produtos[indice+1:]...)
			return http.StatusOK, nil
		}
	}
	return http.StatusNotFound, &ErroValidacao{Mensagem: "Produto não encontrado."}
}

// ========== TIPO DE ERRO ==========
type ErroValidacao struct {
	Campo     string
	Mensagem  string
}

func (e *ErroValidacao) Error() string {
	return e.Mensagem
}