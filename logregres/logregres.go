package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	//data,label:= LoadDataSet()
	//fmt.Println(data)
	//fmt.Println(label)

	v := []float64{1,2,3,4,5,6,7,8,9,10,11,12}
	A := mat.NewDense(3, 4, v)
	MatPrint(A)

	
}

func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func LoadDataSet() ([][]float64, []float64) {
	absPath, _ := filepath.Abs("logregres/testSet.txt")
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	dataMat := make([][]float64, 0)
	labelMat := make([]float64, 0)
	for scanner.Scan() {
		lineArr := strings.Split(scanner.Text(),"\t")
		fmt.Println(lineArr)
		lineData := []float64{1.0}
		for i,v:=range lineArr {
			if i!=len(lineArr)-1 {
				val , _ := strconv.ParseFloat(v, 64)
				lineData = append(lineData, val)
			}  else {
				label, _:= strconv.ParseFloat(v, 64)
				labelMat = append(labelMat, label)
			}
		}
		dataMat = append(dataMat, lineData)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dataMat , labelMat
}

func Sigmoid(x float64) float64 {
	return 1.0/(1+math.Exp(-1*x))
}

func GradAscent(data [][]float64, labels []float64) {
	m:=len(data)
	n:=len(data[0])
	lineData := make([]float64, m*n)

	for _,v:=range data {
		for _,w:=range v {
			lineData = append(lineData, w)
		}
	}

	dataMatrix := mat.NewDense(m, n, lineData)
	labelMat := mat.NewDense(1, len(labels), labels).T()

	alpha := 0.001
	maxCycle := 500

	weights := make([]float64, n)
	for i:=0;i<n;i++{
		weights = append(weights,1.0)
	}
	weightsMat := mat.NewDense(n,1, weights)

	for i:=0;i<maxCycle;i++{

		var c mat.Dense
		c.Mul(dataMatrix, weightsMat)
		h:=Sigmoid(c.Trace())

		err := labelMat. - h


		//d.Product(dataMatrix, weightsMat)


	}









}


