package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	label string
	no    *TreeNode
	yes   *TreeNode
	val   int
}

/**
  决策树
*/
func main() {
	dataSet, labels := createDataSet()
	tree := CreateTree(dataSet, labels)
	fmt.Println(tree)
	tree.Print()
}

func createDataSet() ([][]int, []string) {
	dataSet := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	labels := []string{
		"no surfacing",
		"flippers",
	}

	return dataSet, labels
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

		}
	}
	return retDataSet
}

/**
  选择最好的特征列
*/
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

		for k := range uniMap {
			subDataSet := SplitDataSet(dataSet, i, k)
			prob := float64(len(subDataSet)) / float64(len(dataSet))

			newEntropy += prob * CalcShannonEnt(subDataSet)
		}

		infoGain := baseEntropy - newEntropy
		if infoGain > baseInfoGain {
			baseInfoGain = infoGain
			bestFeature = i
		}

	}
	return bestFeature

}

/**
  返回出现 次数最多的分类 . 返回主要决定的特征
*/
func majorityCnt(classList []int) int {
	classCnt := make(map[int]int, 0)
	for vote := range classList {
		if _, ok := classCnt[vote]; ok == false {
			classCnt[vote] = 0
		}
		classCnt[vote]++
	}

	maxIndex := 0
	maxVal := 0

	for k, v := range classCnt {
		if v > maxVal {
			maxIndex = k
			maxVal = v
		}
	}
	return maxIndex
}

func CreateTree(dataSet [][]int, labels []string) *TreeNode {
	fmt.Println("create Tree:", dataSet)
	//结果分类:
	featNum := len(dataSet[0]) - 1
	classList := make([]int, 0)
	for _, item := range dataSet {
		classList = append(classList, item[featNum])
	}

	//分类内所有元素完全相同, 停止分类
	firstVal := classList[0]
	allSame := true
	for _, v := range classList {
		if v != firstVal {
			allSame = false
			break
		}
	}

	if allSame == true {
		fmt.Println("return first Val:", firstVal)
		return &TreeNode{
			val: firstVal,
		}
	}

	//所有特征都已经遍历完事,返回最主要的决定值
	if len(dataSet[0]) == 1 {
		return &TreeNode{
			val: majorityCnt(classList),
		}
	}
	//开始找最有决定权的列
	bestFeat := ChooseBestFeatureToSplit(dataSet)

	myTree := TreeNode{
		label: labels[bestFeat],
	}
	//del elem by index
	labels = append(labels[:bestFeat], labels[bestFeat+1:]...)

	//获取最佳列的值map
	uniVal := make(map[int]int)
	for _, v := range dataSet {
		val := v[bestFeat]
		if _, ok := uniVal[val]; ok == false {
			uniVal[val] = 1
		}
	}

	/**
	  遍历最佳列的值列表,
	*/
	for v, _ := range uniVal {
		subLabels := labels[:]
		if v == 0 {
			myTree.no = CreateTree(SplitDataSet(dataSet, bestFeat, v), subLabels)
			fmt.Println("[myTree.child]:", myTree.no)
		} else if v == 1 {
			myTree.yes = CreateTree(SplitDataSet(dataSet, bestFeat, v), subLabels)
			fmt.Println("[myTree.child]:", myTree.yes)
		}

	}

	return &myTree

}

func (n *TreeNode) Print() {
	//中, 左, 右

	if n.label != "" {
		fmt.Println("label=", n.label)
	}
	if n.no == nil && n.yes == nil {
		fmt.Println("val=", n.val)
	}

	if n.no != nil {
		fmt.Println("no branch>")
		n.no.Print()
	}
	if n.yes != nil {
		fmt.Println("yes branch>")
		n.yes.Print()
	}
	if n.no == nil && n.yes == nil {
		return
	}

}
