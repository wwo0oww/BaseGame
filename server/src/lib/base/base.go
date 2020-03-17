package base

import (
	"os"
	"runtime"
)


func EnsureDir(path string) error{
	Exist,err := PathExists(path)
	if err != nil{
		return nil
	}
	if !Exist{
		os.Mkdir(path,os.ModePerm)
	}
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetLineBreak() string{
	switch runtime.GOOS{
	case `windows`:
		return "\r\n"
	default:
		return "\n"
	}
}