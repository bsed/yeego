/**
 * Created by angelina on 2017/4/17.
 */

package yeeStrings

import (
	"testing"
	"github.com/yeeyuntech/yeego"
)

func TestIsInSlice(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	str1 := "a"
	str2 := "d"
	yeego.OK(IsInSlice(testSlice, str1))
	yeego.OK(!IsInSlice(testSlice, str2))
}

func TestMapFunc(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	f := func(a string) string {
		return a + "?"
	}
	newSlice := MapFunc(testSlice, f)
	yeego.Equal(newSlice[0], "a?")
	yeego.Equal(newSlice[1], "b?")
	yeego.Equal(newSlice[2], "c?")
}

func TestAddURLParam(t *testing.T) {
	old := "www.baidu.com"
	old = AddURLParam(old, "a", "b")
	yeego.Equal(old, "www.baidu.com?a=b")
	old = AddURLParam(old, "c", "d")
	yeego.Equal(old, "www.baidu.com?a=b&c=d")
}

func TestStringToIntArray(t *testing.T) {
	str1 := "1,2,3"
	yeego.Equal(len(StringToIntArray(str1, ",")), 3)
	str2 := "aa,aa"
	yeego.Equal(len(StringToIntArray(str2, ",")), 0)
}
