/**
 * Created by angelina-zf on 17/3/13.
 */
package aes

import (
	"testing"
	"fmt"
	"github.com/yeeyuntech/yeego"
)

func TestAesDecrypt(t *testing.T) {
	for _, origin := range [][]byte{
		[]byte(""),
		[]byte("1"),
		[]byte("12"),
		[]byte("123"),
		[]byte("1234"),
		[]byte("12345"),
		[]byte("123456"),
		[]byte("1234567"),
		[]byte("12345678"),
		[]byte("123456789"),
		[]byte("1234567890"),
		[]byte("123456789012345"),
		[]byte("1234567890123456"),
		[]byte("12345678901234567"),
	} {
		ob, err1 := AesEncrypt([]byte("1"), origin)
		yeego.Equal(err1, nil)
		ret, err2 := AesDecrypt([]byte("1"), ob)
		yeego.Equal(err2, nil)
		yeego.Equal(ret, origin)
	}
}

func TestAesEncrypt(t *testing.T) {
	data, err := AesEncrypt([]byte("test"), []byte("hello"))
	yeego.Equal(err, nil)
	fmt.Println(string(data))
}

