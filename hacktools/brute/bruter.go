package brute

import (
	"bufio"
	"errors"
	"io"
	"os"
	"pyIngo/pgbar"
	"sort"
)

type Worker interface {
	goWork(datain, result chan interface{}, args ...interface{})
}

// type SingleTask func(b interface{}, args ...interface{}) error

type Bruter struct {
	SingleTask func(b interface{}, args ...interface{}) error

	Threads int
	taskNum int

	dataIn chan interface{}
	result chan interface{}

	ResultData []interface{}
	ErrData    []error
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

type IntBruter struct {
	SingleTask func(b int, args ...interface{}) error

	Threads int
	taskNum int

	dataIn chan int
	result chan int

	ResultData []int
	ErrData    []error
}

func (bt *IntBruter) goWork(args ...interface{}) {

	for b := range bt.dataIn {
		if err := bt.SingleTask(b, args...); err != nil {
			bt.result <- 0
		} else {
			bt.result <- b
			continue
		}
	}
}

type StringBruter struct {
	SingleTask func(b string, args ...interface{}) error

	Threads int
	taskNum int

	dataIn chan string
	result chan string

	ResultData []string
	ErrData    []error
}

func (bt *StringBruter) goWork(args ...interface{}) {

	for b := range bt.dataIn {
		if err := bt.SingleTask(b, args...); err != nil {
			bt.result <- ""
		} else {
			bt.result <- b
			continue
		}
	}
}

// todo: The data inflow here still needs to be optimized
// todo: bs -> []interface{}  or  chan interface{}
func (bt *Bruter) Start(bs []interface{}, args ...interface{}) error {
	bt.taskNum = len(bs)

	bt.dataIn = make(chan interface{}, bt.Threads)
	bt.result = make(chan interface{})

	defer close(bt.dataIn)
	defer close(bt.result)

	if bt.Threads > bt.taskNum {
		return errors.New("Threads > taskNum Error!")
	}

	// for i := 1; i < cap(bt.dataIn); i++ {
	for range cap(bt.dataIn) {
		go bt.goWork(args...)
	}

	go func(bs []interface{}) {
		var pgchan = make(chan int)
		defer close(pgchan)

		pgbar.Play(pgchan, bs)
		for pg, b := range bs {
			pgchan <- (pg + 1)
			bt.dataIn <- b
		}
		// pgbar.Clear()
	}(bs)

	for ; bt.taskNum > 0; bt.taskNum-- {
		r := <-bt.result
		if err, ok := r.(error); ok {
			bt.ErrData = append(bt.ErrData, err)
		} else {
			bt.ResultData = append(bt.ResultData, r)
		}
	}

	return nil
}

func (bt *IntBruter) StartWithInt(numRange [2]int, args ...interface{}) error {
	bt.taskNum = numRange[1] - numRange[0]

	bt.dataIn = make(chan int, bt.Threads)
	bt.result = make(chan int)

	defer close(bt.dataIn)
	defer close(bt.result)

	if bt.Threads > bt.taskNum {
		return errors.New("Threads > taskNum Error!")
	}

	// for i := 1; i < cap(bt.dataIn); i++ {
	for range cap(bt.dataIn) {
		go bt.goWork(args...)
	}

	go func(numRange [2]int) {
		var pgchan = make(chan int)
		defer close(pgchan)

		pgbar.Play(pgchan, bt.taskNum)
		// go <1.22
		/*
			for pg := numRange[0]; pg < numRange[1]; pg++ {
				pgchan <- (pg - numRange[0] + 1)
				bt.dataIn <- pg
			}*/
		// go 1.22
		for pg := range numRange[1] - numRange[0] {
			bt.dataIn <- numRange[0] + pg
			pgchan <- pg + 1
		}
	}(numRange)

	for ; bt.taskNum > 0; bt.taskNum-- {
		r := <-bt.result
		if r == 0 {
			continue
		}
		bt.ResultData = append(bt.ResultData, r)
	}
	sort.Ints(bt.ResultData)

	return nil
}

func (bt *StringBruter) StartWithFile(file *os.File, args ...interface{}) error {

	bt.dataIn = make(chan string, bt.Threads)
	bt.result = make(chan string)

	defer close(bt.dataIn)
	defer close(bt.result)
	defer file.Close()

	// for i := 1; i < cap(bt.dataIn); i++ {
	for range cap(bt.dataIn) {
		go bt.goWork(args...)
	}

	var linenum int
	var fileScanner = bufio.NewScanner(file)
	for fileScanner.Scan() {
		linenum++
	}
	bt.taskNum = linenum

	if bt.Threads > bt.taskNum {
		return errors.New("Threads > taskNum Error!")
	}

	go func() {
		var pgchan = make(chan int)
		defer close(pgchan)

		pgbar.Play(pgchan, bt.taskNum)
		linenum = 0

		file.Seek(0, io.SeekStart)
		fileScanner = bufio.NewScanner(file)
		for fileScanner.Scan() {
			pgchan <- (linenum + 1)
			bt.dataIn <- fileScanner.Text()
			linenum++
		}
		// pgbar.Clear()
	}()

	for ; bt.taskNum > 0; bt.taskNum-- {
		r := <-bt.result
		if r == "" {
			continue
		}
		bt.ResultData = append(bt.ResultData, r)
	}

	return nil
}

/*
---------------****************---------------
******** The Follow Function and Variable e.g. is just an Example

********/
// Recommended Example Template
// singleTask
func singleTask(b int, args ...interface{}) (err error) {

	var (
		ok       bool
		yourVar1 string
		yourVar2 int
	)

	// If your function needs to control the number of variables
	// If you want to check the type of a variable

	// You should check you argsList type in this function
	if len(args) != 2 {
		err = errors.New("args number is wrong")
	} else {
		yourVar1, ok = args[0].(string)
		if !ok {
			err = errors.New("args0 type is wrong")
			return
		}
		yourVar2, ok = args[1].(int)
		if !ok {
			err = errors.New("args1 type is wrong")
			return
		}

		// _, ok = b.(int)
		if !ok {
			err = errors.New("The Data type is Wrong which from chan interface{}")
		}
	}

	// Implement your function logic here
	_ = yourVar1
	_ = yourVar2
	// YourFunc(b, yourVar1,yourVar2)

	return
}

// Recommended Example Template
// How to recv Result

// Recommended Example Template
// Recommended CoreLogic
func exampleScan(thread int, arg0, arg1 string) []int {
	var (
		bt      *IntBruter
		results []int // whatever type you want
	)

	bt.SingleTask = singleTask // function pointer
	bt.Threads = thread

	bt.StartWithInt([2]int{1, 65536}, arg0, arg1)

	for _, rd := range bt.ResultData {
		results = append(results, rd)
	}

	return results
}
