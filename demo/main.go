package main

import (
	"github.com/smartwalle/nox"
	"fmt"
)

func main() {
	var r = nox.NewRequest("GET", "http://www.baidu.com")
	rep := r.Exec()
	fmt.Println(rep.String())
}
