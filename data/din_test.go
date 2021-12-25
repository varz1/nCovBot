package data

import (
	"testing"
)

func TestGetAreaData(t *testing.T) {
	high := GetRiskLevel("2")
	mid := GetRiskLevel("1")
	t.Logf("高风险地区数量:%d\n中风险地区数量:%d", len(high), len(mid))
}
