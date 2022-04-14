// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cystem

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "cystem.h"
import "C"
import (
	"unsafe"
)

func Run(command string) int {
	commandName := C.CString(command)
	defer C.free(unsafe.Pointer(commandName))

	commandReturn := C.cystem(commandName)
	return int(commandReturn)
}
