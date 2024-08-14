package main

import (
	"flag"
	"os"
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	go func() {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		main()
	}()
	time.Sleep(5 * time.Second)
}
