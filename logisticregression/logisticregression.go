package main

import (
	"fmt"
	"math"
)

/**
逻辑回归
*/
func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{0, 0, 0, 1, 1}

	alpha := 0.1
	iters := 1500
	m := len(x)

	var theta0 float64
	var theta1 float64

	for j := 0; j < iters; j++ {
		cost0 := float64(0)
		for i, _ := range x {
			cost0 += h(x[i], theta0, theta1) - y[i]
		}
		temp0 := theta0 - alpha*(1/float64(m))*cost0
		cost1 := float64(0)
		for i, _ := range x {
			cost1 += (h(x[i], theta0, theta1) - y[i]) * x[i]
		}
		temp1 := theta1 - alpha*(1/float64(m))*cost1
		theta0 = temp0
		theta1 = temp1

		fmt.Printf("%.2f, %.2f, %.2f \n", theta0, theta1, cost(theta0, theta1, x, y))
	}
}

/**
  将一个数转化为概率值 位于0-1之间
*/
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(z*(-1)))
}

/**
  逻辑回归方法，
   调用sigmoid的线性回归
   将线性分类结果概率化

*/
func h(x, theta0, theta1 float64) float64 {
	return sigmoid(theta0 + theta1*x)
}

/**
计算均差
*/
func cost(theta0, theta1 float64, x, y []float64) float64 {
	var distance float64
	for i, v := range x {
		if y[i] == 1 {
			//取对数
			distance += -math.Log(h(v, theta0, theta1))
		} else {
			distance += math.Log(1 - h(v, theta0, theta1))
		}
	}
	distance = distance / float64(len(x))
	return distance
}
