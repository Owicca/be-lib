package complex

import (
	"fmt"

	mn "github.com/Owicca/be-lib/object"
	"github.com/Owicca/be-lib/object/primitive"
)

type List struct {
	value []mn.Object
}

func (l List) Run() {} // just for the Object interface

func NewList(val string) (*List, error) {
	ls := &List{}
	if err := ls.SetValue(val); err != nil {
		return nil, fmt.Errorf("Invalid value (%s)!", err)
	}

	return ls, nil
}

func (l *List) SetValue(val string) error { //TODO: a sequential algorithm that creates the underlying Object list
	integer, err := primitive.NewInteger("i44e")
	if err != nil {
		return fmt.Errorf("Could not create child object (%s)", err)
	}

	l.value = []mn.Object{*integer}

	return nil
}

func (l *List) GetValue() []mn.Object {
	return l.value
}
