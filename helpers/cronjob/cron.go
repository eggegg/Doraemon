package cronjob

import (
	"time"
	"github.com/eggegg/Doraemon/helpers/env"		
)

const TASK_RUN_INTERVAL = 300

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker()  {
	//  run once
	go w.runner.StartAll()
	//run periodlly
	for {
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start(h *env.Dbhandler)  {
	r:=NewRunner(true,h, AdInitialLoadExecutor)
	w:=NewWorker(TASK_RUN_INTERVAL, r)
	go w.startWorker()
}