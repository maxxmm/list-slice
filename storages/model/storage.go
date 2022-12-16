package model

type Storage interface {
	Add(data any) (index int64, err error)
	Remove(index int64) (err error)
	Print()
	Sort(func(i, j any) bool) error
	Len() int64
}
