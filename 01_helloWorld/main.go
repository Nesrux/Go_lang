package main

import "fmt"

const prefixPortugues = "Olá, "

const prefixEspanhol = "Hola, "
const espanhol = "espanhol"

const prefixFrances = "Bonjur, "
const frances = "frances"

const prefixRusso = "Priviet, "
const russo = "russo"

func main() {
	fmt.Println(Ola("mundo", ""))
}
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoSaudacao(idioma) + nome

}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixFrances
	case espanhol:
		prefixo = prefixEspanhol
	case russo:
		prefixo = prefixRusso
	default:
		prefixo = prefixPortugues
	}
	return
}
