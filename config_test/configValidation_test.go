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

func TestStaticReturn(t *testing.T){
	methods := make(map[string]string)
	methods["aaa"] = "aaa_static"
	methods["aaad"] = "aaad_static"
	methods["ccc"] = "c_static"

	cfg := config.GetConfig()
	cfg.StaticReturn = methods

	if v,ok := cfg.IsStaticReturn("aaa"); !ok  || v != "aaa_static"  {
		t.Errorf("aaa should have aaa_static value, but %s",v )
	}

	if v,ok := cfg.IsStaticReturn("aaad"); !ok  || v != "aaa_static"  {
		t.Errorf("aaad should have aaa_static value, but %s",v )
	}

	if v,ok := cfg.IsStaticReturn("cccccc.xml"); !ok  || v != "c_static"  {
		t.Errorf("cccccc should have c_static value, but %s",v )
	}


	if v,ok := cfg.IsStaticReturn("bbb"); ok  {
		t.Errorf("bbb should not have static value, but %s",v )
	}
}