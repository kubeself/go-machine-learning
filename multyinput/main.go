package main

import "fmt"

func main() {
	toes := []float64{8.5, 9.5, 9.9, 9.0}   // 脚趾数
	wlrec := []float64{0.65, 0.8, 0.8, 0.9} // 过去是否胜利
	nfans := []float64{1.2, 1.3, 0.5, 1.0}  // 粉丝数

	weight := []float64{0.1, 0.2, 0} //每个条件的权重
	for i := range toes {
		input := make([]float64, 0)
		input = append(input, toes[i])
		input = append(input, wlrec[i])
		input = append(input, nfans[i])
		fmt.Println(neuralNetwork(input, weight))
	}

}

func neuralNetwork(input []float64, weight []float64) (float64, error) {
	// 求两个向量的点积(加权沉积的和)
	res := float64(0)
	if len(input) != len(weight) {
		return 0, fmt.Errorf("len wrong")
	}
	for i, v := range input {
		res += v * weight[i]
	}
	return res, nil

}
