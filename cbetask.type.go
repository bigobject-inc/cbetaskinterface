package cbetask

import (
	"os"
	"sync"

	cbeutility "bitbucket.org/bigobject/cbe-utility/v3"
	cbe "github.com/bigobject-inc/stailib/cbe"
	ttType "github.com/bigobject-inc/stailib/tt"
)

type Set struct {
	context *TaskStruct
	data    []cbe.CBEObjStruct // [Set number]
}

type Event struct { // ??
	eType        string
	toInstanceID int // Sequence instance that receive the event

	toSequenceName string        // Sequence name that receive the event
	targets        []ttType.Node // Sequence name pair with list of globalIDs. map[varName]global_id

	location  string // the location where the event occurs, area/cctv name
	timestamp int    // timestamp when the event occurs
}

type TaskStruct struct {
	taskName  string
	timestamp int
	frameID   string

	objects *Set     // input for Task Process()
	events  []*Event // input for Task Process()

	// sequenceStartTable    map[string]func(*TaskStruct) []Sequence
	defineSequenceTable   map[string]*defineSequence // static - one for a sequence type
	sequenceInstanceTable map[string][]Sequence
	sequenceBlackList     map[string]map[string]bool // map[seqName]map[GlobalID]bool an object (GID) can't be used by multiple sequence instances

	raiseExternalFunc func(string, ...ttType.Node) error

	defineSetTable map[string]func(*TaskStruct) *Set // static - one for a set type
	setMap         map[string]*Set

	instanceIDcounter int

	eventTypeForInstanceID map[string]bool // true: has toInstanceID. false: has no toInstanceID
	logPath                string
	utility                *cbeutility.CBEutility
	CBEServ                cbe.CBEServ
	logFile                *os.File
	mutex                  sync.Mutex
}

type SequenceStruct struct {
	sequenceName string
	instanceID   int
	currentState int

	context   *TaskStruct
	objectMap map[string]ObjectMapStruct // use it to remember the object this sequence need
	define    *defineSequence
	repeats   map[int]*Repeat // (stateID + 1) * repeatDivider + conditionIndex, so 0~repeatDivider is for Sequence, repeatDivider~max is for States, repeatDivider for each

	//  states []State
	//	stateFunc      []func(*TaskStruct, *Event)

	stateSignal    chan int
	sequenceSignal chan int
	intervalSignal chan int

	// lifetime int
	// interval int
}

type ObjectMapStruct struct {
	savedData   ttType.Node
	currentData cbe.CBEObjStruct
}

// defineSequence and defineState define constants and function pointers
// static functions/variables are shared by sequence instances
type defineSequence struct {
	startFunc func(*TaskStruct, *Event) []Sequence

	states                []*defineState
	timeoutSequence       int
	timeoutSequenceAction func(*SequenceStruct)
	interval              int
	intervalAction        func(*SequenceStruct)
}

type defineState struct {
	endState bool                               // 2020/10/26, since exitSequence() is not mandatory in endState, seems like endState has no use anymore
	eventNum int                                // number of event cond functions
	guard    func(*SequenceStruct) (bool, bool) // check if objects are the targets
	conds    []func(*SequenceStruct) bool       // nil for end state, the first eventNum elements are event
	// repeats    	[]string
	condsActions       []func(*SequenceStruct, *Event)
	action             func(*SequenceStruct, *Event)
	timeoutSec         int // max time to stay in the state
	timeoutStateAction func(*SequenceStruct, *Event)
}

type Repeat struct {
	cnt           int
	max           int
	interval      int // second
	lastTimestamp int // second
}
