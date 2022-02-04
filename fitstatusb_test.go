package fitstatusb

import (
	"testing"
	"time"
)

var fsu *FitStatUSB

func TestSetup(t *testing.T) {
	var err error
	fsu, err = New()
	if err != nil {
		t.Error(err)
	}
}

func TestBlue(t *testing.T) {
	err := fsu.Write("#0000FF")
	if err != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)
}

func TestPulse(t *testing.T) {
	err := fsu.Write("B#FF0000-0200#000000-0600")
	if err != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)
}

func TestGreen(t *testing.T) {
	err := fsu.Write("#002200")
	if err != nil {
		t.Error(err)
	}
}
