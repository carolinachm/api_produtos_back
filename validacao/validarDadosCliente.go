package validacao

import (
	"api_produtos/erro"
	"regexp"
)

// ========== VALIDAÇÕES ==========
func ValidarCliente(nome, email, senha, telefone, endereco, cidade, uf, cep string) error {
	if nome == "" {
		return &erro.ErroValidacao{Campo: "nome", Mensagem: "O campo 'nome' é obrigatório."}
	}
	if len(nome) < 2 {
		return &erro.ErroValidacao{Campo: "nome", Mensagem: "O 'nome' deve ter pelo menos 2 caracteres."}
	}

	if email == "" {
		return &erro.ErroValidacao{Campo: "email", Mensagem: "O campo 'email' é obrigatório."}
	}
	regexEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !regexEmail.MatchString(email) {
		return &erro.ErroValidacao{Campo: "email", Mensagem: "Formato de e-mail inválido."}
	}

	if senha == "" {
		return &erro.ErroValidacao{Campo: "senha", Mensagem: "O campo 'senha' é obrigatório."}
	}
	if len(senha) < 6 {
		return &erro.ErroValidacao{Campo: "senha", Mensagem: "A senha deve ter pelo menos 6 caracteres."}
	}

	if endereco == "" {
		return &erro.ErroValidacao{Campo: "endereco", Mensagem: "O campo 'endereco' é obrigatório."}
	}
	if cidade == "" {
		return &erro.ErroValidacao{Campo: "cidade", Mensagem: "O campo 'cidade' é obrigatório."}
	}
	if uf == "" {
		return &erro.ErroValidacao{Campo: "uf", Mensagem: "O campo 'uf' é obrigatório."}
	}
	if cep == "" {
		return &erro.ErroValidacao{Campo: "cep", Mensagem: "O campo 'cep' é obrigatório."}
	}

	return nil
}