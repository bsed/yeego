/**
 * Created by angelina-zf on 17/2/27.
 */
package yeeTime

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	println(DateFormat(time.Now(), "YYYY-MM-DD HH:MM:ss"))
}
