package list

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type List struct {
	itemType  reflect.Type
	len       int64
	firstNode *Node
	mutex     *sync.Mutex
}

type Node struct {
	index    int64
	Data     any
	nextNode *Node
}

func InitList() (*List) {
	l := &List{}
	l.mutex = &sync.Mutex{}
	return l
}

func (l *List) Add(data any) (index int64, err error) {
	if l.mutex == nil {
		return -1, errors.New("First initialize the list")
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	n := &Node{}
	if l.firstNode == nil {
		n.index = 0
		n.Data = data
		l.firstNode = n
		l.len = 1
		l.itemType = reflect.TypeOf(data)
		return n.index, nil
	}
	
	lastNode, err := l.Retrieve(l.len - 1)

	if err != nil {
		return -1, err
	}

	dataType := reflect.TypeOf(data)
	if dataType != l.itemType {
		return -1, errors.New(fmt.Sprintf("Can't add an element of %v type to the [list]%v\n", dataType, l.itemType))
	}

	n.Data = data
	lastNode.nextNode = n
	n.index = l.len
	l.len++
	return n.index, nil
}

func (l *List) Print() {
	if l.mutex == nil {
		fmt.Println("First initialize the list")
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.firstNode == nil {
		fmt.Println("Empty list")
		return
	}
	cn := l.firstNode
	for {
		if cn.nextNode == nil {
			fmt.Println(cn.Data)
			break
		}
		fmt.Println(cn.Data)
		cn = cn.nextNode
	}
}

func (l *List) Len() int64 {
	return l.len
}

func (l *List) Remove(index int64) (err error) {
	if l.mutex == nil {
		return errors.New("First initialize the list")
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index == 0 {
		// remove the last item in the list
		newNextNode, err := l.Retrieve(index + 1)
		if err != nil {
			l.len = 0
			l.firstNode = nil
			l.itemType = nil
			return nil
		}

		l.firstNode = newNextNode
		l.len--
		return nil
	}

	prevNode, err := l.Retrieve(index - 1)
	if err != nil {
		return err
	}

	if index == l.len-1 {
		prevNode.nextNode = nil
		l.len--
		return err
	}

	newNextNode, err := l.Retrieve(index + 1)
	if err != nil {
		return err
	}

	prevNode.nextNode = newNextNode
	l.len--
	return nil
}

func (l *List) Retrieve(index int64) (*Node, error) {
	if l.mutex.TryLock() {
		l.mutex.Lock()
		defer l.mutex.Unlock()
	}
	if index > l.len-1 {
		return nil, errors.New("index is out of range")
	}
	if index < 0 {
		return nil, errors.New("index must be positive")
	}
	cn := l.firstNode
	for {
		if cn.index == index {
			return cn, nil
		}
		cn = cn.nextNode
	}
}

func (l *List) Sort(less func(i, j any) bool) error {
	if l.mutex == nil {
		return errors.New("First initialize the list")
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.firstNode == nil {
		return errors.New("Empty list")
	}

	for i := int64(0); i < l.len-1; i++ {
		currentNode := l.firstNode
		for {
			if currentNode.nextNode == nil {
				break
			}
			if less(i, i+1) {
				currentNode.nextNode.Data, currentNode.Data = currentNode.Data, currentNode.nextNode.Data
			}
			currentNode = currentNode.nextNode
		}
	}
	return nil
}
