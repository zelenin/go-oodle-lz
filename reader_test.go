package oodle

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

func generateTestData(size int) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func TestReader(t *testing.T) {
	testCases := []struct {
		name string
		size int
	}{
		{"100B", 100},
		{"100KB", 100 * 1024},
		{"BLOCK_LEN", BLOCK_LEN},
		{"500KB", 500 * 1024},
		{"5MB", 5 * 1024 * 1024},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := generateTestData(tc.size)

			compData, err := Compress(CompressorKraken, data, CompressionLevelMax)
			if err != nil {
				t.Fatalf("compression failed: %v", err)
			}

			rc, err := NewReader(bytes.NewReader(compData), int64(len(data)))
			if err != nil {
				t.Fatalf("failed to create reader: %v", err)
			}
			defer rc.Close()

			decompData, err := io.ReadAll(rc)
			if err != nil {
				t.Fatalf("failed to read decompressed data: %v", err)
			}

			if !bytes.Equal(data, decompData) {
				t.Fatalf("decompressed data does not match original")
			}
		})
	}
}
