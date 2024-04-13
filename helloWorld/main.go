package main

import "fmt"

const prefix = "Olá, "

func main() {
	fmt.Println(Ola("mundo"))
}
func Ola(nome string) string {
	if(nome == ""){
		nome = "Mundo"
	}
	
	return prefix + nome
}
