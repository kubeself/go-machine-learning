package main

import (
	"fmt"
	"math"
)

func main() {

	x := [][]float64{
		{1, 1, 2},
		{1, 2, 6},
		{1, 1.5, 7},
		{1, 1, 0},
		{1, 2, 3},
		{1, 2.5, 6},
	}

	y := []float64{1, 1, 1, 0, 0, 0}
	//假设每一列对结果的权重都是相同的
	//w := []float64{1,1,1}
	//收集每一列数据到dataX
	dataX := make([][]float64, len(x[0]))
	for i := range dataX {
		dataX[i] = make([]float64, len(x))
		for j := range dataX[i] {
			dataX[i][j] = x[j][i]
		}
	}
	bestTheta := make([][]float64, len(x[0]))
	//分别计算每一列的最佳theta0 & theta1
	for i := range bestTheta {
		bestTheta[i] = make([]float64, 2) // 存储theta0 & theta1
		theta0, theta1 := calTheta(0.1, 200, dataX[i], y)
		bestTheta[i][0] = theta0
		bestTheta[i][1] = theta1
		//disp([theta_0, theta_1, cost(theta_0, theta_1, x, y)]);
		fmt.Printf("%.2f,%.2f,%.2f\n", theta0, theta1, cost(theta0, theta1, dataX[i], y))
	}
	fmt.Println(bestTheta)
}

/**
最佳推测
*/
func h(x, theta0, theta1 float64) float64 {
	return sigmoid(theta0 + theta1*x)
}

func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(z*(-1)))
}

func calTheta(alpha float64, iters int, x []float64, y []float64) (float64, float64) {
	m := len(x)
	theta0 := float64(0)
	theta1 := float64(0)

	for i := 0; i < iters; i++ {
		cost0 := float64(0)
		for j := 0; j < m; j++ {
			cost0 += h(x[j], theta0, theta1) - y[j]
		}

		cost1 := float64(0)
		for j := 0; j < m; j++ {
			cost1 += (h(x[j], theta0, theta1) - y[j]) * x[j]
		}

		theta0 -= cost0 / float64(m) * alpha
		theta1 -= cost1 / float64(m) * alpha
	}

	return theta0, theta1
}

func cost(theta0, theta1 float64, x, y []float64) float64 {
	var distance float64
	for i, v := range x {
		if y[i] == 1 {
			//取对数
			distance += -1 * math.Log(h(v, theta0, theta1))
		} else {
			distance += -1 * math.Log(1-h(v, theta0, theta1))
		}
	}
	distance = distance / float64(len(x))
	return distance
}
