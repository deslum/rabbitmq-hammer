package compression

import (
	"bytes"
	"io/ioutil"

	"github.com/golang/snappy"
)

type Snappy struct{}

func (o *Snappy) Encode(data []byte) ([]byte, error) {
	var buff bytes.Buffer
	w := snappy.NewBufferedWriter(&buff)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (o *Snappy) Decode(data []byte) ([]byte, error) {
	r := snappy.NewReader(bytes.NewReader(data))
	result, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *Snappy) String() string {
	return "Snappy"
}
