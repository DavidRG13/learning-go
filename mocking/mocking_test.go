package mocking

import (
	"reflect"
	"testing"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	countdownOperationsSpy := &CountdownOperationsSpy{}

	Countdown(countdownOperationsSpy, countdownOperationsSpy)

	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, countdownOperationsSpy.Calls) {
		t.Errorf("wanted calls %v got %v", want, countdownOperationsSpy.Calls)
	}
}
