package memoryCacheMutex

import (
	"sync"
)

type entry struct {
	val   result
	ready chan interface{}
}

type result struct {
	value interface{}
	err   error
}

type Func func(key string)(interface{}, error)

type Memorycache struct {
	f Func
	guard sync.Mutex
	cache map[string]*entry // point is neccerssaly
}

func New(f Func) *Memorycache {
	return &Memorycache{f: f, cache: make(map[string]*entry)}
}

func (this *Memorycache)Get(key string) (interface{}, error) {
	this.guard.Lock()
	e := this.cache[key]
	if e == nil {
		e = &entry{ready: make(chan interface{})}
		this.cache[key] = e
		this.guard.Unlock()

		e.val.value, e.val.err = this.f(key) // slow function, realse lock first
		close(e.ready) // broadcast all
	} else {
		this.guard.Unlock()
		<-e.ready // cause ready has been closed before, so,second time would't block, get 0 immedary
	}

	return e.val.value, e.val.err
}
