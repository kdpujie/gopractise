package main

import (
	"fmt"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	it := th.Iter()
	for {
		v, ok := <-it
		if !ok {
			break
		}
		fmt.Printf("%s%v \n", "ch", v)
	}

}
