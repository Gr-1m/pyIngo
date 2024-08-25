package brute

import (
	"bufio"
	"errors"
	"io"
	"os"
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

// todo: The data inflow here still needs to be optimized
// todo:
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
		for _, b := range bs {
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

func (bt *Bruter) StartWithFile(file *os.File, args ...interface{}) {

	bt.dataIn = make(chan interface{}, bt.Threads)
	bt.result = make(chan interface{})

	defer close(bt.dataIn)
	defer close(bt.result)
	defer file.Close()

	for i := 1; i < cap(bt.dataIn); i++ {
		go bt.goWork(args...)
	}

	var linenum int
	var fileScanner = bufio.NewScanner(file)
	for fileScanner.Scan() {
		linenum++
	}
	bt.taskNum = linenum

	go func() {
		file.Seek(0, io.SeekStart)
		fileScanner = bufio.NewScanner(file)
		for fileScanner.Scan() {
			bt.dataIn <- fileScanner.Text()
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
// singleTask
func singleTask(b interface{}, args ...interface{}) (err error) {

	var (
		ok       bool
		yourVar1 string
		yourVar2 int
	)

	// If your function needs to control the number of variables
	// If you want to check the type of a variable
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

	// Implement your function logic here
	_ = yourVar1
	_ = yourVar2

	return
}

// Recommended Example Template
// How to recv Result

// Recommended Example Template
// Recommended CoreLogic
func exampleScan(thread int, arg0, arg1 string) []int {
	var (
		bt      *Bruter
		data    []interface{}
		results []int // whatever type you want, same as func return
	)

	data = make([]interface{}, 65535)

	for i := 0; i < 65535; i++ {
		data[i] = i + 1
	}

	bt.SingleTask = singleTask // function pointer
	bt.Threads = thread

	bt.Start(data, arg0, arg1)

	for _, rd := range bt.ResultData {
		v, ok := rd.(int)
		if ok {
			results = append(results, v)
		}
	}
	// sort.Ints(results)

	return results
}
