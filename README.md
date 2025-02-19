# go-oodle-lz

Go wrapper for [OodleLZ](https://www.radgametools.com/oodlecompressors.htm)

### Compress
```go
compressedData, err := oodle.Compress(oodle.CompressorKraken, data, oodle.CompressionLevelMax)
```

### Decompress
```go
decompressedData, err := oodle.Decompress(compressedData, outputSize)
```

### Reader
```go
rc, err := oodle.NewReader(r, uncompSize)
if err != nil {
	log.Fatalf("failed to create reader: %v", err)
}
defer rc.Close()
```

## Notes

* WIP. Library API can be changed in the future

## Author

[Aleksandr Zelenin](https://github.com/zelenin/), e-mail: [aleksandr@zelenin.me](mailto:aleksandr@zelenin.me)
