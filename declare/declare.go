package declare

// #include <stdio.h>
// #include <stdarg.h>
// typedef void (*cb) (char*);
// void bridge(cb f, char* message) {
//	f(message);
// }
import "C"
import "unsafe"

func CallMe(pointer unsafe.Pointer) {
	C.bridge(C.cb(pointer), C.CString("enter sandman"))
}