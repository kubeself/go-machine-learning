package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
)

func main() {
	x := []float64{1000, 2000, 4000}       // area of houses
	y := []float64{200000, 250000, 300000} // price of houses

	//测试当theta_0 = 0 ， theta_1为 75 和 160 分别的结果
	c1 := cost(x, y, 0, 75)
	c2 := cost(x, y, 0, 160)

	fmt.Printf("theta1=75, cost=%.2f, theta1=160 , cost= %.2f\n", c1, c2)

	//使用gonum库的方法计算最佳值
	theta0, theta1 := stat.LinearRegression(x, y, nil, false)
	fmt.Printf("theta_0=%.2f, theta_1=%.2f, cost=%.2f\n", theta0, theta1, cost(x, y, theta0, theta1))
}

/**
  线性回归 y = theta_0 + theta_1 * x
   平均偏差值， 距离Y的总差的平均数

*/
func cost(x, y []float64, theta0, theta1 float64) float64 {
	var distance float64
	for k, v := range x {
		actY := theta0 + theta1*v
		distance += math.Abs(float64(actY) - float64(y[k]))
	}
	distance = distance / float64(len(x))
	return distance
}
