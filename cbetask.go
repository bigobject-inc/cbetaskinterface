package cbetask

import (
	"os"

	cbeutility "bitbucket.org/bigobject/cbe-utility/v3"
	cbe "github.com/bigobject-inc/stailib/cbe"
	ttType "github.com/bigobject-inc/stailib/tt"
)

type Task interface {
	Init(*os.File, *cbeutility.CBEutility, func(string, ...ttType.Node) error) error
	Prepare(string, int, cbe.CBE) error
	Process() error
	GetName() string
}

type Event interface {
	Scan()
}

type Set interface {
	Add(cbe.CBEObjStruct)
}
