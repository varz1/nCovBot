package data

import (
	"log"
	"testing"
)

func TestGetAreaData(t *testing.T) {
	var s []int = make([]int, 4)
	s = append(s[0:1], 1)
	s = append(s[1:2], 2)
	s = append(s[2:3], 3)
	log.Println(s)
}
