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
	"strings"
	"unsafe"
)

func RunString(command string) int {
	commandName := C.CString(command)
	defer C.free(unsafe.Pointer(commandName))

	commandReturn := C.cystem(commandName)
	return int(commandReturn)
}

func RunStringSlice(command []string) int {
	commandName := C.CString(strings.Join(command, " "))
	defer C.free(unsafe.Pointer(commandName))

	commandReturn := C.cystem(commandName)
	return int(commandReturn)
}
