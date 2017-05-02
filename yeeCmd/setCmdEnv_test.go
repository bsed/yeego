/**
 * Created by angelina on 2017/5/2.
 */

package yeeCmd_test

import (
	"testing"
	"github.com/yeeyuntech/yeego/yeeCmd"
	"github.com/yeeyuntech/yeego"
)

func TestSetCmdEnv(t *testing.T) {
	cmd := yeeCmd.CmdBash("echo $GOPATH")
	yeeCmd.SetCmdEnv(cmd.GetExecCmd(), "GOPATH", "testGOPATH")
	b, err := cmd.RunAndReturnOutput()
	yeego.Equal(err, nil)
	yeego.Equal(string(b), "testGOPATH\n")
}
