package b

import (
	"a"
	"fmt"
	"log/syslog"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
	syslog.New(0, "")
}
