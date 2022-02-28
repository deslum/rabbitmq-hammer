package compression

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

type ZLib struct{}

func (o *ZLib) Encode(data []byte) ([]byte, error) {
	var buff bytes.Buffer
	w := zlib.NewWriter(&buff)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (o *ZLib) Decode(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := r.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func (o *ZLib) String() string {
	return "ZLib"
}
