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

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		verificaDeposito(t, carteira, Bitcoin(10))
	})
	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		verificaDeposito(t, carteira, Bitcoin(10))
	})
}
