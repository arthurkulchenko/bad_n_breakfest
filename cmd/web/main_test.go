package main

import "testing"

func TestRun(test *testing.T) {
	err := run()
	if err != nil {
		test.Error("failed run()")
	}
}
