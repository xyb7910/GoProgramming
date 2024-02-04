package clog

import "testing"

func TestChoke(t *testing.T) {
	chokeWithWaitGroup()
	chokeWithSelect()
	chokeWithChannel()
	chokeWithTimeAfter()
	chokeWithFor()
	chokeWithMutex()
}
