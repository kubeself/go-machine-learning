package main

import (
	"fmt"
	"math"
)

/**
决策树
*/
func main() {

	labelSet := []string{"yes", "yes", "no", "no", "no"}
	res := CalcShannonEnt(labelSet)
	fmt.Println(res)

	labelSet = []string{"maybe", "yes", "no", "no", "no"}
	res = CalcShannonEnt(labelSet)
	fmt.Println(res)
}

/**
计算给定数据集合的香农熵
*/
func CalcShannonEnt(labelSet []string) float64 {
	numEntries := len(labelSet)
	labelCnt := make(map[string]int)

	for _, v := range labelSet {
		if _, ok := labelCnt[v]; ok == false {
			labelCnt[v] = 0
		}
		labelCnt[v]++
	}
	fmt.Println(labelCnt)

	shannonEnt := float64(0)
	for _, v := range labelCnt {
		prob := float64(v) / float64(numEntries)
		shannonEnt -= prob * math.Log2(prob) // 每个词出现概率 * log以2为底求值
	}
	return shannonEnt
}

/**
  按照特征拆分数据
   @param axis : 抽取第几列数据
*/
func SplitDataSet(dataSet [][]int, labelSet []string, axis int, value int) [][]int {
	retDataSet := make([][]int, 0)
	for _, v := range dataSet {
		if v[axis] == value {
			//如果相等的话进行抽取
			retDataSet = append(retDataSet, v)
		}

	}
	return retDataSet
}
