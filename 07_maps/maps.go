package maps

import "errors"

var ErrNaoEncontrado = errors.New("não foi possivel encontrar a palavra que você porcurou")

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
func (d Dicionario) Adiciona(palavra, definicao string) {
	d[palavra] = definicao
}
