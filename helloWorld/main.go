package main

import "fmt"

const prefix = "Olá, "

func main() {
	fmt.Println(Ola("mundo"))
}
func Ola(nome string) string {
	return prefix + nome
}
