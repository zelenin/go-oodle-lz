package oodle

import (
	"io"
)

func NewReader(r io.Reader, outSize int64) (*Reader, error) {
	bufSize := BLOCK_LEN + 64*1024

	return &Reader{
		r:       r,
		outSize: outSize,
		decoder: DecoderCreate(CompressorInvalid, outSize, nil),
		compBuf: make([]byte, bufSize),
		buf:     make([]byte, bufSize),
	}, nil
}

type Reader struct {
	r                  io.Reader
	outSize            int64
	decoder            Decoder
	compBuf            []byte
	compBufR, compBufW int
	buf                []byte
	bufR, bufW         int
	hasBeenRead        bool
}

func (r *Reader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	if r.hasBeenRead && r.bufR == r.bufW && r.compBufR == r.compBufW {
		return 0, io.EOF
	}

	if r.bufR == r.bufW {
		r.bufR = 0
		r.bufW = 0
		if r.compBufR != r.compBufW {
			copy(r.compBuf, r.compBuf[r.compBufR:r.compBufW])
			r.compBufW = r.compBufW - r.compBufR
			r.compBufR = 0
		}
		err := r.fill()
		if err != nil {
			return 0, err
		}

		out, err := DecoderDecodeSome(
			r.decoder,
			r.buf,
			r.bufW,
			int(r.outSize),
			len(r.buf)-r.bufW,
			r.compBuf[r.compBufR:r.compBufW],
			FuzzSafeYes,
			CheckCRCYes,
			VerbosityNone,
			DecodeUnthreaded,
		)
		if err != nil {
			return 0, err
		}

		r.compBufR += int(out.CompBufUsed)
		r.bufW += int(out.DecodedCount)

		if out.DecodedCount == 0 && r.compBufR == r.compBufW {
			err = r.fill()
			if err != nil {
				return 0, err
			}
		}
	}

	n := copy(p, r.buf[r.bufR:r.bufW])
	r.bufR += n

	return n, nil
}

func (r *Reader) fill() error {
	if r.hasBeenRead {
		return nil
	}
	n, err := r.r.Read(r.compBuf[r.compBufW:])
	if err != nil {
		if err == io.EOF {
			if n == 0 {
				r.hasBeenRead = true
			} else {
				r.compBufW += n
				r.compBufR = 0
			}
		} else {
			return err
		}
	} else {
		r.compBufW += n
	}

	return nil
}

func (r *Reader) Close() error {
	DecoderDestroy(r.decoder)
	return nil
}
