package models

// Produto — estrutura de dados do produto
type ProdutoModel struct {
	Codigo    string  `json:"codigo" form:"codigo"`
	Nome      string  `json:"nome" form:"nome"`
	Categoria string  `json:"categoria" form:"categoria"`
	Valor     float64 `json:"valor" form:"valor"`
	Imagem    string  `json:"imagem" form:"imagem"`
}