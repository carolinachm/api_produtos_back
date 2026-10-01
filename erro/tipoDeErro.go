package erro

// ========== TIPO DE ERRO ==========
type ErroValidacao struct {
	Campo     string
	Mensagem  string
}

func (e *ErroValidacao) Error() string {
	return e.Mensagem
}