package main

// #include <stdlib.h>

import "C"
import (
	"unsafe"
	"github.com/atrifan/io-filter/declare"
)

//export IOProxyValidate
func IOProxyValidate(
	_authToken *C.char,
	_apiKey *C.char,
	_validateAPIKey C.int,
	_validateUserToken C.int,
	_validateServiceToken C.int,
	_logCallback unsafe.Pointer,
) *C.char {
	declare.CallMe(_logCallback)
	return C.CString("OK")
}

func main() {}
