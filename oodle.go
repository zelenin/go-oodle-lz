package oodle

import (
	"compress/gzip"
	"fmt"
	"github.com/ebitengine/purego"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type libMeta struct {
	url    string
	hash   string
	offset int64
	size   int64
	name   string
}

var libOnce struct {
	sync.Once
	handle uintptr
	err    error
}

func init() {
	var err error

	var downloaded bool

	if !isLibExists() {
		log.Printf("oodle lib is not exists, downloading...")
		err = download()
		if err != nil {
			log.Fatalf("oodle lib download: %s", err)
		}
		downloaded = true
	}

	_, err = loadLib()
	if err != nil {
		log.Fatalf("oodle lib load: %s", err)
	}

	if downloaded {
		var configValues ConfigValues
		GetConfigValues(&configValues)

		log.Printf("oodle, v2.%d.%d", (configValues.HeaderVersion>>16)&0xff, (configValues.HeaderVersion>>8)&0xff)
	}
}

var possibleLibPaths = []string{
	meta.name,
	getTempDllPath(),
}

func isLibExists() bool {
	_, err := resolveLibPath()
	return err == nil
}

func resolveLibPath() (string, error) {
	for _, libPath := range possibleLibPaths {
		_, err := os.Stat(libPath)
		if !os.IsNotExist(err) {
			return libPath, nil
		}
	}

	return "", fmt.Errorf("`%s` is not resolve", meta.name)
}

func getTempDllPath() string {
	return filepath.Join(os.TempDir(), "go-oodle-lz", meta.name)
}

func download() error {
	req, err := http.NewRequest("GET", meta.url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gzr.Close()

	_, err = io.CopyN(io.Discard, gzr, meta.offset)
	if err != nil {
		return err
	}

	lr := io.LimitReader(gzr, meta.size)

	filePath := getTempDllPath()
	err = os.MkdirAll(filepath.Dir(filePath), 0777)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, lr)
	if err != nil {
		return err
	}

	return nil
}

func registerLibFuncs(libHandle uintptr) {
	purego.RegisterLibFunc(&compress, libHandle, "OodleLZ_Compress")
	purego.RegisterLibFunc(&decompress, libHandle, "OodleLZ_Decompress")
	purego.RegisterLibFunc(&decoderCreate, libHandle, "OodleLZDecoder_Create")
	purego.RegisterLibFunc(&decoderMemorySizeNeeded, libHandle, "OodleLZDecoder_MemorySizeNeeded")
	purego.RegisterLibFunc(&threadPhasedBlockDecoderMemorySizeNeeded, libHandle, "OodleLZ_ThreadPhased_BlockDecoderMemorySizeNeeded")
	purego.RegisterLibFunc(&decoderDestroy, libHandle, "OodleLZDecoder_Destroy")
	purego.RegisterLibFunc(&decoderReset, libHandle, "OodleLZDecoder_Reset")
	purego.RegisterLibFunc(&decoderDecodeSome, libHandle, "OodleLZDecoder_DecodeSome")
	purego.RegisterLibFunc(&makeValidCircularWindowSize, libHandle, "OodleLZDecoder_MakeValidCircularWindowSize")
	purego.RegisterLibFunc(&makeSeekChunkLen, libHandle, "OodleLZ_MakeSeekChunkLen")
	purego.RegisterLibFunc(&getNumSeekChunks, libHandle, "OodleLZ_GetNumSeekChunks")
	purego.RegisterLibFunc(&getSeekTableMemorySizeNeeded, libHandle, "OodleLZ_GetSeekTableMemorySizeNeeded")
	purego.RegisterLibFunc(&fillSeekTable, libHandle, "OodleLZ_FillSeekTable")
	purego.RegisterLibFunc(&createSeekTable, libHandle, "OodleLZ_CreateSeekTable")
	purego.RegisterLibFunc(&freeSeekTable, libHandle, "OodleLZ_FreeSeekTable")
	purego.RegisterLibFunc(&checkSeekTableCrcs, libHandle, "OodleLZ_CheckSeekTableCRCs")
	purego.RegisterLibFunc(&findSeekEntry, libHandle, "OodleLZ_FindSeekEntry")
	purego.RegisterLibFunc(&getSeekEntryPackedPos, libHandle, "OodleLZ_GetSeekEntryPackedPos")
	purego.RegisterLibFunc(&compressionLevelGetName, libHandle, "OodleLZ_CompressionLevel_GetName")
	purego.RegisterLibFunc(&compressorGetName, libHandle, "OodleLZ_Compressor_GetName")
	purego.RegisterLibFunc(&jobifyGetName, libHandle, "OodleLZ_Jobify_GetName")
	purego.RegisterLibFunc(&compressOptionsGetDefault, libHandle, "OodleLZ_CompressOptions_GetDefault")
	purego.RegisterLibFunc(&compressOptionsValidate, libHandle, "OodleLZ_CompressOptions_Validate")
	purego.RegisterLibFunc(&getCompressScratchMemBound, libHandle, "OodleLZ_GetCompressScratchMemBound")
	purego.RegisterLibFunc(&getCompressScratchMemBoundEx, libHandle, "OodleLZ_GetCompressScratchMemBoundEx")
	purego.RegisterLibFunc(&getCompressedBufferSizeNeeded, libHandle, "OodleLZ_GetCompressedBufferSizeNeeded")
	purego.RegisterLibFunc(&getDecodeBufferSize, libHandle, "OodleLZ_GetDecodeBufferSize")
	purego.RegisterLibFunc(&getInPlaceDecodeBufferSize, libHandle, "OodleLZ_GetInPlaceDecodeBufferSize")
	purego.RegisterLibFunc(&getCompressedStepForRawStep, libHandle, "OodleLZ_GetCompressedStepForRawStep")
	purego.RegisterLibFunc(&getAllChunksCompressor, libHandle, "OodleLZ_GetAllChunksCompressor")
	purego.RegisterLibFunc(&getFirstChunkCompressor, libHandle, "OodleLZ_GetFirstChunkCompressor")
	purego.RegisterLibFunc(&getChunkCompressor, libHandle, "OodleLZ_GetChunkCompressor")
	purego.RegisterLibFunc(&getConfigValues, libHandle, "Oodle_GetConfigValues")
}
