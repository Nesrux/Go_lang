package main

import (
	"fmt"
	"io"
	"os"
)

func Cumprimenta(esc io.Writer, nome string) {
	fmt.Fprintf(esc, "Olá, %s", nome)
}

func main() {
	Cumprimenta(os.Stdout, "Marcos")
}
