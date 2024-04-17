package common

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "common"})
}

type Timer interface {
	Stop()
	Start() //Start or Reset timer
}

type ueTimer struct {
	d       time.Duration
	e       Worker
	abort   chan bool
	_t      *time.Timer
	expired bool
	fn      func() //callback
	wg      sync.WaitGroup
}

// create an idling timer
func NewTimer(d time.Duration, fn func(), e Worker) Timer {
	return &ueTimer{
		d:       d,
		fn:      fn,
		e:       e,
		expired: false,
	}
}

//start/restart the timer
func (t *ueTimer) Start() {
	//make sure the timer is stopped
	t.Stop()

	t.abort = make(chan bool)
	t._t = time.NewTimer(t.d) //always create a new one
	t.expired = false

	t.wg.Add(1)
	task := func() {
		defer t.wg.Done()
		select {
		case <-t.abort:
			//just exit
		case <-t._t.C:
			t.expired = true
			//execute callback function
			if t.fn != nil {
				t.fn()
			}
		}
	}
	if t.e != nil { //submit to workerpool
		t.e.Submit(task)
	} else {
		go task() //create new go routine
	}

}

func (t *ueTimer) Stop() {
	if t.abort != nil { //make the call idempotem
		close(t.abort) //trigger canceling
		t.wg.Wait()    //make sure the running goroutine to complete

		//stop internal timer if it was not expired
		if t.expired {
			t._t.Stop()
		}
		t.abort = nil //make the call idempotem
	}
}
