package brute

import (
	"errors"
)

type Worker interface {
	goWork(datain, result chan interface{}, args ...interface{})
}

type Bruter struct {
	SingleTask func(b interface{}, args ...interface{}) error

	Threads int
	taskNum int

	dataIn chan interface{}
	result chan interface{}

	ResultData []interface{}
	ErrData    []error
	// timeout int
}

func (bt *Bruter) goWork(args ...interface{}) {

	for b := range bt.dataIn {
		if err := bt.SingleTask(b, args...); err != nil {
			bt.result <- err
		} else {
			bt.result <- b
			continue
		}
	}
}

func (bt *Bruter) Start(bs []interface{}, args ...interface{}) {
	bt.taskNum = len(bs)

	bt.dataIn = make(chan interface{}, bt.Threads)
	bt.result = make(chan interface{})

	defer close(bt.dataIn)
	defer close(bt.result)

	for i := 1; i < cap(bt.dataIn); i++ {
		go bt.goWork(args...)
	}

	go func() {
		for b := range bs {
			bt.dataIn <- b
		}
	}()

	for ; bt.taskNum > 0; bt.taskNum-- {
		r := <-bt.result
		if err, ok := r.(error); ok {
			bt.ErrData = append(bt.ErrData, err)
		} else {
			bt.ResultData = append(bt.ResultData, r)
		}
	}

	return
}

// Recommended Example Template
func (bt *Bruter) singleTask(b interface{}, args ...interface{}) (err error) {

	var (
		ok       bool // Recommended Var
		yourVar1 string
		yourVar2 int
	)

	// your SingleTask want args length
	if len(args) != 2 {
		err = errors.New("args number is wrong")
	} else {
		yourVar1, ok = args[0].(string)
		if !ok {
			err = errors.New("args0 type is wrong")
		}
		yourVar2, ok = args[1].(int)
		if !ok {
			err = errors.New("args1 type is wrong")
		}

		_, ok = b.(int)
		if !ok {
			err = errors.New("The Data type is Wrong which from chan interface{}")
		}
	}

	// your SingleTask funcLogic
	_ = yourVar1
	_ = yourVar2

	return
}

// Recommended Example Template
// How to recv Result

// Recommended Example Template
// Recommended CoreLogic
func ExampleScan() {}
