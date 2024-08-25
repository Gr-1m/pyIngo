package pgbar

import (
	"fmt"
	"time"
)

type Bar struct {
	percent uint8
	current int
	total   int
	rate    string
	graph   string
}

var DefaultBar Bar

func loading(ms int) {
	var jdt = "-\\|/"
	if ms == 0 {
		ms = 120
	}
	for i := 0; true; i++ {
		fmt.Printf("\x1b[01;40;36m>[%c] ----\x1b[0m\x1b[K\r", jdt[i%len(jdt)])
		time.Sleep(time.Millisecond * time.Duration(ms))
	}
}

func vlen(v interface{}) int {
	switch vt := v.(type) {
	case int:
		return vt
	case string:
		return len(vt)
	case []interface{}:
		return len(vt)
	default:
		fmt.Errorf("Not supported type")
		return 0
	}
}

func NewBar(v interface{}) (interface{}, *Bar) {
	var b *Bar

	b.graph = "#"
	b.total = vlen(v)

	// Any operation that changes the current property should be update by exec Bar.setRate
	// Strongly Recommended
	b.current = 0
	b.setRate(0)

	return v, b
}

func (b *Bar) getPercent() uint8 {
	return uint8(float32(b.current) / float32(b.total) * 100)
}

func (b *Bar) setRate(incret int) {

	switch incret {
	case 0:
	case 1:
		b.rate += b.graph
		b.percent = b.getPercent()
	default:
		for i := 0; i < incret; i++ {
			b.rate += b.graph
		}
		b.percent = b.getPercent()
	}

}

func (b *Bar) Play(cur int) {
	var jdt = "-\\|/"

	b.current = cur
	b.setRate(int(b.getPercent() - b.percent))

	fmt.Printf("\r\x1b[01;40;36m>[%c][%-100s]%3d%% \x1b[0m%8d/%d\x1b[K", jdt[int(b.percent)%len(jdt)], b.rate, b.percent, b.current, b.total)

}

func InitDBar(v interface{}) interface{} {
	// Initialize Default Bar

	DefaultBar.graph = "#"
	DefaultBar.total = vlen(v)
	DefaultBar.current = 0
	DefaultBar.setRate(0)

	return v
}

func DPlay(cur int) {
	go DefaultBar.Play(cur)
}

func Play(cur int, v interface{}) {
	// If you have performance requirements, please perform InitDBar before the loop
	// If you don't mind, you can use this garbage function for progress bar rendering
	//
	// (The author hasn't yes figured out how to optimize it, If you have any good suggestions, please pull request, Thanks very much)
	//

	InitDBar(v)
	go DefaultBar.Play(cur)

	if DefaultBar.percent == 100 {
		fmt.Printf("\r\x1b[01;40;36m[%-100s]100%% \x1b[0m%8d/%d\x1b[K\n", DefaultBar.rate, DefaultBar.total, DefaultBar.total)
	}
}

func Clear(){
	fmt.Println("\r\x1b[0m\x1b[K")
}
