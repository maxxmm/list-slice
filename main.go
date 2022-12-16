package main

import (
	"education/list/storages/list"
	"education/list/storages/model"
	"education/list/storages/slice"
	"fmt"
	"math/rand"
	"sync"
    "log"
)

var wg sync.WaitGroup

func main() {
    var arr model.Storage

    //test for slice
    arr = slice.InitSlice()
    fillArr(arr)

    //test for list
    arr = list.InitList()
    fillArr(arr)
}

func fillArr(arr model.Storage) {
    var err error
    for i := 0; i < 5000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            _, err = arr.Add(rand.Intn(50000))
            if err != nil {
                log.Fatal(err)
            }
        }()
    }
    wg.Wait()
    fmt.Println("length of the array should be 5000")
    fmt.Printf("it is: %v\n", arr.Len())
}
