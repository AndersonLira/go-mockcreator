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

func TestShouldClearCache(t *testing.T){
	methods := make(map[string][]string)
	cfg := config.GetConfig()
	methods["aaa"] = []string{"aaa","bbb","ddd"}
	cfg.ClearCache = methods

	if list,ok := cfg.ShouldClearCache("aaa"); !ok  || len(list) != 3  {
		t.Errorf("aaa should have 3 items, but %v",list )
	}

	if list,ok := cfg.ShouldClearCache("bbb"); ok  || len(list) != 0 {
		t.Errorf("bbb should not exists, but %v",list )
	}
}

func TestConfigIsCacheEvict(t *testing.T){
	methods := []string{"aaa","bbb","ddd"}
	cfg := config.GetConfig()
	cfg.CacheEvict = methods

	if !cfg.IsCacheEvict("aaa.xml")  {
		t.Error("aaa.xml should be cache evict, but false")
	}

	if !cfg.IsCacheEvict("bbbc")  {
		t.Error("bbbc should be cache evict, but false")
	}

	if cfg.IsCacheEvict("dd")  {
		t.Error("dd should not be cache evict, but true")
	}


	if cfg.IsCacheEvict("ccc")  {
		t.Error("ccc should not be cached evict, but true")
	}
}