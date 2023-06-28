package plus

import "testing"

func TestGetLastMonthLastDsFormat(t *testing.T) {
	if ans, error := GetLastMonthLastDsFormat(); ans != 202301 && error != nil {
		t.Errorf("获得上月的数据有问题")
	}
}
