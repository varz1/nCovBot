package data

import (
	"fmt"
	"github.com/varz1/nCovBot/model"
	"github.com/varz1/nCovBot/variables"
	"testing"
)

func Test1(t *testing.T) {
	var tencent model.Overall
	resp, _ := request.R().SetResult(&tencent).Get(variables.LOCAL)
	if resp.IsSuccess() {
		fmt.Println(tencent)
	}
}
