package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var fl []string

func main() {
	wd, _ := os.Getwd()
	// 读取所有文件/文件夹、子文件/子文件夹
	AllLayer(wd)
	// 读取当前路径的所有文件/文件夹
	CurrentLayer(wd)
}

// 使用go内置包filepath中的walk函数
// 会完全遍历出指定目录下的所有文件/文件夹，及所有子目录中的子文件/子文件夹
// 可以在回调函数allFiles中对walk()传入的fileinfo进行过滤处理
func AllLayer(path string) {
	filepath.Walk(path, allFiles)
}

// 此处过滤掉所有文件夹，只取文件并存入到切片f1中
func allFiles(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		fl = append(fl, path)
	}
	return nil
}

// 读取给定路径的当前目录中的所有文件及文件夹
// 此处对文件夹进行过滤，只打印路径+文件名
func CurrentLayer(path string) error {
	rd, err := ioutil.ReadDir(path)
	for _, f := range rd {
		if !f.IsDir() {
			fmt.Println(filepath.Join(path, f.Name()))
		}
	}
	return err
}
