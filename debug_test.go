/**
 * Created by angelina-zf on 17/2/25.
 */
package yeego

import (
	"testing"
)

func TestSprint(t *testing.T) {
	str := Sprint(0, 1)
	if str != "[yeegoDebug] at TestSprint() [debug_test.go:11]\n0\n1\n" {
		t.Error("TestSprint Error ")
	}
}

func TestPrint(t *testing.T) {
	//Print(0, 1)
}
