package main

import (
	"fmt"
)

func main() {
	input := 0.65                      // 过去
	weight := []float64{0.3, 0.2, 0.9} // 是否受伤，是否胜利，是否难过
	fmt.Println(neuralNetwork(input, weight))
}

func neuralNetwork(input float64, weight []float64) float64 {
	res := float64(0)
	for _, v := range weight {
		res += v * input
	}
	return res
}
