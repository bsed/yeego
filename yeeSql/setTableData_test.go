/**
 * Created by angelina on 2017/4/16.
 */

package yeeSql_test

import (
	"testing"
	"yeego/yeeSql"
)

func TestMustSetTableDataToml(t *testing.T) {
	initTestDbTable()
	yeeSql.MustSetTableDataToml(tomlData)
}
