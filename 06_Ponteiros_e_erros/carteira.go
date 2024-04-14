package ponteiros

import (
	"errors"
	"fmt"
)

var ErroSaldoInsulficiente = errors.New("não é possível retirar: saldo insuficiente")

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

func (c *Carteira) Retirar(qtda Bitcoin) error {
	if qtda > c.saldo {
		return ErroSaldoInsulficiente
	}
	c.saldo -= qtda
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Depositar(qtda Bitcoin) {
	c.saldo += qtda
}
