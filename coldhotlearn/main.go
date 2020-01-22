package main

import (
	"fmt"
	"math"
)

func main() {

	//假设初始权重=0.5, 输入0.5, 目标0.8 ,每次调整权重0.001,
	weight := 0.5
	input := 0.5
	goalPrediction := 0.8
	step := 0.001
	loops := 1101

	//开始循环计算
	for i := 0; i < loops; i++ {
		realVal := weight * input
		err := math.Pow(realVal-goalPrediction, 2)
		fmt.Printf("current err:%.6f\n", err)
		if err == 0 {
			break
		}

		//接下去对up和down的步骤分别计算err
		upVal := (weight + step) * input
		upErr := math.Pow(upVal-goalPrediction, 2)

		downVal := (weight - step) * input
		downErr := math.Pow(downVal-goalPrediction, 2)

		if upErr < downErr {
			weight += step

		} else {
			weight -= step
		}

	}

	fmt.Println(weight)
	fmt.Println(weight * input)
}
