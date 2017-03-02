/**
 * Created by angelina-zf on 17/2/28.
 */
package yeeCache

import (
	"testing"
	"github.com/agelinazf/yee/yeeTest"
	"time"
)

func TestNewMemoryCache(t *testing.T) {
	mc := NewMemoryCache(DefaultExpiration, 0)
	a, found := mc.Get("a")
	yeeTest.Equal(found, false)
	yeeTest.Equal(a, nil)
	b, found := mc.Get("b")
	yeeTest.Equal(found, false)
	yeeTest.Equal(b, nil)
	c, found := mc.Get("c")
	yeeTest.Equal(found, false)
	yeeTest.Equal(c, nil)
	mc.Set("a", 1, DefaultExpiration)
	mc.Set("b", "b", DefaultExpiration)
	mc.Set("c", 3.5, DefaultExpiration)
	x, found := mc.Get("a")
	yeeTest.Equal(found, true)
	yeeTest.Equal(x, 1)
	y, found := mc.Get("b")
	yeeTest.Equal(found, true)
	yeeTest.Equal(y, "b")
	z, found := mc.Get("c")
	yeeTest.Equal(found, true)
	yeeTest.Equal(z, 3.5)
}

func TestMemoryCacheTime(t *testing.T) {
	var found bool
	mc := NewMemoryCache(40*time.Millisecond, time.Millisecond)
	mc.Set("a", 1, DefaultExpiration)
	mc.Set("b", 2, NoExpiration)
	mc.Set("c", 3, 30*time.Millisecond)
	mc.Set("d", 4, 50*time.Millisecond)
	<-time.After(30 * time.Millisecond)
	_, found = mc.Get("c")
	yeeTest.Equal(found, false)
	<-time.After(15 * time.Millisecond)
	_, found = mc.Get("a")
	yeeTest.Equal(found, false)
	<-time.After(15 * time.Millisecond)
	_, found = mc.Get("d")
	yeeTest.Equal(found, false)
	_, found = mc.Get("b")
	yeeTest.Equal(found, true)
}
