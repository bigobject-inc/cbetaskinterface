package cbetask

import (
	"os"

	cbeutility "bitbucket.org/bigobject/cbe-utility/v3"
	ttType "github.com/bigobject-inc/stailib/tt"
)

type Task interface {
	Init(*os.File, *cbeutility.CBEutility, func(string, ...ttType.Node) error) error
	Prepare(string, int, *Set, []*Event) error
	Process() error
	GetName() string
}

type Event interface {
}

type Set interface {
}
