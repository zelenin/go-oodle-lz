package oodle

import (
	"errors"
	"unsafe"
)

var compress func(Compressor, []byte, int, []byte, CompressionLevel, *CompressOptions, int, uintptr, unsafe.Pointer, int) uintptr

func Compress(compressor Compressor, rawBuf []byte, compressionLevel CompressionLevel) ([]byte, error) {
	rawLen := len(rawBuf)
	outputSize := GetCompressedBufferSizeNeeded(compressor, rawLen)
	compBuf := make([]byte, outputSize)

	pOptions := (*CompressOptions)(nil)
	dictionaryBase := 0
	var lrm uintptr = 0
	scratchMem := unsafe.Pointer(nil)
	scratchMemSize := 0

	n := compress(
		compressor,
		rawBuf,
		rawLen,
		compBuf,
		compressionLevel,
		pOptions,
		dictionaryBase,
		lrm,
		scratchMem,
		scratchMemSize,
	)

	if n == 0 {
		return nil, errors.New("compress failure")
	}

	data := make([]byte, n)
	copy(data, compBuf)

	return data, nil
}

var decompress func([]byte, int, []byte, int, FuzzSafe, CheckCrc, Verbosity, unsafe.Pointer, int, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int, DecodeThreadPhase) int

func Decompress(compBuf []byte, rawLen int) ([]byte, error) {
	compBufSize := len(compBuf)
	rawBuf := make([]byte, rawLen)

	decBufBase := unsafe.Pointer(nil)
	decBufSize := 0

	fpCallback := unsafe.Pointer(nil)
	callbackUserData := unsafe.Pointer(nil)

	decoderMemory := unsafe.Pointer(nil)
	decoderMemorySize := 0

	n := decompress(
		compBuf,
		compBufSize,
		rawBuf,
		rawLen,
		FuzzSafeYes,
		CheckCRCYes,
		VerbosityNone,
		decBufBase,
		decBufSize,
		fpCallback,
		callbackUserData,
		decoderMemory,
		decoderMemorySize,
		DecodeUnthreaded,
	)
	if n == 0 {
		return nil, errors.New("decompress failure")
	}

	return rawBuf, nil
}

var decoderCreate func(Compressor, int64, []byte, int) Decoder

func DecoderCreate(compressor Compressor, rawLen int64, memory []byte) Decoder {
	return decoderCreate(compressor, rawLen, memory, len(memory))
}

var decoderMemorySizeNeeded func(Compressor, int) int32

func DecoderMemorySizeNeeded(compressor Compressor, rawLen int) (int32, error) {
	res := decoderMemorySizeNeeded(compressor, rawLen)
	if res == 0 {
		return 0, errors.New("decoder memory size needed")
	}

	return res, nil
}

var threadPhasedBlockDecoderMemorySizeNeeded func() int32

func ThreadPhasedBlockDecoderMemorySizeNeeded() int32 {
	return threadPhasedBlockDecoderMemorySizeNeeded()
}

var decoderDestroy func(Decoder)

func DecoderDestroy(decoder Decoder) {
	decoderDestroy(decoder)
}

var decoderReset func(Decoder, int, int) bool

func DecoderReset(decoder Decoder, decPos int, decLen int) bool {
	return decoderReset(decoder, decPos, decLen)
}

var decoderDecodeSome func(Decoder, *DecodeSomeOut, []byte, int, int, int, []byte, int, FuzzSafe, CheckCrc, Verbosity, DecodeThreadPhase) bool

func DecoderDecodeSome(decoder Decoder, decBuf []byte, decBufPos int, decBufferSize int, decBufAvail int, compPtr []byte, fuzzSafe FuzzSafe, checkCrc CheckCrc, verbosity Verbosity, threadPhase DecodeThreadPhase) (*DecodeSomeOut, error) {
	out := &DecodeSomeOut{}

	res := decoderDecodeSome(
		decoder,
		out,
		decBuf,
		decBufPos,
		decBufferSize,
		decBufAvail,
		compPtr,
		len(compPtr),
		fuzzSafe,
		checkCrc,
		verbosity,
		threadPhase,
	)
	if !res {
		return nil, errors.New("decoder failure")
	}

	return out, nil
}

var makeValidCircularWindowSize func(Compressor, int32) int32

func MakeValidCircularWindowSize(compressor Compressor, minWindowSize int32) int32 {
	return makeValidCircularWindowSize(compressor, minWindowSize)
}

var makeSeekChunkLen func(int64, int32) int32

func MakeSeekChunkLen(rawLen int64, desiredSeekPointCount int32) int32 {
	return makeSeekChunkLen(rawLen, desiredSeekPointCount)
}

var getNumSeekChunks func(int64, int32) int32

func GetNumSeekChunks(rawLen int64, seekChunkLen int32) int32 {
	return getNumSeekChunks(rawLen, seekChunkLen)
}

var getSeekTableMemorySizeNeeded func(int32, SeekTableFlags) int

func GetSeekTableMemorySizeNeeded(numSeekChunks int32, flags SeekTableFlags) int {
	return getSeekTableMemorySizeNeeded(numSeekChunks, flags)
}

var fillSeekTable func(*SeekTable, SeekTableFlags, int32, []byte, int, []byte, int) bool

func FillSeekTable(pTable *SeekTable, flags SeekTableFlags, seekChunkLen int32, rawBuf []byte, rawLen int, compBuf []byte, compLen int) bool {
	return fillSeekTable(pTable, flags, seekChunkLen, rawBuf, rawLen, compBuf, compLen)
}

var createSeekTable func(SeekTableFlags, int32, []byte, int, []byte, int) *SeekTable

func CreateSeekTable(flags SeekTableFlags, seekChunkLen int32, rawBuf []byte, rawLen int, compBuf []byte, compLen int) *SeekTable {
	return createSeekTable(flags, seekChunkLen, rawBuf, rawLen, compBuf, compLen)
}

var freeSeekTable func(*SeekTable)

func FreeSeekTable(pTable *SeekTable) {
	freeSeekTable(pTable)
}

var checkSeekTableCrcs func([]byte, int, *SeekTable) bool

func CheckSeekTableCrcs(rawBuf []byte, rawLen int, seekTable *SeekTable) bool {
	return checkSeekTableCrcs(rawBuf, rawLen, seekTable)
}

var findSeekEntry func(int64, *SeekTable) int32

func FindSeekEntry(rawPos int64, seekTable *SeekTable) int32 {
	return findSeekEntry(rawPos, seekTable)
}

var getSeekEntryPackedPos func(int32, *SeekTable) int64

func GetSeekEntryPackedPos(seekI int32, seekTable *SeekTable) int64 {
	return getSeekEntryPackedPos(seekI, seekTable)
}

var compressionLevelGetName func(CompressionLevel) string

func CompressionLevelGetName(compressSelect CompressionLevel) string {
	return compressionLevelGetName(compressSelect)
}

var compressorGetName func(Compressor) string

func CompressorGetName(compressor Compressor) string {
	return compressorGetName(compressor)
}

var jobifyGetName func(Jobify) string

func JobifyGetName(jobify Jobify) string {
	return jobifyGetName(jobify)
}

var compressOptionsGetDefault func(Compressor, CompressionLevel) *CompressOptions

func CompressOptionsGetDefault(compressor Compressor, lzLevel CompressionLevel) *CompressOptions {
	return compressOptionsGetDefault(compressor, lzLevel)
}

var compressOptionsValidate func(*CompressOptions)

func CompressOptionsValidate(pOptions *CompressOptions) {
	compressOptionsValidate(pOptions)
}

var getCompressScratchMemBound func(Compressor, CompressionLevel, int, *CompressOptions) int

func GetCompressScratchMemBound(compressor Compressor, level CompressionLevel, rawLen int, pOptions *CompressOptions) int {
	return getCompressScratchMemBound(compressor, level, rawLen, pOptions)
}

var getCompressScratchMemBoundEx func(Compressor, CompressionLevel, CompressScratchMemBoundType, int, *CompressOptions) int

func GetCompressScratchMemBoundEx(compressor Compressor, level CompressionLevel, boundType CompressScratchMemBoundType, rawLen int, pOptions *CompressOptions) int {
	return getCompressScratchMemBoundEx(compressor, level, boundType, rawLen, pOptions)
}

var getCompressedBufferSizeNeeded func(Compressor, int) int

func GetCompressedBufferSizeNeeded(compressor Compressor, rawSize int) int {
	return getCompressedBufferSizeNeeded(compressor, rawSize)
}

var getDecodeBufferSize func(Compressor, int, bool) int

func GetDecodeBufferSize(compressor Compressor, rawSize int, corruptionPossible bool) int {
	return getDecodeBufferSize(compressor, rawSize, corruptionPossible)
}

var getInPlaceDecodeBufferSize func(Compressor, int, int) int

func GetInPlaceDecodeBufferSize(compressor Compressor, compLen int, rawLen int) int {
	return getInPlaceDecodeBufferSize(compressor, compLen, rawLen)
}

var getCompressedStepForRawStep func([]byte, int, int, int, *int, *int) int

func GetCompressedStepForRawStep(comp []byte, compAvail int, startRawPos int, rawSeekBytes int, pEndRawPos *int, pIndependent *int) int {
	return getCompressedStepForRawStep(comp, compAvail, startRawPos, rawSeekBytes, pEndRawPos, pIndependent)
}

var getAllChunksCompressor func([]byte, int, int) Compressor

func GetAllChunksCompressor(compBuf []byte, compBufSize int, rawLen int) Compressor {
	return getAllChunksCompressor(compBuf, compBufSize, rawLen)
}

var getFirstChunkCompressor func([]byte, int, *bool) Compressor

func GetFirstChunkCompressor(compChunk []byte, compBufAvail int, pIndependent *bool) Compressor {
	return getFirstChunkCompressor(compChunk, compBufAvail, pIndependent)
}

var getChunkCompressor func([]byte, int, *bool) Compressor

func GetChunkCompressor(compChunk []byte, compBufAvail int, pIndependent *bool) Compressor {
	return getChunkCompressor(compChunk, compBufAvail, pIndependent)
}

var getConfigValues func(*ConfigValues) Compressor

func GetConfigValues(configValues *ConfigValues) Compressor {
	return getConfigValues(configValues)
}

func nullablePointer[T any](val *T) unsafe.Pointer {
	var valPtr unsafe.Pointer
	if val == nil {
		valPtr = unsafe.Pointer(nil)
	} else {
		valPtr = unsafe.Pointer(&val)
	}

	return valPtr
}
