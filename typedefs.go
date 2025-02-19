package oodle

import "unsafe"

type DecompressCallback func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer, int, int, int)
