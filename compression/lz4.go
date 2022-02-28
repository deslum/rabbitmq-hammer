package compression

import (
	"bytes"
	"io/ioutil"

	"github.com/pierrec/lz4"
)

type LZ4 struct{}

func (o *LZ4) Encode(data []byte) ([]byte, error) {
	var buff bytes.Buffer
	w := lz4.NewWriter(&buff)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (o *LZ4) Decode(data []byte) ([]byte, error) {
	r := lz4.NewReader(bytes.NewReader(data))
	result, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *LZ4) String() string {
	return "LZ4"
}
