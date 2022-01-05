package data

import (
	"os"
	"strconv"
	"testing"
	"time"
)

//创建文件
func Test1(t *testing.T) {
	create, err := os.Create("test.png")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(create.Name())
	}
}

// 删除文件
func Test2(t *testing.T) {
	err := os.Remove("test.png")
	if err != nil {
		t.Log(err)
	} else {
		t.Log("删除成功")
	}
}

// 获取时间戳
func Test3(t *testing.T) {
	t1 := strconv.FormatInt(time.Now().Unix(), 10)
	t.Log(t1)
}

//
func Test4(t *testing.T) {
	var s []string
	s = append(s, "xx")
	s = append(s, "aa")
	for _, fi := range s {
		t.Log(fi)
	}
}
