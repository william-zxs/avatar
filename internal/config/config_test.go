package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	res := ReadConfig("../../conf.yaml")
	fmt.Printf("%+v", res)
}
