/**
 * Created by angelina-zf on 17/2/25.
 */

// file 文件处理相关函数
package yeeFile

import (
	"time"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"io"
	"path/filepath"
)

// FileGetBytes 通过给定的文件名称或者url地址以及超时时间获取文件的[]byte数据.
func FileGetBytes(filenameOrURL string, timeout ...time.Duration) ([]byte, error) {
	if strings.Contains(filenameOrURL, "://") {
		if strings.Index(filenameOrURL, "file://") == 0 {
			filenameOrURL = filenameOrURL[len("file://"):]
		} else {
			client := http.DefaultClient
			if len(timeout) > 0 {
				client = &http.Client{Timeout: timeout[0]}
			}
			r, err := client.Get(filenameOrURL)
			if err != nil {
				return nil, err
			}
			defer r.Body.Close()
			if r.StatusCode < 200 || r.StatusCode > 299 {
				return nil, fmt.Errorf("%d: %s", r.StatusCode, http.StatusText(r.StatusCode))
			}
			return ioutil.ReadAll(r.Body)
		}
	}
	return ioutil.ReadFile(filenameOrURL)
}

// FileSetBytes 向指定的文件设置[]byte内容.
func FileSetBytes(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0660)
}

// FileAppendBytes 向指定的文件追加[]byte内容.
func FileAppendBytes(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// FileGetString 通过给定的文件名称或者url地址以及超时时间获取文件的string数据.
func FileGetString(filenameOrURL string, timeout ...time.Duration) (string, error) {
	bytes, err := FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FileSetString 向指定的文件设置string内容.
func FileSetString(filename string, data string) error {
	return FileSetBytes(filename, []byte(data))
}

// FileAppendString 向指定的文件追加string内容.
func FileAppendString(filename string, data string) error {
	return FileAppendBytes(filename, []byte(data))
}

// FileExists 文件或者文件夹是否存在.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// FileTimeModified 返回文件的最后修改时间
// 如果有错误则返回空time.Time.
func FileTimeModified(filename string) time.Time {
	info, err := os.Stat(filename)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

// FileIsDir 判断是否是文件夹.
func FileIsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	return err == nil && info.IsDir()
}

// FileFind 在给定的文件夹中查找某文件.
func FileFind(searchDirs []string, filenames ...string) (filePath string, found bool) {
	for _, dir := range searchDirs {
		for _, filename := range filenames {
			filePath = path.Join(dir, filename)
			if FileExists(filePath) {
				return filePath, true
			}
		}
	}
	return "", false
}

// FileGetPrefix 获取文件的前缀.
func FileGetPrefix(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[0:i]
		}
	}
	return ""
}

// FileGetExt 获取文件的后缀.
func FileGetExt(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i + 1:]
		}
	}
	return ""
}

// FileCopy 将文件从原地址拷贝到目的地.
func FileCopy(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		si, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, si.Mode())
		}
	}
	return err
}

// DirSize 返回文件夹的大小
func DirSize(path string) int64 {
	var dirSize int64 = 0
	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
		}
		return nil
	}
	filepath.Walk(path, readSize)
	return dirSize
}
