package main

import "fmt"

const prefixPortugues = "Olá, "

const prefixEspanhol = "Hola, "
const espanhol = "espanhol"

const prefixFrances = "Bonjur, "
const frances = "frances"

func main() {
	fmt.Println(Ola("mundo", ""))
}
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	if idioma == espanhol {
		return prefixEspanhol + nome
	}

	if idioma == frances {
		return prefixFrances + nome
	}

	return prefixPortugues + nome
}
