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
