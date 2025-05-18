package main

import (
	"fmt"
	"reflect"
)

func main() {
	sliceExample := []string{"FirstItem", "SecondItem", "ThirdItem"}
	var arrayExample [3]string = [3]string{"FirstItem", "SecondItem", "ThirdItem"}

	fmt.Println("Slice example:", sliceExample)
	fmt.Println("Array example:", arrayExample)

	fmt.Println("Slice TypeOf:", reflect.TypeOf(sliceExample))
	fmt.Println("Array TypeOf:", reflect.TypeOf(arrayExample))

	fmt.Println("Len of slice:", len(sliceExample))
	fmt.Println("Capacity of slice:", cap(sliceExample))

	// When you append a new item to a slice and the underlying array needs to grow,
	// the capacity of the new slice is often twice as big as the old capacity
	sliceExample = append(sliceExample, "FourthItem")
	fmt.Println("Slice after inserting a new item with append:", sliceExample)
	fmt.Println("Slice len after inserting a new item:", len(sliceExample))
	fmt.Println("Slice capacity after inserting a new item:", cap(sliceExample))
	fmt.Println("Slice TypeOf after inserting a new item:", reflect.TypeOf(sliceExample))
}