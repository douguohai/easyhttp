package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//获取当前执行文件的父目录的绝对路径
func RunAbsPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	return path
}
