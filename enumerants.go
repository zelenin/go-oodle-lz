package oodle

type Verbosity int

const (
	VerbosityNone    Verbosity = 0
	VerbosityMinimal Verbosity = 1
	VerbositySome    Verbosity = 2
	VerbosityLots    Verbosity = 3
	VerbosityForce32 Verbosity = 0x40000000
)

type Compressor int

const (
	CompressorInvalid Compressor = -1
	CompressorNone    Compressor = 3
	// new compressors
	CompressorKraken    Compressor = 8
	CompressorLeviathan Compressor = 13
	CompressorMermaid   Compressor = 9
	CompressorSelkie    Compressor = 11
	CompressorHydra     Compressor = 12
	// deprecated compressors
	CompressorBitKnit Compressor = 10
	CompressorLZB16   Compressor = 4
	CompressorLZNA    Compressor = 7
	CompressorLZH     Compressor = 0
	CompressorLZHLW   Compressor = 1
	CompressorLZNIB   Compressor = 2
	CompressorLZBLW   Compressor = 5
	CompressorLZA     Compressor = 6
	//
	CompressorCount   Compressor = 14
	CompressorForce32 Compressor = 0x40000000
)

type PackedRawOverlap int

const (
	PackedRawOverlapNo      PackedRawOverlap = 0
	PackedRawOverlapYes     PackedRawOverlap = 1
	PackedRawOverlapForce32 PackedRawOverlap = 0x40000000
)

type CheckCrc int

const (
	CheckCRCNo      CheckCrc = 0
	CheckCRCYes     CheckCrc = 1
	CheckCRCForce32 CheckCrc = 0x40000000
)

type Profile int

const (
	ProfileMain    Profile = 0
	ProfileReduced Profile = 1
	ProfileForce32 Profile = 0x40000000
)

type DecompressCallbackRet int

const (
	DecompressCallbackRetContinue DecompressCallbackRet = 0
	DecompressCallbackRetCancel   DecompressCallbackRet = 1
	DecompressCallbackRetInvalid  DecompressCallbackRet = 2
	DecompressCallbackRetForce32  DecompressCallbackRet = 0x40000000
)

type CompressionLevel int

const (
	CompressionLevelNone       CompressionLevel = 0
	CompressionLevelSuperFast  CompressionLevel = 1
	CompressionLevelVeryFast   CompressionLevel = 2
	CompressionLevelFast       CompressionLevel = 3
	CompressionLevelNormal     CompressionLevel = 4
	CompressionLevelOptimal1   CompressionLevel = 5
	CompressionLevelOptimal2   CompressionLevel = 6
	CompressionLevelOptimal3   CompressionLevel = 7
	CompressionLevelOptimal4   CompressionLevel = 8
	CompressionLevelOptimal5   CompressionLevel = 9
	CompressionLevelHyperFast1 CompressionLevel = -1
	CompressionLevelHyperFast2 CompressionLevel = -2
	CompressionLevelHyperFast3 CompressionLevel = -3
	CompressionLevelHyperFast4 CompressionLevel = -4
	CompressionLevelHyperFast  CompressionLevel = CompressionLevelHyperFast1
	CompressionLevelOptimal    CompressionLevel = CompressionLevelOptimal2
	CompressionLevelMax        CompressionLevel = CompressionLevelOptimal5
	CompressionLeve_Min        CompressionLevel = CompressionLevelHyperFast4
	CompressionLevelForce32    CompressionLevel = 0x40000000
	CompressionLevelInvalid    CompressionLevel = CompressionLevelForce32
)

type Jobify int

const (
	JobifyDefault    Jobify = 0
	JobifyDisable    Jobify = 1
	JobifyNormal     Jobify = 2
	JobifyAggressive Jobify = 3
	JobifyCount      Jobify = 4
	JobifyForce32    Jobify = 0x40000000
)

type DecodeThreadPhase int

const (
	DecodeThreadPhase1   DecodeThreadPhase = 1
	DecodeThreadPhase2   DecodeThreadPhase = 2
	DecodeThreadPhaseAll DecodeThreadPhase = 3
	DecodeUnthreaded     DecodeThreadPhase = DecodeThreadPhaseAll
)

type FuzzSafe int

const (
	FuzzSafeNo  FuzzSafe = 0
	FuzzSafeYes FuzzSafe = 1
)

type SeekTableFlags int

const (
	SeekTableFlagsNone        SeekTableFlags = 0
	SeekTableFlagsMakeRawCRCs SeekTableFlags = 1
	SeekTableFlagsForce32     SeekTableFlags = 0x40000000
)

type CompressScratchMemBoundType int

const (
	CompressScratchMemBoundTypeWorstCase CompressScratchMemBoundType = 0
	CompressScratchMemBoundTypeTypical   CompressScratchMemBoundType = 1
	CompressScratchMemBoundTypeForce32   CompressScratchMemBoundType = 0x40000000
)
