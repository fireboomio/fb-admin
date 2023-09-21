package util

import (
	"github.com/labstack/gommon/log"
	"os"
	"path/filepath"
)

func GetAbsolutePath(filename string) (string, error) {
	// 获取项目根目录路径
	rootDir, err := os.Getwd()
	if err != nil {
		log.Error("获取项目根目录失败：%v", err.Error())
	}

	// 构造文件的绝对路径
	absPath := filepath.Join(rootDir, filename)
	absPath = filepath.Clean(absPath)

	return absPath, nil
}
