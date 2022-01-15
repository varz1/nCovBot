package data

import (
	"io/ioutil"
	"log"
	"os"
)

func YaHeiFontData() []byte {
	p,_ := os.Getwd()
	fontBytes, err := ioutil.ReadFile(p+".fonts/WeiRuanYaHei-1.ttf")
	if err != nil {
		log.Println(err)
		return nil
	}
	return fontBytes
}
