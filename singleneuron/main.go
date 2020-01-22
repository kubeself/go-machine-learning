package main

import "fmt"

func main() {
	//单个神经元
	//输入脚趾数目，它的权重=0.1 。预测是否胜利
	data := []float64{8.5, 9.5, 10, 9}
	w := 0.1
	for _, k := range data {
		fmt.Println(neuralNetwork(k, w))
	}
}

func neuralNetwork(input float64, weight float64) float64 {
	return input * weight
}
