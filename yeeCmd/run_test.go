/**
 * Created by angelina on 2017/5/2.
 */

package yeeCmd

import (
	"testing"
	"github.com/yeeyuntech/yeego"
)

func TestRun(t *testing.T) {
	yeego.Equal(Run("ls -al"), nil)
}

func TestRunSlice(t *testing.T) {
	yeego.Equal(RunSlice([]string{"ls", "-al"}), nil)
}

func TestWhich(t *testing.T) {
	yeego.OK(Which("ls"))
	yeego.OK(!Which("xxx"))
}
