package ponteiros

import (
	"testing"
)

func TestXxx(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcoin(10))

	res := carteira.Saldo()
	esp := Bitcoin(10)

	if res != esp {
		t.Errorf("resultado %s, esperado %s", res, esp)
	}
}

func Test_Carteira(t *testing.T) {
	verificaDeposito := func(t *testing.T, carteira Carteira, esp Bitcoin) {
		t.Helper()
		res := carteira.Saldo()
		if res != esp {
			t.Errorf("resultado %s, esperado %s", res, esp)
		}
	}
	confirmaErro := func(t *testing.T, erro error, esperado string) {
		t.Helper()
		if erro == nil {
			t.Fatal("Esperava um erro, mas nenhum aconteceu")
		}
		resultado := erro.Error()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		verificaDeposito(t, carteira, Bitcoin(10))
	})
	t.Run("Retirar com saldo Suficiente", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		verificaDeposito(t, carteira, Bitcoin(10))
	})
	t.Run("Retirar com saldo insulficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		Carteira := Carteira{saldoInicial}
		err := Carteira.Retirar(Bitcoin(100))

		verificaDeposito(t, Carteira, saldoInicial)
		confirmaErro(t, err, ErroSaldoInsulficiente.Error())
	})
}
