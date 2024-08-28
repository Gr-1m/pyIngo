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

type Bruter struct {
	SingleTask func(b interface{}, args ...interface{}) error

	Threads int
	taskNum int

	dataIn chan interface{}
	result chan interface{}

	ResultData []interface{}
	ErrData    []error
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

type StringBruter struct {
	SingleTask func(b string, args ...interface{}) error

	Threads int
	taskNum int

	dataIn chan string
	result chan string

	ResultData []string
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

func (bt *IntBruter) goWork(args ...interface{}) {

	for b := range bt.dataIn {
		if err := bt.SingleTask(b, args...); err != nil {
			bt.result <- 0
			bt.ErrData = append(bt.ErrData, err)
		} else {
			bt.result <- b
			continue
		}
	}
}

func (bt *StringBruter) goWork(args ...interface{}) {

	for b := range bt.dataIn {
		if err := bt.SingleTask(b, args...); err != nil {
			bt.result <- ""
			bt.ErrData = append(bt.ErrData, err)
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

	for i := 1; i < cap(bt.dataIn); i++ {
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

	for i := 1; i < cap(bt.dataIn); i++ {
		go bt.goWork(args...)
	}

	go func(numRange [2]int) {
		var pgchan = make(chan int)
		defer close(pgchan)

		pgbar.Play(pgchan, bt.taskNum)
		// go 1.22
		for pg := range numRange[1] - numRange[0] {
			pgchan <- pg + 1
			bt.dataIn <- numRange[0] + pg
		}
		// go < 1.22
		/*
			for pg := numRange[0]; pg < numRange[1]; pg++ {
				pgchan <- (pg - numRange[0] + 1)
				bt.dataIn <- pg
			}*/
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

	for i := 1; i < cap(bt.dataIn); i++ {
		go bt.goWork(args...)
	}

	var linenum int
	var fileScanner = bufio.NewScanner(file)
	for fileScanner.Scan() {
		linenum++
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}

	bt.taskNum = linenum
	if bt.Threads > bt.taskNum {
		return errors.New("Threads > taskNum Error!")
	}

	go func(file *os.File) {
		var pgchan = make(chan int)
		defer close(pgchan)

		pgbar.Play(pgchan, bt.taskNum)
		linenum = 0

		fileScanner = bufio.NewScanner(file)
		for fileScanner.Scan() {
			pgchan <- (linenum + 1)
			bt.dataIn <- fileScanner.Text()
			linenum++
		}
		if err := fileScanner.Err(); err != nil {
			panic(err)
		}
	}(file)

	for ; bt.taskNum > 0; bt.taskNum-- {
		r := <-bt.result
		if r == "" {
			continue
		}
		bt.ResultData = append(bt.ResultData, r)
	}

	return nil
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
		results []int // whatever type you want
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
