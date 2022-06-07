package data

import (
	"github.com/varz1/nCovBot/model"
	"github.com/varz1/nCovBot/variables"
	"testing"
)

func TestGetRisk(t *testing.T) {
	var res struct {
		Data []model.RiskArea `json:"data"`
	}
	resp, err := request.R().SetResult(&res).Get(variables.RISK)
	if err != nil || resp.StatusCode() != 200 {
		t.Log(err)
		return
	}
	t.Log(res)
}

func Test1(t *testing.T) {
	t.Log("test")
}
