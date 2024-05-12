package cron

import (
	"context"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron   *cron.Cron
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context) *Scheduler {
	c, cancel := context.WithCancel(ctx)
	return &Scheduler{
		cron:   cron.New(),
		ctx:    c,
		cancel: cancel,
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cancel()
	s.cron.Stop()
}

func (s *Scheduler) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return s.cron.AddFunc(spec, cmd)
}

func (s *Scheduler) AddJob(spec string, cmd cron.Job) (cron.EntryID, error) {
	return s.cron.AddJob(spec, cmd)
}

func (s *Scheduler) Remove(id cron.EntryID) {
	s.cron.Remove(id)
}

func (s *Scheduler) Entries() []cron.Entry {
	return s.cron.Entries()
}

func (s *Scheduler) Run() {
	<-s.ctx.Done()
}

func (s *Scheduler) Context() context.Context {
	return s.ctx
}
