package common

import (
	"fmt"
	"time"
)

type EventInfo struct {
	infotype uint8
	info     interface{}
}

func NewEventInfo(t uint8, info interface{}) *EventInfo {
	return &EventInfo{
		infotype: t,
		info:     info,
	}
}

func (evinfo *EventInfo) Info() interface{} {
	return evinfo.info
}
func (evinfo *EventInfo) InfoType() uint8 {
	return evinfo.infotype
}

type AsyncJob struct {
	*EventInfo

	errch chan error
	t     *time.Timer

	callback func() //called before quit job
}

func NewAsyncJob(jtype uint8, info interface{}, timeout int) *AsyncJob {
	job := &AsyncJob{
		EventInfo: NewEventInfo(jtype, info),
		errch:     make(chan error, 1),
	}
	if timeout > 0 {
		job.t = time.NewTimer(time.Duration(timeout) * time.Millisecond)
	}
	return job
}

func (job *AsyncJob) SetCallback(fn func()) {
	job.callback = fn
}
func (job *AsyncJob) Wait() (err error) {
	defer func() {
		if job.callback != nil {
			job.callback()
		}
	}()
	//blocking (no timeout)
	if job.t == nil {
		err = <-job.errch
		return
	}

	//with timeout
	select {
	case err = <-job.errch:
		job.t.Stop()
	case <-job.t.C:
		err = fmt.Errorf("Async Job timeout")
	}
	return
}

func (job *AsyncJob) Done(err error) {
	job.errch <- err
}
