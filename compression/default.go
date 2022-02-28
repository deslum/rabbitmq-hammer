package compression

type Default struct{}

func (o *Default) Encode(data []byte) ([]byte, error) {
	return data, nil
}

func (o *Default) Decode(data []byte) ([]byte, error) {
	return data, nil
}

func (o *Default) String() string {
	return "None"
}
