package main

type stack []string

func (st *stack) Push(val string) {
	*st = append(*st, val)
}

func (st *stack) Pop() (val string) {
	val = (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return val
}

func (st stack) Size() (size int) {
	return len(st)
}
