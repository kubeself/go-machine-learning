package main

import "fmt"

func main() {
	toes := []float64{8.5, 9.5, 9.9, 9.0}   // 脚趾数
	wlrec := []float64{0.65, 0.8, 0.8, 0.9} // 过去是否胜利
	nfans := []float64{1.2, 1.3, 0.5, 1.0}  // 粉丝数

	//神经网络的堆叠, 一个网络的输出作为输入给下一个网络. 在图像分类上有用到
	w2 := [][]float64{
		{0.1, 0.2, -0.1}, //hid[0] = v1 * w1 + v2 * w2 + w3 *w3
		{-0.1, 0.1, 0.9}, //hid[1]
		{0.1, 0.4, 0.1},  //hid[2]
	}
	w1 := [][]float64{
		{0.3, 1.1, -0.3}, //是否受伤
		{0.1, 0.2, 0},    //是否胜利
		{0.0, 1.3, 0.1},  //是否难过
	}

	for i, v := range toes {
		data := []float64{v, wlrec[i], nfans[i]}
		hData := neuralNetwork(data, w2)
		fmt.Printf("hData = %v \n", hData)
		rData := neuralNetwork(hData, w1)
		fmt.Printf("rData = %v\n", rData)
	}

}

func neuralNetwork(data []float64, weights [][]float64) []float64 {
	res := make([]float64, 0)
	for _, w := range weights {
		val := float64(0)
		for j, v := range data {
			//val = v1*w[0]+v2*w[1]+v3*w[2]
			val += v * w[j]
		}
		res = append(res, val)
	}
	return res
}
