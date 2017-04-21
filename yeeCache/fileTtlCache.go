/**
 * Created by angelina-zf on 17/2/27.
 */

// yeeCache
// 文件缓存类
package yeeCache

import (
	"time"
	"github.com/yeeyuntech/yeego/yeeFile"
	"encoding/json"
)

type fileTtlCache struct {
	Value   []byte
	Timeout time.Time
}

// FileTtlCache
// 文件ttl缓存
func FileTtlCache(fileName string, f func() (b []byte, ttl time.Duration, err error)) (b []byte, err error) {
	cache := &fileTtlCache{}
	now := time.Now()
	fileContent, err := yeeFile.GetBytes(fileName)
	// 获取到了内容
	if err == nil {
		// 成功解析为cache
		if err := json.Unmarshal(fileContent, &cache); err == nil {
			// 超时时间在现在后面
			if cache.Timeout.After(now) {
				return cache.Value, nil
			}
		}
	}
	b, ttl, err := f()
	if err != nil {
		return nil, err
	}
	cache.Value = b
	cache.Timeout = now.Add(ttl)
	if err := yeeFile.MkdirForFile(fileName); err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(&cache)
	if err != nil {
		return nil, err
	}
	if err := yeeFile.SetBytes(fileName, jsonData); err != nil {
		return nil, err
	}
	return b, nil
}
