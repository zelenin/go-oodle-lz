package oodle

import (
	"bytes"
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestCompressDecompress(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			compressed, err := tc.Compress()
			if err != nil {
				t.Fatalf("Compression failed: %v", err)
			}

			if len(tc.data) > 0 && len(compressed) == 0 {
				t.Error("Compressed data is empty")
			}

			decompressed, err := Decompress(
				compressed,
				len(tc.data),
			)
			if err != nil {
				t.Fatalf("Decompression failed: %v", err)
			}

			if !bytes.Equal(tc.data, decompressed) {
				t.Error("Decompressed data doesn't match original input")
			}
		})
	}
}

func TestGetCompressedBufferSizeNeeded(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			size := GetCompressedBufferSizeNeeded(tc.compressor, len(tc.data))

			if size <= 0 {
				t.Errorf("Buffer size is too small: %d", size)
			}
		})
	}
}

func TestGetDecodeBufferSize(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			size := GetDecodeBufferSize(tc.compressor, len(tc.data), false)

			if size < len(tc.data) {
				t.Errorf("Buffer size is too small: %d", size)
			}
		})
	}
}

func TestGetInPlaceDecodeBufferSize(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			compressed, err := tc.Compress()
			if err != nil {
				t.Fatalf("Compression failed: %v", err)
			}

			if len(tc.data) > 0 && len(compressed) == 0 {
				t.Error("Compressed data is empty")
			}

			size := GetInPlaceDecodeBufferSize(tc.compressor, len(compressed), len(tc.data))

			if size <= len(compressed) {
				t.Errorf("Buffer size is too small: %d", size)
			}
		})
	}
}

func TestGetAllChunksCompressor(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			compressed, err := tc.Compress()
			if err != nil {
				t.Fatalf("Compression failed: %v", err)
			}

			if len(tc.data) > 0 && len(compressed) == 0 {
				t.Error("Compressed data is empty")
			}

			compressor := GetAllChunksCompressor(compressed, len(compressed), len(compressed))

			if compressor != tc.compressor {
				t.Logf("Compressor doesn't match. Want: %s, got: %s", CompressorGetName(tc.compressor), CompressorGetName(compressor))
			} else {
				t.Logf("Compressor: %s", CompressorGetName(compressor))
			}
		})
	}
}

func TestGetFirstChunkCompressor(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			compressed, err := tc.Compress()
			if err != nil {
				t.Fatalf("Compression failed: %v", err)
			}

			if len(tc.data) > 0 && len(compressed) == 0 {
				t.Error("Compressed data is empty")
			}

			compressor := GetFirstChunkCompressor(compressed, len(compressed), nil)

			if compressor != tc.compressor {
				t.Logf("Compressor doesn't match. Want: %s, got: %s", CompressorGetName(tc.compressor), CompressorGetName(compressor))
			} else {
				t.Logf("Compressor: %s", CompressorGetName(compressor))
			}
		})
	}
}

func TestGetChunkCompressor(t *testing.T) {
	testCases := getCompressionVariants()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.data) == 0 {
				t.Skipf("Empty data")
			}

			compressed, err := tc.Compress()
			if err != nil {
				t.Fatalf("Compression failed: %v", err)
			}

			if len(tc.data) > 0 && len(compressed) == 0 {
				t.Error("Compressed data is empty")
			}

			compressor := GetChunkCompressor(compressed, len(compressed), nil)

			if compressor != tc.compressor {
				t.Logf("Compressor doesn't match. Want: %s, got: %s", CompressorGetName(tc.compressor), CompressorGetName(compressor))
			} else {
				t.Logf("Compressor: %s", CompressorGetName(compressor))
			}
		})
	}
}

type compressionVariant struct {
	name       string
	data       []byte
	compressor Compressor
	level      CompressionLevel
}

func (v *compressionVariant) Compress() ([]byte, error) {
	return Compress(v.compressor, v.data, v.level)
}

func getCompressionVariants() []*compressionVariant {
	compressors := map[Compressor]string{
		CompressorKraken:    CompressorGetName(CompressorKraken),
		CompressorLeviathan: CompressorGetName(CompressorLeviathan),
		CompressorMermaid:   CompressorGetName(CompressorMermaid),
		CompressorSelkie:    CompressorGetName(CompressorSelkie),
		CompressorHydra:     CompressorGetName(CompressorHydra),
	}

	compressionLevels := map[CompressionLevel]string{
		CompressionLevelFast:   CompressionLevelGetName(CompressionLevelFast),
		CompressionLevelNormal: CompressionLevelGetName(CompressionLevelNormal),
		CompressionLevelMax:    CompressionLevelGetName(CompressionLevelMax),
	}

	data := map[string][]byte{
		"text":   []byte("Hello, World! This is a test string that should be compressed and then decompressed."),
		"empty":  []byte{},
		"binary": []byte{0x00, 0xFF, 0x80, 0x7F, 0x3A, 0x12, 0xB4},
	}

	buf := bytes.NewBuffer(nil)
	var n int
	for i := 0; i < 5_000_000; i += n {
		n, _ = buf.WriteString("Hello, World! ")
	}
	data["5mb+"] = buf.Bytes()

	variants := make([]*compressionVariant, 0, len(compressors)*len(compressionLevels)*len(data))

	for _, compressor := range slices.Sorted(maps.Keys(compressors)) {
		for _, compressionLevel := range slices.Sorted(maps.Keys(compressionLevels)) {
			for _, d := range slices.Sorted(maps.Keys(data)) {
				variants = append(variants, &compressionVariant{
					name:       fmt.Sprintf("%s-%s-%q", compressors[compressor], compressionLevels[compressionLevel], d),
					data:       data[d],
					compressor: compressor,
					level:      compressionLevel,
				})
			}
		}
	}

	return variants
}
