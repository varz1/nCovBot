package data

import (
	"github.com/varz1/nCovBot/model"
	"testing"
)

func TestGetRisk(t *testing.T) {
	t.Log("fake")
	var riskdata model.Risks
	risk, _ := C.Get("risk")
	riskdata = risk.(model.Risks)
	t.Logf("high %v", len(riskdata.High))
	t.Logf("mid %v", len(riskdata.Mid))
	t.Log(riskdata)
}

func Test1(t *testing.T) {
	t.Log("test")
}
