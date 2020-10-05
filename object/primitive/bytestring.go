package primitive

import (
	"fmt"
	"regexp"
	"strings"
)

type ByteString struct {
	value string
}

func NewByteString(val string) (*ByteString, error) {
	bs := &ByteString{}
	if err := bs.SetValue(val); err != nil {
		return nil, fmt.Errorf("Could not set value (%s)", err)
	}

	return bs, nil
}

func (i *ByteString) GetValue() string {
	parts := strings.Split(i.value, ":")

	return parts[1]
}

func (i *ByteString) SetValue(val string) error {
	reg := regexp.MustCompile(`^\d{1,}\:\w{0,}`)

	if !reg.Match([]byte(val)) {
		return fmt.Errorf("Invalid value!")
	}
	i.value = val

	return nil
}
