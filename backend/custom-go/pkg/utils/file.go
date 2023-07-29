package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type fileCache struct {
	info    os.FileInfo
	content []byte
}

var fileCacheMap map[string]*fileCache

func ReadBytesAndCacheFile(path string) ([]byte, error) {
	if nil == fileCacheMap {
		fileCacheMap = make(map[string]*fileCache)
	}

	fileInfo, err := os.Stat(path)
	if nil != err {
		return nil, err
	}

	cache := fileCacheMap[path]
	if nil == cache {
		content, err := os.ReadFile(path)
		if nil != err {
			return nil, err
		}

		cache = &fileCache{fileInfo, content}
		fileCacheMap[path] = cache
	} else if !reflect.DeepEqual(cache.info.ModTime(), fileInfo.ModTime()) {
		content, err := os.ReadFile(path)
		if nil != err {
			return nil, err
		}

		cache.info = fileInfo
		cache.content = content
	}
	return cache.content, nil
}

func ReadStructAndCacheFile(path string, result interface{}) error {
	bytesData, err := ReadBytesAndCacheFile(path)
	if nil != err {
		return err
	}

	return json.Unmarshal(bytesData, &result)
}

func JoinString(sep string, str ...string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.Join(str, sep)
}

func JoinPathAndToSlash(path ...string) string {
	return filepath.ToSlash(filepath.Join(path...))
}

func ConvertType[S, T any](s *S) (t *T) {
	sBytes, _ := json.Marshal(s)
	_ = json.Unmarshal(sBytes, &t)
	return
}
