package complex

import (
	"github.com/Owicca/be-lib/object"
)

type Complex interface {
	object.Object
	GetValue() interface{}
	SetValue(interface{}) error
}
