package main

import (
	data_structures "LT-data-structures/data-structures"
	"fmt"
	"time"
)

func main() {
	//Slice
	//sliceKeyValues := data_structures.NewSliceKeyValues(10000)
	//
	//timeStart := time.Now()
	//sliceKeyValues.Set(5, "foo")
	//timeEnd := time.Since(timeStart)
	//fmt.Println(timeEnd)
	//
	//timeIndexStart := time.Now()
	//_ = sliceKeyValues.KeyValues[10000]
	//timeIndexEnd := time.Since(timeIndexStart)
	//fmt.Println(timeIndexEnd)
	//
	////Map
	//mapKeyValues := data_structures.NewMapKeyValues(100)
	//
	//mapStart := time.Now()
	//mapKeyValues.Set(5, "foo")
	//mapEnd := time.Since(mapStart)
	//
	//fmt.Println(mapEnd)
	//HashMap
	hashMap := data_structures.NewHashMap()
	hashMap.Set(1, "aaa")
	hashMap.Set(3, "bbb")
	hashMap.Set(5, "ccc")
	hashMap.Set(7, "fff")
	hashMap.Set(10, "ggg")
	hashMap.Set(20, "fffggg")
	fmt.Println(hashMap)

	timeHashStart := time.Now()
	hashMap.Get(3)
	timeHashEnd := time.Since(timeHashStart)
	fmt.Println(timeHashEnd)

	key := 5
	for i := 0; i < 10000; i++ {
		key += 10
		hashMap.Set(key, "someValue")
	}
	timeLinkedListStart := time.Now()
	hashMap.Get(9555)
	timeLinkedListEnd := time.Since(timeLinkedListStart)
	fmt.Println(timeLinkedListEnd)

}
