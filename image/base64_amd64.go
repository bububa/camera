// +build amd64

package image

import (
	base64 "github.com/libgoost/encoding-base64"
)

// EncodeToString encode bytes to base64 string
func EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
