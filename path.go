/**
 * Created by WillkYang on 2017/3/17.
 */

package yeego

import (
	"os"
	"path/filepath"
	"strings"
)

var WORK_PATH string

func GetCurrentPath() string {
	if WORK_PATH != "" {
		return WORK_PATH
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		print(err.Error())
	}

	WORK_PATH = strings.Replace(dir, "\\", "/", -1)
	return WORK_PATH
}
