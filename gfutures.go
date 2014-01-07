package gfutures

//Holds the future result and error
type Future struct {
	result interface{}
	err    error
	ch     chan interface{}
}

//Retrieve the result or error as returned by the target function.
func (f *Future) Get() (interface{}, error) {
	result, ok := <-f.ch
	if ok {
		switch result.(type) {
		case error:
			f.err = result.(error)
		default:
			f.result = result
		}
		close(f.ch)
	}
	return f.result, f.err
}

// Create a new future to hold the result and error of the target function.
func NewFuture(target func() interface{}) *Future {
	ch := make(chan interface{})
	go func() {
		ch <- target()
	}()
	return &Future{ch: ch}
}
