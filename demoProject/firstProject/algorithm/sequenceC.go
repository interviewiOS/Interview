package algorithm

import (
	"fmt"
)

func PrintTest() {
	var oriList = []int{
		-11, 4, 3, 3, 5, 8, 80, 2, 13, 15, 26, 100, 93, -1, -2,
	}
	//mappao(oriList)
	//xuanze(oriList)
	//kuaipai(oriList)
	//charu(oriList)
	//shellPaixu(oriList)
	fmt.Printf("\n长度 %d", len(oriList))
	fmt.Printf("\n%+v", oriList)

	resultList := mergePaixu(oriList)
	fmt.Printf("\n%+v", resultList)
	fmt.Printf("\n长度 %d", len(resultList))

}

func shellPaixu(oriList []int) {
	if len(oriList) <= 1 {
		return
	}
	length := len(oriList)
	gap := length / 2
	for gap > 0 {
		i := gap
		for ; i < length; i++ {
			fmt.Printf("i:%d ", i)
			temp := oriList[i]
			j := i - gap
			for j >= 0 && oriList[j] > temp {
				oriList[j+gap] = oriList[j]
				j -= gap
			}
			oriList[j+gap] = temp
		}

		gap = gap / 2
	}
}

func mergePaixu(oriList []int) []int {
	length := len(oriList)
	if length <= 1 {
		return oriList
	}
	half := length / 2
	leftList := oriList[0:half]
	rightList := oriList[half:]
	return subMergePaixu(mergePaixu(leftList), mergePaixu(rightList))
}
func subMergePaixu(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func charu(oriList []int) {
	length := len(oriList)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		cNum := oriList[i]
		j := i - 1
		for j >= 0 && oriList[j] > cNum {
			oriList[j+1] = oriList[j]
			j--
		}
		oriList[j+1] = cNum
	}
}

func kuaipai(oriList []int) {
	subKuaiPai(oriList, 0, len(oriList)-1)
}
func subKuaiPai(oriList []int, start, end int) {
	if start >= end {
		return
	}
	length := len(oriList)
	if end >= length {
		return
	}
	cNum := oriList[start]
	pre := start
	post := end
	for pre < post {
		for oriList[post] <= cNum && pre < post {
			post--
		}
		oriList[pre] = oriList[post]

		for oriList[pre] >= cNum && pre < post {
			pre++
		}
		oriList[post] = oriList[pre]
	}
	oriList[post] = cNum

	subKuaiPai(oriList, start, pre-1)
	subKuaiPai(oriList, pre+1, end)
}

func mappao(oriList []int) []int {
	tList := oriList
	length := len(tList)
	if length <= 1 {
		return tList
	}
	for i := 0; i < length; i++ {
		for j := length - 1; i < j; j-- {
			if tList[j] > tList[j-1] {
				var temp = tList[j]
				tList[j] = tList[j-1]
				tList[j-1] = temp
			}
		}
	}
	return tList
}

func xuanze(oriList []int) []int {
	tList := oriList
	length := len(tList)
	if length <= 1 {
		return tList
	}
	for i := 0; i < length; i++ {
		sIndex := i
		for j := length - 1; i < j; j-- {
			if tList[j] < tList[sIndex] {
				sIndex = j
			}
		}
		if i != sIndex {
			var temp = tList[sIndex]
			tList[sIndex] = tList[i]
			tList[i] = temp
		}
	}
	return tList
}
