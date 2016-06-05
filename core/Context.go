package core

type Context interface {
	// get current goroutine cid.
	Cid() int
}

type context int

var __cid int = 100

func NewContext() Context {
	v := context(__cid)
	__cid++
	return v
}

func (v context) Cid() int {
	return int(v)
}




