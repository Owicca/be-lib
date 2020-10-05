package primitive

import (
	"github.com/Owicca/be-lib/object"
)

type Primitive interface {
	object.Object
	GetValue() interface{}
	SetValue(interface{}) error
}
