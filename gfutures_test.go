package gfutures

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var future Future

func answer(s interface{}) *Future {
	return NewFuture(func() interface{} {
		if s != nil {
			res := time.Duration(rand.Intn(1e3)) * time.Millisecond
			time.Sleep(res)
			return fmt.Sprintf("%s", s)
		}
		return errors.New("Error")
	})
}

func TestSet(t *testing.T) {
	future := answer("foo")
	res, err := future.Get()
	if nil != err {
		t.Errorf("Expected nil. Result %v", err)
		return
	}
	if "foo" != res {
		t.Errorf("Expected foo. Result %v", res)
	}
	future = answer("bar")
	res, err = future.Get()
	if "bar" != res {
		t.Errorf("Expected bar. Result %v", res)
	}
}

func TestReSet(t *testing.T) {
	future := answer(nil)
	res, err := future.Get()
	if nil != res {
		t.Errorf("Expected nil. Result %v", res)
	}
	if "Error" != err.Error() {
		t.Errorf("Expected Error. Result %v", err)
	}
}

func TestReSetAgain(t *testing.T) {
	future := answer("ok cool")
	res, err := future.Get()
	if nil != err {
		t.Errorf("Expected nil. Result %v", err)
		return
	}
	if "ok cool" != res {
		t.Errorf("Expected ok cool. Result %v", res)
	}
}
