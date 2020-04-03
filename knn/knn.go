package main

import (
	"fmt"
	"sort"
)

func main() {
	// 二维数据
	data := [][]float64{
		{1, 1.1},
		{1, 1},
		{0, 0},
		{0, 0.1},
	}
	labels := []string{
		"A", "A", "B", "B",
	}

	res := Classify0([]float64{0, 0}, data, labels, 3)
	fmt.Println(res)
}

type CmpData struct {
	data  []float64
	dist  float64
	label string
}

func Classify0(testData []float64, data [][]float64, labels []string, k int) string {
	// 距离计算
	cmpDataList := make([]CmpData, 0)
	for i, v := range data {
		d := GetDistance(testData, v)
		cmpDataList = append(cmpDataList, CmpData{
			data:  v,
			dist:  d,
			label: labels[i],
		})
	}
	fmt.Println("distance:", cmpDataList)

	sort.Slice(cmpDataList[:], func(i, j int) bool {
		return cmpDataList[i].dist < cmpDataList[j].dist
	})

	fmt.Println("distance:", cmpDataList)

	// 获得最近的k个点
	kList := cmpDataList[:k]

	// 每个label出现的概率计算
	labelTimes := make(map[string]int)
	labelTimesVal := make([]int, 0)
	for _, v := range kList {
		if _, ok := labelTimes[v.label]; ok == false {
			labelTimes[v.label] = 1
		} else {
			labelTimes[v.label]++
		}
	}

	fmt.Println("labelTimes:", labelTimes)

	for _, v := range labelTimes {
		labelTimesVal = append(labelTimesVal, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(labelTimesVal)))

	fmt.Println("label times val:", labelTimesVal)
	for k, v := range labelTimes {
		for _, v2 := range labelTimesVal {
			if v == v2 {
				fmt.Println("found v, v2", v, v2)
				return k
			}
		}
	}
	// 将概率降序排列. 去第一个值
	return ""

}

/**
  公式:(x1-x2)^2 + (y1-y2)^2 + (z1- z2) ^2 + ....
*/
func GetDistance(testData, v []float64) float64 {
	sum := float64(0)
	for i, item := range testData {
		sum += (item - v[i]) * (item - v[i])
	}
	return sum
}
