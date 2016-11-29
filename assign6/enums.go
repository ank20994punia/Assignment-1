package main

import "fmt"

func main() {
        var eval Eval=1
        var eval_1 Eval=2
        var eval_2 Eval=3
        var eval_3 Eval=4
	fmt.Printf("%s",eval)
	fmt.Printf("\n%s",eval_1)
        fmt.Printf("\n%s",eval_2)
        fmt.Printf("\n%s\n",eval_3)
}

type Eval int

const (
        _start=iota
	One
	Two
	Three
        Four
)

var myTypeValues = []string{
	One:   "One",
	Two:   "Two",
	Three: "Three",
        Four: "Four",
}

func (t Eval) String() string {
	if t <=0 || int(t) >= len(myTypeValues) {
		return "Invalid"
	}
	return myTypeValues[t]
}
