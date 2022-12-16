package slice

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Slice struct {
	Sl []any
	itemType reflect.Type
	mu *sync.Mutex
}

func InitSlice() (*Slice) {
	sl := &Slice{}
	sl.mu = &sync.Mutex{}
	return sl
}

func (s *Slice) Add(data any) (index int64, err error) {
	if s.mu == nil {
		return -1, errors.New("First initialize the list")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.Sl) == 0 {
		s.Sl = append(s.Sl, data)
		s.itemType = reflect.TypeOf(data)
		return 0, nil
	}

	dataType := reflect.TypeOf(data)
	if dataType != s.itemType {
		return -1, errors.New(fmt.Sprintf("Can't add an element of %v type to the []%v\n", dataType, s.itemType))
	}

	s.Sl = append(s.Sl, data)
	return int64(len(s.Sl)-1), nil
}

func (s *Slice) Remove(index int64) error {
	if s.mu == nil {
		return errors.New("First initialize the list")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 {
		return errors.New("index must be positive") 
	}

	if index > int64(len(s.Sl)-1) {
		return errors.New("index is out of range")
	}

	s.Sl = append(s.Sl[:index], s.Sl[index+1:]...)
	return nil
}

func (s *Slice) Print() {
	for _, num := range s.Sl {
        fmt.Println(num)
    }
}

func (s *Slice) Len() int64 {
	return int64(len(s.Sl))
}

func (s *Slice) Sort(less func(i, j any) bool) error {
	if s.mu == nil {
		return errors.New("First initialize the list")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.Sl) == 0 {
		return errors.New("Empty slice")
    }
	
	sl := s.Sl
	for k := 0; k < len(s.Sl)-1; k++ {
	  	if less(sl[k], sl[k+1]) {
			sl[k+1], sl[k] = sl[k], sl[k+1]
		}
	}
	return nil
}