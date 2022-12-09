package main

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import <Foundation/Foundation.h>
//#include "application.h"
//#include "window.h"
import "C"

func main() {

	title := C.CString("School 21")

	C.InitApplication()
	
	p := C.Window_Create(500, 500, 300, 200, title)
	C.Window_MakeKeyAndOrderFront(p)

	C.RunApplication()
	
}