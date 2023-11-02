package main

import "sync"

type Pool[T any] struct {
	sync.Mutex
	inst []T
	new  func() T
}

func NewPool[T any](new func() T) *Pool[T] {
	return &Pool[T]{
		new: new,
	}
}

func (p *Pool[T]) Borrow() (res T) {
	p.Lock()
	defer p.Unlock()
	if len(p.inst) == 0 {
		return p.new()
	}
	res = p.inst[len(p.inst)-1]
	p.inst = p.inst[:len(p.inst)-1]
	return
}

func (p *Pool[T]) Put(v T) {
	p.Lock()
	defer p.Unlock()
	p.inst = append(p.inst, v)
}
