package main

import (
	"testing"
	"errors"
)

type ExecutorA struct{
	Next Executor
}

func (self *ExecutorA) Get(xml string) (string,error) {
	if self.GetNext() != nil {
		return self.GetNext().Get(xml)
	}
	return "A", errors.New("A error")
}

func (self *ExecutorA) GetNext() Executor{
	return self.Next
}


type ExecutorB struct{
	Next Executor
}

func (self ExecutorB) Get(xml string) (string,error) {
	if self.GetNext() != nil {
		return self.GetNext().Get(xml)
	}
	return "B", errors.New("B error")
}

func (self *ExecutorB) GetNext() Executor{
	return self.Next
}

type ExecutorC struct{
	Next Executor
}

func (self ExecutorC) Get(xml string) (string,error) {
	if self.GetNext() != nil {
		return self.GetNext().Get(xml)
	}
	return "C", errors.New("C error")
}

func (self *ExecutorC) GetNext() Executor{
	return self.Next
}


func TestChainOfResponsibilityBehaviour(t *testing.T){
	a := ExecutorA{}

	_,err := a.Get("")

	if err == nil {
		t.Errorf("Error should be 'A error' but nil %v",err)
	}

	b := ExecutorB{}
	c := ExecutorC{}

	a.Next = &b
	b.Next = &c
    r, err := a.Get("")

	if r != "C" {
		t.Errorf("Result should be 'C' but %s", r)
	}

	if err.Error() != "C error" {
		t.Errorf("Error should be 'C error' but %s",err.Error())
	}
}