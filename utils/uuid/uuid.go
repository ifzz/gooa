package uuid

import (
	"crypto/rand"
	"fmt"

	"github.com/iissy/gooa/utils/base58"
)

func genUUIDv4() []byte {
	var u = make([]byte, 16)
	_, _ = rand.Read(u[:])

	// Set version
	u[6] = (u[6] & 0x0F) | (4 << 4)
	// Set variant bits
	u[8] = (u[8] | 0x40) & 0x7F
	return u
}

// 生成 RFC4122 V4 版本的 UUID，长度为 36
func V4() string {
	u := genUUIDv4()
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// 生成 Base58 表示的 RFC4122 V4 版本的 UUID，长度为 22
func V4Compressed() string {
	return base58.Encode(genUUIDv4())
}

func DecompressV4(compressed string) string {
	u := base58.Decode(compressed)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
