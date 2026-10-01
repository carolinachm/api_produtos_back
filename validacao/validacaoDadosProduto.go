package validacao

import (
	
	"api_produtos/erro"
	
	"path/filepath"
	"strconv"

	"uuid"

	"github.com/gin-gonic/gin"
)

// ========== VALIDAÇÕES ==========
func validarProduto(nome, categoria, valorStr string) (float64, error) {
	if nome == "" {
		return 0, &erro.ErroValidacao{Campo: "nome", Mensagem: "O campo 'nome' é obrigatório."}
	}
	if len(nome) < 2 {
		return 0, &erro.ErroValidacao{Campo: "nome", Mensagem: "O 'nome' deve ter pelo menos 2 caracteres."}
	}
	if categoria == "" {
		return 0, &erro.ErroValidacao{Campo: "categoria", Mensagem: "O campo 'categoria' é obrigatório."}
	}
	if len(categoria) < 2 {
		return 0, &erro.ErroValidacao{Campo: "categoria", Mensagem: "A 'categoria' deve ter pelo menos 2 caracteres."}
	}
	if valorStr == "" {
		return 0, &erro.ErroValidacao{Campo: "valor", Mensagem: "O campo 'valor' é obrigatório."}
	}
	valorConvertido, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		return 0, &erro.ErroValidacao{Campo: "valor", Mensagem: "Valor inválido. Informe um número válido (ex: 5500 ou 199.90)."}
	}
	if valorConvertido <= 0 {
		return 0, &erro.ErroValidacao{Campo: "valor", Mensagem: "O valor deve ser maior que zero."}
	}
	return valorConvertido, nil
}

func validarImagem(c *gin.Context) (string, string, error) {
	imagem, err := c.FormFile("imagem")
	if err != nil {
		return "", "", &erro.ErroValidacao{Campo: "imagem", Mensagem: "A imagem é obrigatória. Selecione um arquivo de imagem."}
	}

	extensao := filepath.Ext(imagem.Filename)
	extensoesPermitidas := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !extensoesPermitidas[extensao] {
		return "", "", &erro.ErroValidacao{Campo: "imagem", Mensagem: "Formato de imagem inválido. Use: .jpg, .jpeg, .png ou .gif."}
	}

	tamanhoMaximo := int64(5 * 1024 * 1024)
	if imagem.Size > tamanhoMaximo {
		return "", "", &erro.ErroValidacao{Campo: "imagem", Mensagem: "A imagem é muito grande. Tamanho máximo: 5MB."}
	}

	codigo := uuid.New().String()
	nomeImagem := codigo + extensao
	caminho := filepath.Join("uploads", nomeImagem)

	if err := c.SaveUploadedFile(imagem, caminho); err != nil {
		return "", "", &erro.ErroValidacao{Campo: "imagem", Mensagem: "Não foi possível salvar a imagem. Verifique a pasta 'uploads'."}
	}

	return nomeImagem, caminho, nil
}