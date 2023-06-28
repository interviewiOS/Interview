package algorithm

func Feibonice(lastNum int) int {
	if lastNum == 1 || lastNum == 2 {
		return 1
	}
	return Feibonice(lastNum-1) + Feibonice(lastNum-2)
}

func FeboniceForCacul(lastNum int) int {
	if lastNum == 1 || lastNum == 2 {
		return 1
	}
	start := 1
	second := 1
	lastN := 0
	for i := 3; i <= lastNum; i++ {
		lastN = start + second
		start = second
		second = lastN
	}
	return lastN
}
