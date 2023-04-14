package math

type Math struct {
	A, B, C int
}

func NewMath(a, b, c int) Math {
	return math{A: a, B: b, C: c}
}

func (m Math) Add() int {
	return m.A + m.B
}
