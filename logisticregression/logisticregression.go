package main

import (
	"fmt"
	"math"
)

/**
逻辑回归
*/
func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{0, 0, 0, 1, 1, 1}

	alpha := 0.1 // （斜率）
	theta0, theta1 := calTheta(x, y, 1500, alpha)

	//计算后打印 x=1-10的概率值
	for i := 0; i < 6; i++ {
		fmt.Printf("x=%d, p=%.2f\n", i, h(float64(i), theta0, theta1)*float64(100))
	}

}

func calTheta(x, y []float64, iters int, alpha float64) (float64, float64) {

	m := len(x)

	//theta0 & theta1 需要同事更新
	var theta0 float64
	var theta1 float64

	//y = 1 / (1 + math.Exp((theta0 + theta1*x)*(-1)))
	//按照指定的循环次数
	for j := 0; j < iters; j++ {
		//计算分母均差
		cost0 := float64(0)
		for i, _ := range x {
			cost0 += h(x[i], theta0, theta1) - y[i]
		}
		cost0 = cost0 / float64(m)

		//计算分子均差
		cost1 := float64(0)
		for i, _ := range x {
			cost1 += (h(x[i], theta0, theta1) - y[i]) * x[i]
		}
		cost1 = cost1 / float64(m)

		//得到下一个点的位置
		theta0 = theta0 - alpha*cost0
		theta1 = theta1 - alpha*cost1

		fmt.Printf("%.2f, %.2f, %.2f \n", theta0, theta1, cost(theta0, theta1, x, y))
	}

	return theta0, theta1

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
