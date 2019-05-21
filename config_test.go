package main

import (
	"sync"
	"testing"
)

func TestConfigSingleton(t *testing.T){
	wg := sync.WaitGroup{}
	ca := GetConfig()
	cb := GetConfig()

	if ca  != cb {
		t.Errorf("Objects should be same but %p and %p", ca, cb)
	}

	size := 100
	wg.Add(size)
	many := make([]*Config, size)

	for i := 0; i < size ;i++ {
		go func(index int) {
			defer wg.Done()
			many[index] = GetConfig()
		}(i)
	}

	wg.Wait()

	for i:= 0; i< size;i++ {
		if many[i] != ca {
			t.Errorf("Objects should be same but %p and %p. Index: %d", ca, many[i],i)
		}
	}

}