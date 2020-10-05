package primitive

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Integer struct {
	value int
}

func (i Integer) Run() {} //just for the Object interface

func NewInteger(val string) (*Integer, error) {
	integer := &Integer{}
	if err := integer.SetValue(val); err != nil {
		return nil, fmt.Errorf("Could not set value (%s)", err)
	}

	return integer, nil
}

func (i *Integer) GetValue() int {
	return i.value
}

func (i *Integer) SetValue(val string) error {
	reg := regexp.MustCompile(`^i\d{1,}e$`)

	if !reg.Match([]byte(val)) {
		return fmt.Errorf("Invalid value!")
	}

	result := strings.TrimPrefix(val, "i")
	result = strings.TrimSuffix(result, "e")

	in, err := strconv.Atoi(result)
	if err != nil {
		return fmt.Errorf("Could not convert %s (%s)!", result, err)
	}
	i.value = in

	return nil
}
