package go_hivereg

import (
	"github.com/jc-lab/go-hivereg/model/regtype"
	"io"
)

type Store interface {
	io.Closer
	AddKey(key string) error
	DeleteKey(key string) error
	AddValue(key string, valueName string, dataType regtype.RegType, separator string, data string) error
	DeleteValue(key string, valueName string) error
}
