package main

import (
	"fmt"
	"math"
)

//解决多元概率问题是从一元进行扩展的,每一元与结果直接计算theta0和theta1，最后进行相加除以1就是结果
//多元 将每一列作为一个独立因素，对Y有影响的独立因素，分别进行计算他们的theta0 和 theta 1
// y = theta0 + theta1 * X1
// y = theta0 + theta1 * X2
// y = theta0 + theta1 * X3
// 全部计算后会得到[(theta0,theta1),(theta0,theta1),(theta0,theta1)] 对于所有列的合适的取值。最后取个平均数就可以算出总概率
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
	bestTheta := calGroupTheta(x, y, 0.1, 1000)
	//fmt.Println(bestTheta)
	//得到 y = 0+0*x1+1.91+(-1.15)*x2+(-1.40)+(0.35*x3)
	//打印每一行的概率
	//计算后打印 x=1-10的概率值
	for i := 0; i < len(x); i++ {
		p := float64(0)
		for j := 0; j < len(x[i]); j++ {
			p += h(x[i][j], bestTheta[j][0], bestTheta[j][1])
		}
		p /= float64(len(x[i]))
		fmt.Printf("i = %d, p = %.2f\n", i, p)
	}

	//测试指定记录的概率： {1, 9, 888}
	testData := []float64{1, 2.8, 6.1}
	p := float64(0)
	for i, v := range testData {
		p += h(v, bestTheta[i][0], bestTheta[i][1])
	}
	p /= float64(len(testData))
	fmt.Printf("property for %v = %.2f\n", testData, p)
}

func calGroupTheta(x [][]float64, y []float64, alpha float64, loopNum int) [][]float64 {
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
		theta0, theta1 := calTheta(alpha, loopNum, dataX[i], y)
		bestTheta[i][0] = theta0
		bestTheta[i][1] = theta1
		//disp([theta_0, theta_1, cost(theta_0, theta_1, x, y)]);
		fmt.Printf("theta0 = %.2f,theta1 = %.2f,cost = %.2f\n", theta0, theta1, cost(theta0, theta1, dataX[i], y))
	}
	return bestTheta
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
