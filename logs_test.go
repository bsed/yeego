/**
 * Created by angelina-zf on 17/2/25.
 */
package yeego

import (
	"testing"
)

func TestLogDebug(t *testing.T) {
	MustInitLogs("logs/", "")
	LogDebug("debug", LogFields{
		"event": "a",
		"topic": "b",
		"key":   "C",
	})
}

func TestLogError(t *testing.T) {
	LogError("error", LogFields{
		"event": "a",
		"topic": "b",
		"key":   "C",
	})
}

func TestLogInfo(t *testing.T) {
	LogInfo("info", LogFields{
		"event": "a",
		"topic": "b",
		"key":   "C",
	})
}
