package service

import (
	"api_produtos/erro"
	"api_produtos/models"
	"api_produtos/validacao"
	"net/http"

	"github.com/gin-gonic/gin"
	"uuid"
)

// ========== BANCO EM MEMÓRIA ==========
var clientes = []models.ClienteModel{}

// ListarTodosClientes retorna todos os clientes
func ListarTodosClientes() []models.ClienteModel {
	return clientes
}

// CadastrarCliente — corrigido para usar campos do cliente
func CadastrarCliente(c *gin.Context) (models.ClienteModel, int, error) {
	// Obter dados do formulário
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	senha := c.PostForm("senha")
	telefone := c.PostForm("telefone")
	endereco := c.PostForm("endereco")
	complemento := c.PostForm("complemento")
	cidade := c.PostForm("cidade")
	uf := c.PostForm("uf")
	cep := c.PostForm("cep")

	// Validar
	err := validacao.ValidarCliente(nome, email, senha, telefone, endereco, cidade, uf, cep)
	if err != nil {
		return models.ClienteModel{}, http.StatusBadRequest, err
	}

	// Gerar código
	codigo := uuid.New().String()

	// Criar cliente
	novoCliente := models.ClienteModel{
		Codigo:      codigo,
		Nome:        nome,
		Email:       email,
		Senha:       senha,
		Telefone:    telefone,
		Endereco:    endereco,
		Complemento: complemento,
		Cidade:      cidade,
		UF:          uf,
		CEP:         cep,
	}

	// Salvar
	clientes = append(clientes, novoCliente)
	return novoCliente, http.StatusCreated, nil
}

// AlterarCliente — corrigido para usar lista clientes e campos corretos
func AlterarCliente(codigo string, c *gin.Context) (*models.ClienteModel, int, error) {
	for indice, cliente := range clientes {
		if cliente.Codigo == codigo {
			// Obter dados
			nome := c.PostForm("nome")
			email := c.PostForm("email")
			senha := c.PostForm("senha")
			telefone := c.PostForm("telefone")
			endereco := c.PostForm("endereco")
			complemento := c.PostForm("complemento")
			cidade := c.PostForm("cidade")
			uf := c.PostForm("uf")
			cep := c.PostForm("cep")

			// Atualizar apenas o que foi enviado
			if nome != "" {
				if len(nome) < 2 {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "O 'nome' deve ter pelo menos 2 caracteres."}
				}
				clientes[indice].Nome = nome
			}
			if email != "" {
				clientes[indice].Email = email
			}
			if senha != "" {
				if len(senha) < 6 {
					return nil, http.StatusBadRequest, &erro.ErroValidacao{Mensagem: "A senha deve ter pelo menos 6 caracteres."}
				}
				clientes[indice].Senha = senha
			}
			if telefone != "" {
				clientes[indice].Telefone = telefone
			}
			if endereco != "" {
				clientes[indice].Endereco = endereco
			}
			if complemento != "" {
				clientes[indice].Complemento = complemento
			}
			if cidade != "" {
				clientes[indice].Cidade = cidade
			}
			if uf != "" {
				clientes[indice].UF = uf
			}
			if cep != "" {
				clientes[indice].CEP = cep
			}

			return &clientes[indice], http.StatusOK, nil
		}
	}
	return nil, http.StatusNotFound, &erro.ErroValidacao{Mensagem: "Cliente não encontrado."}
}

// RemoverCliente — corrigido para usar lista clientes
func RemoverCliente(codigo string) (int, error) {
	for indice, cliente := range clientes {
		if cliente.Codigo == codigo {
			clientes = append(clientes[:indice], clientes[indice+1:]...)
			return http.StatusOK, nil
		}
	}
	return http.StatusNotFound, &erro.ErroValidacao{Mensagem: "Cliente não encontrado."}
}