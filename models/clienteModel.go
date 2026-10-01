package models

// Cliente — usuário que faz pedidos
type ClienteModel struct{
	
	Codigo string `json:"codigo"`
	Nome string `json:"nome"`
	Email string `json:"email"`
	Telefone string `json:"telefone"`
	Senha string `json:"senha,omitempty"`
	Endereco string `json:"endereco"`
	Complemento string `json:"complemento,omitempty"`
	Cidade string `json:"cidade"`
	UF string `json:"uf"`
	CEP string `json:"cep"`


}