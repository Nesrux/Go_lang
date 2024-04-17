package maps

const (
	ErrNaoEncontrado    = ErrDicionario("não foi possivel encontrar a palavra que você porcurou")
	ErrPalavraExistente = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

type Dicionario map[string]string

func Busca(dicionario map[string]string, palavra string) string {
	return dicionario[palavra]
}

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, exist := d[palavra]
	if !exist {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}
func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}
	return nil
}
