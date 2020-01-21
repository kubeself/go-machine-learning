package main

import (
	"fmt"
	"math"
)

func main() {

	//x = 1:100;
	//y = [0, 0, 0, 1, 1, 1, 1, zeros(1, 93)];
	x := make([]float64, 0)
	for i := 1; i <= 100; i++ {
		x = append(x, float64(i))
	}

	y := make([]float64, 0)
	for i := 0; i < 100; i++ {
		if i < 3 {
			y = append(y, 0)
		} else if i < 7 {
			y = append(y, 1)
		} else {
			y = append(y, 0)
		}
	}

	//fmt.Println(x)
	//fmt.Println(y)

	theta := calTheta(0.1, 1500, x, y)

	for i := 0; i < 100; i++ {
		fmt.Println(h(float64(i), theta))
	}

}

/**
  将一个数转化为概率值 位于0-1之间
*/
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(z*(-1)))
}

/**
  非线性逻辑回归公式
  theta(1) + theta(2) * x + theta(3) * ((x - theta(4))^2)

*/
func h(x float64, theta []float64) float64 {
	return sigmoid(theta[0] + theta[1]*x + theta[2]*(math.Pow(x-theta[3], 2)))
}

func calTheta(alpha float64, iters int, x []float64, y []float64) []float64 {
	m := len(x)

	theta := make([]float64, 4)
	theta0 := float64(0)
	theta1 := float64(0)
	theta2 := float64(0)
	theta3 := float64(0)

	for i := 0; i < iters; i++ {
		theta[0] = theta0
		theta[1] = theta1
		theta[2] = theta2
		theta[3] = theta3

		cost0 := float64(0)
		for j := 0; j < m; j++ {
			cost0 += h(x[j], theta) - y[j]
		}

		//theta(1) + theta(2) * x + theta(3) * ((x - theta(4))^2)
		cost1 := float64(0)
		for j := 0; j < m; j++ {
			cost1 += (h(x[j], theta) - y[j]) * x[j]
		}

		cost2 := float64(0)

		for j := 0; j < m; j++ {
			cost2 += (h(x[j], theta) - y[j]) * math.Pow(x[j], 2)
		}
		theta0 -= cost0 / float64(m) * alpha
		theta1 -= cost1 / float64(m) * alpha
		theta2 -= cost2 / float64(m) * alpha
	}

	theta[0] = theta0
	theta[1] = theta1
	theta[2] = theta2
	theta[3] = theta3

	return theta
}
