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
	getName() string
}

type Sequence interface {
	init(*TaskStruct)
	getContext() *TaskStruct
	process(*Set, *Event)
	getName() string
	getID() int
	enterState(int, *Event) int
	//	exitState()
	exitSequence()
	setDefine(*defineSequence)
	getDefine() *defineSequence
	setSequenceTimeout(int) chan int
	setIntervalTimeout(int) chan int
	addBlackList()
	removeBlackList()
	resetIntervalTimeout()

	getObjectMap() map[string]ObjectMapStruct

	print(...interface{})
	endStateAction()
}
