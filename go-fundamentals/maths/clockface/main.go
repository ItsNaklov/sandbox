package main

import (
	"os"
	"time"

	clockface "sandboxtest.com/test/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
