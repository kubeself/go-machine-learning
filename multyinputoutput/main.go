package main

import "fmt"

func main() {
	weight := [][]float64{
		{0.1, 0.1, -0.3}, //是否受伤
		{0.1, 0.2, 0},    //是否胜利
		{0.0, 1.3, 0.1},  //是否难过
	}

	toes := []float64{8.5, 9.5, 9.9, 9.0}   // 脚趾数
	wlrec := []float64{0.65, 0.8, 0.8, 0.9} // 过去是否胜利
	nfans := []float64{1.2, 1.3, 0.5, 1.0}  // 粉丝数

	//对输入的数据.每一列, 执行了3次独立的加权和操作,产生3个预测结果
	for i := range toes {
		data := []float64{toes[i], wlrec[i], nfans[i]}
		fmt.Println(neuralNetwork(data, weight))
	}

}

func neuralNetwork(data []float64, weight [][]float64) []float64 {
	res := make([]float64, 0)
	for _, w := range weight {
		val := float64(0)
		for j, v := range data {
			//val:=v1*w[0] + v2 * w[1] + v3 * w[2]
			val += v * w[j]
		}
		res = append(res, val)
	}
	return res
}
