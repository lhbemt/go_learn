package memoryCacheChan

type Func func(key string) (interface{}, error)

type result struct {
	val interface{}
	err error
}

type entry struct {
	res result
	ready chan interface{}
}

type request struct {
	key string
	resp chan<- result
}

type MemoryCacheChan struct{
	requests chan request
}


func New(f Func) *MemoryCacheChan {
	this := &MemoryCacheChan{requests: make(chan request)}
	go this.server(f)
	return this
}

func(this *entry) call(key string, f Func) {
	this.res.val, this.res.err = f(key)
	close(this.ready) // ready
}

func(this *entry) get(resp chan<- result) {
	<- this.ready
	resp <- result{val: this.res.val, err: this.res.err}
}

func(this *MemoryCacheChan) server(f Func) {
	cache := make(map[string]*entry)
	for req := range this.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan interface{})}
			cache[req.key] = e
			go e.call(req.key, f)
		}
		go e.get(req.resp)
	}
}

func(this *MemoryCacheChan) Get(key string) (interface{}, error) {
	resp := make(chan result)
	req := request{key: key, resp: resp}
	this.requests <- req
	res := <- resp
	return res.val, res.err
}

func(this *MemoryCacheChan) Close() {
	close(this.requests)
}

