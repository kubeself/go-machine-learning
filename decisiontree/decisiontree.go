package main

import (
	"fmt"
	"math"
)

/**
  决策树
*/
func main() {

	dataSet := createDataSet()
	//dataSet2 := createDataSet2()

	/*
		res := CalcShannonEnt(dataSet)
		fmt.Println(res)


		res = CalcShannonEnt(dataSet2)
		fmt.Println(res)


		res1:=SplitDataSet(dataSet, 0 , 1)
		res2 := SplitDataSet(dataSet, 0, 0)

		fmt.Println(res1)
		fmt.Println(res2)


	*/
	res3 := ChooseBestFeatureToSplit(dataSet)
	fmt.Println(res3)

}

func createDataSet() [][]int {
	dataSet := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	return dataSet
}

func createDataSet2() [][]int {
	dataSet := [][]int{
		{1, 1, 2},
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	return dataSet
}

/**
  计算给定数据集合的香农熵:(结果分布的复杂度)
*/
func CalcShannonEnt(dataSet [][]int) float64 {
	numEntries := len(dataSet)
	numFeature := len(dataSet[0]) - 1
	labelCnt := make(map[int]int)

	//根据给定的数据, 看结果列复杂度是怎样的
	for _, v := range dataSet {
		l := v[numFeature]

		if _, ok := labelCnt[l]; ok == false {
			labelCnt[l] = 0
		}
		labelCnt[l]++
	}
	fmt.Println("label cnt=", labelCnt)

	shannonEnt := float64(0)
	for _, v := range labelCnt {
		prob := float64(v) / float64(numEntries)
		shannonEnt -= prob * math.Log2(prob) // 每个词出现概率 * log以2为底求值
	}
	return shannonEnt
}

/**
  按照特征拆分数据, 返回指定特征的指定值的行记录
   @param axis : 抽取第几列数据
   @param value : 要匹配的值
*/
func SplitDataSet(dataSet [][]int, axis int, value int) [][]int {
	retDataSet := make([][]int, 0)

	fmt.Println(dataSet)
	for _, v := range dataSet {
		if v[axis] == value {
			//如果相等的话进行抽掉判断的列 , 添加到结果集中
			reduceFeatVec := make([]int, 0)
			for j, w := range v {
				if j == axis {
					continue
				}
				reduceFeatVec = append(reduceFeatVec, w)
			}
			retDataSet = append(retDataSet, reduceFeatVec)
			fmt.Println("retDataSet:", retDataSet)
		}
	}
	fmt.Println("all-retDataSet:", retDataSet)
	return retDataSet
}

func ChooseBestFeatureToSplit(dataSet [][]int) int {
	// len of feature
	numFeature := len(dataSet[0]) - 1 // 特点一共有多少列
	fmt.Println(numFeature)
	baseEntropy := CalcShannonEnt(dataSet) // 复杂度计算
	baseInfoGain := 0.0
	bestFeature := -1

	for i := 0; i < numFeature; i++ {
		//取得当前列的值的MAP
		uniMap := make(map[int]int)
		for _, v := range dataSet {
			if _, ok := uniMap[v[i]]; ok == false {
				uniMap[v[i]] = 0
			}
			uniMap[v[i]]++
		}

		//当前列复杂度初始化
		newEntropy := 0.0
		//计算当前列   每个值的出现概率 * 具有当前值的行记录(排除当前列的)
		fmt.Println("dataset?", dataSet)
		for k := range uniMap {
			subDataSet := SplitDataSet(dataSet, i, k)
			prob := float64(len(subDataSet)) / float64(len(dataSet))
			fmt.Println("subDataSet", subDataSet)
			fmt.Println("dataset??", dataSet)
			fmt.Println("k, i", k, i)
			fmt.Println("i= ?, prob * shanno:", i, prob, "*", CalcShannonEnt(subDataSet))
			newEntropy += prob * CalcShannonEnt(subDataSet)
		}
		fmt.Println("dataset???", dataSet)

		infoGain := baseEntropy - newEntropy
		fmt.Println("baseEntropy:", baseEntropy)
		fmt.Println("NewEntropy:", newEntropy)
		fmt.Println("infoGain", infoGain)
		if infoGain > baseInfoGain {
			baseInfoGain = infoGain
			bestFeature = i
		}

	}
	return bestFeature

}
