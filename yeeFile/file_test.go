/**
 * Created by angelina-zf on 17/2/25.
 */
package yeeFile

import (
	"testing"
	"yeego"
)

var TestDir string = "data"
var TestPath string = "data/test.txt"
var TestFileName string = "test.txt"
var TestString string = "Hello!"

func TestFileGetString(t *testing.T) {
	str, err := FileGetString(TestPath)
	yeego.Equal(err, nil)
	yeego.Equal(str, TestString)
}

func TestFileSetString(t *testing.T) {
	FileSetString(TestPath, "xxx")
	str, _ := FileGetString(TestPath)
	yeego.Equal(str, "xxx")
	FileSetString(TestPath, TestString)
}
