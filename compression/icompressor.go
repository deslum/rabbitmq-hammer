package compression

type ICompressor interface {
	Encode(data []byte) ([]byte, error)
	Decode(data []byte) ([]byte, error)
	String() string
}
