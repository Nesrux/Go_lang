package ponteiros

import "fmt"

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Retirar(qtda Bitcoin) {
	c.saldo -= qtda
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Depositar(qtda Bitcoin) {
	c.saldo += qtda
}
