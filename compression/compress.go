package compression

import (
	"encoding/json"
	"fmt"
)

type CompressType struct {
	ICompressor
}

func (o *CompressType) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	switch j {
	case "None":
		o.ICompressor = new(Default)
	case "Snappy":
		o.ICompressor = new(Snappy)
	case "LZ4":
		o.ICompressor = new(LZ4)
	case "ZLib":
		o.ICompressor = new(ZLib)
	default:
		return fmt.Errorf("error compress CompressType value")
	}

	return nil
}
