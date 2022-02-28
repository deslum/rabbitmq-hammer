package config

type Message []byte

func (o *Message) UnmarshalJSON(data []byte) error {
	*o = data
	return nil
}
