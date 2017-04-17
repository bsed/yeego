/**
 * Created by angelina-zf on 17/2/25.
 */
package yeeFile_test

import (
	"testing"
	"github.com/yeeyuntech/yeego"
	"github.com/yeeyuntech/yeego/yeeFile"
)

var TestDir string = "data"
var TestPath string = "data/test.txt"
var TestFileName string = "test.txt"
var TestString string = "Hello!"

func TestFileGetString(t *testing.T) {
	str, err := yeeFile.FileGetString(TestPath)
	yeego.Equal(err, nil)
	yeego.Equal(str, TestString)
}

func TestFileSetString(t *testing.T) {
	yeeFile.FileSetString(TestPath, "xxx")
	str, _ := yeeFile.FileGetString(TestPath)
	yeego.Equal(str, "xxx")
	yeeFile.FileSetString(TestPath, TestString)
}
