package main

import "fmt"

/**
  处理多类别问题
  用概率高的作为对应的类别
*/
func main() {
	listOPosts, listClasses := LoadDataSet()
	myVocabList := CreateVocabList(listOPosts)

	fmt.Println(myVocabList)
	fmt.Println(listClasses)
	fmt.Println(SetOfWords2Vec(myVocabList, listOPosts[0]))
	fmt.Println(SetOfWords2Vec(myVocabList, listOPosts[3]))

	//准备矩阵, 包含是否包含在词库的信息, 0=不包含, 1=包含
	trainMat := make([][]int, 0)
	for _, doc := range listOPosts {
		trainMat = append(trainMat, SetOfWords2Vec(myVocabList, doc))
	}

	//训练数据
	TrainNBO(trainMat, listClasses)

}

func LoadDataSet() ([][]string, []int) {
	postingList := [][]string{
		{"my", "dog", "has", "flea", "problems", "help", "please"},
		{"maybe", "not", "take", "him", "to", "dog", "park", "stupid"},
		{"my", "dalmation", "is", "so", "cute", "I", "love", "him"},
		{"stop", "posting", "stupid", "worthless", "garbage"},
		{"mr", "licks", "ate", "my", "steak", "how", "to", "stop", "him"},
		{"quit", "buying", "worthless", "dog", "food", "stupid"},
	}

	classVec := []int{0, 1, 0, 1, 0, 1} // 0 代表正常, 1代表侮辱性
	return postingList, classVec
}

/**
获得所有词汇列表
*/
func CreateVocabList(dataSet [][]string) []string {
	vocMap := make(map[string]int)
	resList := make([]string, 0)
	for _, v := range dataSet {
		for _, w := range v {
			if _, ok := vocMap[w]; ok == false {
				vocMap[w] = 1
				resList = append(resList, w)
			}
		}
	}
	return resList
}

/**
  将要检测的单词集, 存在于词汇表的位置都标记成1
*/
func SetOfWords2Vec(vocabList []string, inputSet []string) []int {
	//创建一个圈0 的向量
	resVec := make([]int, len(vocabList))
	for _, v := range inputSet {
		for j, w := range vocabList {
			if w == v {
				resVec[j] = 1
				break
			}
		}
	}
	return resVec
}

/**
  训练数据
*/
func TrainNBO(trainMatrix [][]int, trainCategory []int) {
	numTrainDoc := len(trainMatrix) // 行数
	numWords := len(trainMatrix[0]) // 列数
	sumTrainCategory := float64(0)

	for _, v := range trainCategory {
		sumTrainCategory += float64(v)
	}
	pAbusive := sumTrainCategory / float64(numTrainDoc) // 总的成功概率

	//初始化概率
	p0num := make([]int, numWords) // 0 的概率
	p1num := make([]int, numWords) // 1 的概率

	p0Denom := float64(0) // 0 的分母
	p1Denom := float64(0) // 1 的分母

	// 遍历每个文档
	for i := 0; i < numTrainDoc; i++ {
		if trainCategory[i] == 1 {
			// 当前行是1
			p1num
		}
	}

}
