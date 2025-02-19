package oodle

import "unsafe"

type CompressOptions struct {
	UnusedWasVerbosity           uint32
	MinMatchLen                  int32
	SeekChunkReset               bool
	SeekChunkLen                 int32
	Profile                      Profile
	DictionarySize               int32
	SpaceSpeedTradeoffBytes      int32
	UnusedWasMaxHuffmansPerChunk int32
	SendQuantumCRCs              bool
	MaxLocalDictionarySize       int32
	MakeLongRangeMatcher         bool
	MatchTableSizeLog2           int32
	Jobify                       Jobify
	JobifyUserPtr                unsafe.Pointer
	FarMatchMinLen               int32
	FarMatchOffsetLog2           int32
	Reserved                     uint32
}

type DecodeSomeOut struct {
	DecodedCount      int32
	CompBufUsed       int32
	CurQuantumRawLen  int32
	CurQuantumCompLen int32
}

type Decoder uintptr

type SeekTable struct {
	Compressor            Compressor
	SeekChunksIndependent bool
	TotalRawLen           int64
	TotalCompLen          int64
	NumSeekChunks         int32
	SeekChunkLen          int32
	SeekChunkCompLens     []uint32
	RawCrcs               []uint32
}

type ConfigValues struct {
	LZLWLRMStep                       int32
	LZLWLRMHashLength                 int32
	LZLWLRMJumpbits                   int32
	LZDecoderMaxStackSize             int32
	LZSmallBufferLZFallbackSizeUnused int32
	LZBackwardsCompatibleMajorVersion int32
	HeaderVersion                     uint32
}
