package config_test

import (
	"testing"

	"github.com/andersonlira/go-mockcreator/config"
)

func TestConfigIsDelayedMethod(t *testing.T){
	methods := []string{"aaa","bbb","ddd"}
	cfg := config.GetConfig()
	cfg.DelayMethods = methods

	if !cfg.IsDelayedMethod("aaa")  {
		t.Error("aaa should be delayed, but false")
	}

	if !cfg.IsDelayedMethod("bbbc")  {
		t.Error("bbbc should be delayed, but false")
	}

	if cfg.IsDelayedMethod("dd")  {
		t.Error("dd should not be delayed, but true")
	}


	if cfg.IsDelayedMethod("ccc")  {
		t.Error("ccc should not be delayed, but true")
	}

}