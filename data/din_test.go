package data

import (
	"log"
	"testing"
)

func TestGetAreaData(t *testing.T) {
	data := GetAreaData("陕西省")
	log.Println(data)
}