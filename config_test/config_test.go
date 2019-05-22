package config_test

import (
	"sync"
	"testing"

	"github.com/andersonlira/go-mockcreator/config"
)

func TestConfigSingleton(t *testing.T){
	wg := sync.WaitGroup{}
	ca := config.GetConfig()
	cb := config.GetConfig()

	if ca  != cb {
		t.Errorf("Objects should be same but %p and %p", ca, cb)
	}

	size := 100
	wg.Add(size)
	many := make([]*config.Config, size)

	for i := 0; i < size ;i++ {
		go func(index int) {
			defer wg.Done()
			many[index] = config.GetConfig()
		}(i)
	}

	wg.Wait()

	for i:= 0; i< size;i++ {
		if many[i] != ca {
			t.Errorf("Objects should be same but %p and %p. Index: %d", ca, many[i],i)
		}
	}

}