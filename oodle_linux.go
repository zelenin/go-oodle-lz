package oodle

import "github.com/ebitengine/purego"

var meta = libMeta{
	url:    "https://cdn.unrealengine.com/dependencies/UnrealEngine-38305390/1b211147634301e52fad51218b52f4db781514b5",
	hash:   "d5b19257d7459b6be53fbae740a2731746b0c251",
	offset: 1397620,
	size:   686376,
	name:   "liboo2corelinux64.so.9",
}

func loadLib() (uintptr, error) {
	libOnce.Do(func() {

		libPath, err := resolveLibPath()
		if err != nil {
			libOnce.err = err
			return
		}

		handle, err := purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			libOnce.err = err
			return
		}

		libOnce.handle = handle

		registerLibFuncs(libOnce.handle)
	})

	return libOnce.handle, libOnce.err
}
