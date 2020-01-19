package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
)

/**
逻辑回归
*/
func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{0, 0, 0, 1, 1, 1}

	//alpha := 0.1 // （斜率）
	theta0, theta1 := stat.LinearRegression(x, y, nil, false)

	//计算后打印 x=1-10的概率值
	for i := 0; i < 6; i++ {
		fmt.Printf("x=%d, p=%.2f\n", i, h(float64(i), theta0, theta1)*float64(100))
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
打印距离的函数
*/
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
