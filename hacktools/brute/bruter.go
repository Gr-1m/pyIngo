package brute

type Bruter struct {
	SingleTask func(b interface{}, args ...interface{}) bool

	TaskNum int
	threads int
	timeout int
}

func (bt *Bruter) Start(bs []interface{}, args ...interface{}) {
	// The default recommended Thread is 700
	// The default recommended Timeout is 60
	bt.TaskNum = len(bs)

	data := make(chan interface{}, bt.threads)
	results := make(chan interface{})

	defer close(results)
	defer close(data)

	for i := 1; i < cap(data); i++ {
		go bt.Worker(data, results, args...)
	}

	go func() {
		for b := range bs {
			data <- b
		}
	}()

	var succs []interface{}
	for ; bt.TaskNum > 0; bt.TaskNum-- {
		r := <-results
		if r != nil {
			succs = append(succs, r)
		}
	}

	return
}

func (bt *Bruter) Worker(data, results chan interface{}, args ...interface{}) {
	for b := range data {
		if bt.SingleTask(b, args...) {
			results <- b
			continue
		} else {
			results <- nil
		}
	}
}
