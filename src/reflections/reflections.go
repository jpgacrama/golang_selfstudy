// Golang program to illustrate
// reflect.ArrayOf() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {
	t := reflect.TypeOf(5)

	// use of ArrayOf method
	arr := reflect.ArrayOf(4, t)
	inst := reflect.New(arr).Interface().(*[4]int)

	for i := 1; i <= 4; i++ {
		inst[i-1] = i * i
	}

	fmt.Println(inst)
}
