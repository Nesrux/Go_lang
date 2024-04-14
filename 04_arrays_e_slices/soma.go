package main

import "fmt"

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
func SomaTudo(slice ...[]int) []int {
	var somas []int
	for _, nums := range slice {
		somas = append(somas, Soma(nums))
	}
	return somas
}

func SomaTodoResto(numerosSoma ...[]int) []int {
	var somas []int
	for _, nums := range numerosSoma {
		if len(nums) == 0 {
			somas = append(somas, 0)
		} else {
			final := nums[1:]
			somas = append(somas, Soma(final))
		}
	}
	return somas
}

func main() {
	var arr []int

	for i := 0; i < 1_000_000_000; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr[len(arr)-1:])
}
