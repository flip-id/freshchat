package freshchat

import "unsafe"

func convertByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
