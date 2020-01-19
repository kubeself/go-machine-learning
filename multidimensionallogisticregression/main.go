package main

import (
	"fmt"
	"math"
)

func main() {

	/**
	data = [[1, 1, 2, 1],
	          [1, 2, 6, 1],
	          [1, 1.5, 7, 1],
	          [1, 1, 0, 0],
	          [1, 2, 3, 0],
	          [1, 2.5, 6, 0]]
	  ds = [e[0:3] for e in data]
	  label = [e[-1] for e in data]
	*/
	data := [][]float64{
		{1, 1, 2, 1},
		{1, 2, 6, 1},
		{1, 1.5, 7, 1},
		{1, 1, 0, 0},
		{1, 2, 3, 0},
		{1, 2.5, 6, 0},
	}

	ds := make([][]float64, len(data))
	label := make([]float64, len(data))
	for i := range data {
		ds[i] = make([]float64, len(data[0])-1)
		for j, val := range data[i] {
			if j == 3 {
				label[i] = val
			} else {
				ds[i][j] = data[i][j]
			}
		}
		fmt.Println()
	}

	w := make([]float64, len(ds))
	for i := 0; i < len(data[0])-1; i++ {
		w[i] = 1
	}

	//提取label,和分析的数据
	fmt.Println(label)
	fmt.Println(ds)
	fmt.Println(w)
	//https://zhuanlan.zhihu.com/p/27161729
	//dmat:=mat.NewDense(len(ds),len(ds[0]), ds)
	//初始化w为【1，1，1】
	/**
	def reduce(ds, label):
	  #转换成矩阵
	  dmat = mat(ds)
	  lmat = mat(label).T
	  #mxn的矩阵的行数和列数
	  m,n = shape(dmat)
	  #步长
	  alpha = 0.1
	  #循环次数
	  loops = 200
	  #初始化w为[1,1,1],即分割线为 1+x+y=0
	  w = ones((n,1))
	  for i in range(loops):
	    h = sigmoid(dmat*w)
	    err = (h - lmat)
	    w = w - alpha * dmat.T* err
	  return w.A[:,0]
	也就是说w=(3.1, -5.5, 1.6), 即w0=3.1, w1=-5.5, w2 = 1.6
	分割线的表达式为：w0+w1x+w2y=0， 代入w后得 3.1-5.5x+1.6y=0, 即y=3.44x-1.9 。 见下图，该直线正确地将图形划分开。
	*/
	loops := 200
	for i := 0; i < loops; i++ {

	}

}

/**
def reduce(ds, label):
  #转换成矩阵
  dmat = mat(ds)
  lmat = mat(label).T
  #mxn的矩阵的行数和列数
  m,n = shape(dmat)
  #步长
  alpha = 0.1
  #循环次数
  loops = 200
  #初始化w为[1,1,1],即分割线为 1+x+y=0
  w = ones((n,1))
  for i in range(loops):
    h = sigmoid(dmat*w)
    err = (h - lmat)
    w = w - alpha * dmat.T* err
  return w.A[:,0]
//[ 3.1007773  -5.54393712  1.60563033]
*/

func reduce(ds [][]float64, label [][]float64) {

}

/**
  将一个数转化为概率值 位于0-1之间
*/
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(z*(-1)))
}
