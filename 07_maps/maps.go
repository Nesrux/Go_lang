package maps

type Dicionario map[string]string

func Busca(dicionario map[string]string, palavra string) string {
	return dicionario[palavra]
}

func (d Dicionario) Busca(palavra string) string {
	return d[palavra]
}
