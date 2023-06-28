package main

import (
	"firstProject/algorithm"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//ds, _ := plus.GetLastMonthFirstDsFormat()
	//print(ds)

	//maopao(&dataList)
	//fmt.Printf("%+v", dataList)
	//
	//dic := make(map[string]float64)
	//dic["we"] = 23.4
	//testMap(dic)
	//print(dic)
	//alist := []int{111, 2222}
	//dataList = append(dataList, alist...)
	//print(dataList)
	//GetLastMonthDayNum()
	//cImage.LoadIamgeTest()

	//sendMsg.SendFileMsg("/Users/yidezhang/Desktop/GolangProject/firstProject/sendMsg/upload.csv")
	//os.RemoveAll("./imageFile/cache")

	//fmt.Printf("%+v", cImage.OneColourType_unknown)

	//ds, error := strconv.Atoi("34e")
	//productId, error := strconv.Atoi("3")
	//fmt.Sprintf("%d,%d,%+v", ds, productId, error)
	//iMs := 20230101
	//lMs := 20230130
	//
	//for index := 0; iMs+index <= lMs; index += 5 {
	//	dayNum := iMs + index
	//	print(dayNum)
	//	print("\n")
	//}
	//
	//print(GetLastMonthDayNum())
	//fmt.Printf("%d\n", algorithm.Feibonice(20))
	//fmt.Printf("%d", algorithm.FeboniceForCacul(20))
	//algorithm.TssTree()
	algorithm.PrintTest()
}

func GetLastMonthDayNum() int {
	now := time.Now()
	firstDay := now.AddDate(0, -1, -now.Day()+1).Format("20060102")
	firstDayI, _ := strconv.Atoi(firstDay)
	lastDay := now.AddDate(0, 0, -now.Day()).Format("20060102")
	lastDayI, _ := strconv.Atoi(lastDay)
	days := lastDayI - firstDayI + 1
	fmt.Printf("上个月的天数为：%d\n", lastDayI-firstDayI+1)
	return days
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

type modelT struct {
	num int
}

func testMap(dic map[string]float64) {
	dic["we"] = 45.3
	dic["q"] = 90.1
}

var dataList = []int{
	2, 3, 9, 5, 1, 3, 4, 2, 74, 32,
}

func maopao(list *[]int) {
	length := len(*list)
	for i := 0; i < length; i++ {
		flag := true
		for j := length - 1; j > i; j-- {
			if (*list)[j] < (*list)[j-1] {
				temp := (*list)[j]
				(*list)[j] = (*list)[j-1]
				(*list)[j-1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
