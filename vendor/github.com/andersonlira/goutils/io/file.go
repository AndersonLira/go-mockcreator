package io

import (
	"fmt"
	"io/ioutil"
	"os"
)

//WriteFile writes a text file with content
//It returns error if some trouble happens
func WriteFile(fileName string, content string) error {
	b := []byte(content)
	err := ioutil.WriteFile(fileName, b, 0644)
	return err
}

//WriteFilef writes on file formatting content
func WriteFilef(fileName string, content string, args ...interface{}) error {
	s := fmt.Sprintf(content, args...)
	return WriteFile(fileName, s)
}

//ReadFile reads a file and return its content as string
func ReadFile(fileName string) (content string, err error) {
	b, err := ioutil.ReadFile(fileName)
	return string(b), err
}

//AppendFile writes on file increasing its content
func AppendFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

//AppendFilef appends text on file formatting content
func AppendFilef(fileName string, content string, args ...interface{}) error {
	s := fmt.Sprintf(content, args...)
	return AppendFile(fileName, s)
}
