package pool

import (
	"go.uber.org/zap"
)

type Pool[T any] struct {
	sem    chan *T
	logger *zap.Logger
}

func NewPool[T any](size int, logger *zap.Logger, factory func() *T) *Pool[T] {
	sem := make(chan *T, size)
	for i := 0; i < size; i++ {
		sem <- factory()
	}
	return &Pool[T]{
		sem:    sem,
		logger: logger,
	}
}

func (p *Pool[T]) Get() *T {
	s := <-p.sem
	p.logger.Sugar().Debug("Got connection from pool")
	return s
}

func (p *Pool[T]) Put(conn *T) {
	p.logger.Sugar().Debug("Putting connection back in pool")
	p.sem <- conn
}
