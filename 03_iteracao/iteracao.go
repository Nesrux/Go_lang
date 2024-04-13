package iteracao

func Repetir(char string, rep int) string {
	var repeticoes string
	for i := 0; i < rep; i++ {
		repeticoes += char
	}
	return repeticoes
}

func ForInfinito(stop int) int {
	i := 0
	for {
		if i == stop {
			break
		}
		i++
	}
	return i
}
