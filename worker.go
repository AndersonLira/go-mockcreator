package main

import (
	"time"
	"github.com/andersonlira/go-mockcreator/chain"
)

type Worker struct {
	Executor chain.Executor
	Request string
	Sleep int32
}

func (w Worker)  Run(){
	go w.proccess()
}

func (w Worker) proccess(){
	time.Sleep(time.Duration(w.Sleep) * time.Second)
	_, err := w.Executor.Get(w.Request)
	if err != nil {
		w.Run()
	}
}
