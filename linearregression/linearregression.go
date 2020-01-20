package main

import (
	"fmt"
	"math"
)

/**
获得一组数据适当的倍数值，
 并且返回缩小后的数据。
*/
func getMultiple(x []float64) (float64, []float64) {
	minI := 0
	min := x[minI]
	fmt.Println("min =", min)
	mul := float64(1)
	for i, v := range x {
		if v < min && v != 0 {
			min = v
			minI = i
		}
	}
	mul = min
	fmt.Println("min =", min)
	fmt.Println("find mul=", mul)
	res := make([]float64, len(x))
	mulTimes := getTenTimes(mul)
	for i, v := range x {
		res[i] = v / mulTimes
	}
	return mulTimes, res

}

/**
  输入的数让它变成0-10之间的个位数，并且返回需要除以的倍数
*/
func getTenTimes(x float64) float64 {
	cnt := float64(0)
	for {
		if x < 10 && x > 0 {
			break
		}
		x /= 10
		cnt++
	}
	return float64(math.Pow(10, cnt))
}

func main() {
	x := []float64{1000, 2000, 4000} // area of houses -> /1000
	mulx, x := getMultiple(x)
	fmt.Println(x)

	//找出x的元素值在1-x范围内，
	y := []float64{200000, 250000, 300000} // price of houses   ->	/100000，
	muly, y := getMultiple(y)
	fmt.Println(y)
	//找出y的元素在1-x范围内

	//测试当theta_0 = 0 ， theta_1为 75 和 160 分别的结果
	c1 := cost(x, y, 0, 75)
	c2 := cost(x, y, 0, 160)

	fmt.Printf("theta1=75, cost=%.2f, theta1=160 , cost= %.2f\n", c1, c2)

	//使用gonum库的方法计算最佳值
	alpha := 0.1
	loopNum := 100000
	theta0, theta1 := calTheta(alpha, loopNum, x, y)
	theta0 *= muly
	theta1 *= mulx * alpha

	fmt.Println("theta0=", theta0, " theta1=", theta1)

	//fmt.Printf("theta_0=%.2f, theta_1=%.2f, cost=%.2f\n", theta0, theta1, cost(x, y, theta0, theta1))
	//theta0, theta1 := stat.LinearRegression(x, y, nil, false)

	//使用最低消耗的参数来估算价格
	res := h(3000, theta0, theta1)
	fmt.Printf("beat price for %df = %.2f\n", 3000, res)
}

/**
  线性回归 y = theta_0 + theta_1 * x
   平均偏差值， 距离Y的总差的平均数

*/
func cost(x, y []float64, theta0, theta1 float64) float64 {
	var distance float64
	for k, v := range x {
		actY := theta0 + theta1*v
		distance += math.Abs(actY - y[k])
	}
	distance = distance / float64(len(x))
	return distance
}

/**
最佳推测
*/
func h(x, theta0, theta1 float64) float64 {
	return theta0 + theta1*x
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
