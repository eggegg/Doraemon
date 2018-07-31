package cronjob

import (
	"github.com/eggegg/Doraemon/helpers/env"		
	
)

const (
	READY_TO_EXECUTE = "e"
	CLOSE = "c"
)

type controlChan chan string

type fn func(h *env.Dbhandler) error

type Runner struct {
	Controller controlChan
	Error controlChan
	longLived bool
	Executor fn
	h *env.Dbhandler
}

func NewRunner(longLived bool,h *env.Dbhandler, e fn) *Runner {
	return &Runner {
		Controller: make(chan string, 1),
		Error: make(chan string , 1),
		longLived: longLived,
		Executor: e,
		h : h,
	}
}

func (r *Runner)StartExecute()  {
	defer func ()  {
		if !r.longLived{
			close(r.Controller)
			close(r.Error)
		}
	}()

	for {
		select {
		case c:= <- r.Controller :
			if c==READY_TO_EXECUTE {
				err := r.Executor(r.h)
				if err != nil{
					r.Error <-CLOSE
				} 
			}
		case e:=<-r.Error:
			if e==CLOSE{
				return
			}
		}
	}
}

func (r *Runner) StartAll()  {
	r.Controller <- READY_TO_EXECUTE
	r.StartExecute()
}