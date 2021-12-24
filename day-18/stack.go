package main

type stack []interface{}

func (st *stack) Push(val interface{}) {
	*st = append(*st, val)
}

func (st *stack) Pop() (val interface{}) {
	val = (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return val
}

func (st stack) Size() (size int) {
	return len(st)
}
