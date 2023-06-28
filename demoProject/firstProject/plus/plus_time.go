package plus

import (
	"strconv"
	"time"
)

// GetLastMonthLastDsFormat 返回上传到obs的日期，一般是一个月最后一天
func GetLastMonthLastDsFormat() (int, error) {
	nowDay := time.Now()
	lastMonthLastDay := nowDay.AddDate(0, 0, -nowDay.Day()).Format("20060102")
	return strconv.Atoi(lastMonthLastDay)
}

// GetLastMonthFirstDsFormat 获取上一个月第一天
func GetLastMonthFirstDsFormat() (int, error) {
	nowDay := time.Now()
	lastMonthFirstDay := nowDay.AddDate(0, -1, -nowDay.Day()+1).Format("20060102")
	return strconv.Atoi(lastMonthFirstDay)
}
