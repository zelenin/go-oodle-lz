package oodle

import (
	"golang.org/x/sys/windows"
)

var meta = libMeta{
	url:    "https://cdn.unrealengine.com/dependencies/UnrealEngine-27563807/51bf6515dd35ac8361c9a324b6deb1736a61240c",
	hash:   "b5120a64e2756eb99978a9ddc4fb36939a3d4cfe",
	offset: 1240856,
	size:   637952,
	name:   "oo2core_9_win64.dll",
}

func loadLib() (uintptr, error) {
	libOnce.Do(func() {
		libPath, err := resolveLibPath()
		if err != nil {
			libOnce.err = err
			return
		}

		handle, err := windows.LoadLibrary(libPath)
		if err != nil {
			libOnce.err = err
			return
		}

		libOnce.handle = uintptr(handle)

		registerLibFuncs(libOnce.handle)
	})

	return libOnce.handle, libOnce.err
}
