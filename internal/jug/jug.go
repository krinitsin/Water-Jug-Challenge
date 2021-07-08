package jug

import "errors"

var errNotEnoughSpace = errors.New("not enough space in target jug")

type Jug interface {
	Name() string
	Fill()
	Empty()
	CurAmount() int
	Size() int
	setCurAmount(int)
}

type jug struct {
	name string
	curAmount int
	size      int
}

func New(size int,name string) Jug {
	return &jug{size: size,name: name}
}

func (j *jug) Name() string {
	return j.name
}

func (j *jug) Fill() {
	j.curAmount = j.size
	return
}

func (j *jug) Empty() {
	j.curAmount = 0
	return
}

func (j *jug) CurAmount() int {
	return j.curAmount
}

func (j *jug) Size() int {
	return j.size
}

func (j *jug) setCurAmount(v int) {
	j.curAmount = v
	return
}

func Transfer(from, to Jug) {
	trAmount := to.Size() - to.CurAmount()
	if from.CurAmount() < trAmount {
		trAmount = from.CurAmount()
	}
	to.setCurAmount(to.CurAmount() + trAmount)
	from.setCurAmount(from.CurAmount() - trAmount)

	return
}
