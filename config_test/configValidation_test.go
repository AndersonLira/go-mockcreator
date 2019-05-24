package config_test

import (
	"testing"

	"github.com/andersonlira/go-mockcreator/config"
)

func TestConfigIsDelayedMethod(t *testing.T){
	methods := []string{"aaa","bbb","ddd"}
	cfg := config.GetConfig()
	cfg.DelayMethods = methods

	if !config.IsDelayedMethod("aaa")  {
		t.Error("aaa should be delayed, but false")
	}

	if !config.IsDelayedMethod("aaa")  {
		t.Error("aaa should be delayed, but false")
	}



	if config.IsDelayedMethod("ccc")  {
		t.Error("ccc should not be delayed, but true")
	}

}