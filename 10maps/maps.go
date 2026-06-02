package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps := make(map[string]string)
	// maps["name"] = "yug"
	// maps["age"] = "18"
	// fmt.Println(maps)
	// fmt.Println(maps["name"])
	// fmt.Println(maps["age"])
	// fmt.Println(len(maps))
	// delete(maps, "age")
	// fmt.Println(maps)
	// clear(maps)
	// fmt.Println(maps)
	// delete(maps, "name")
	// fmt.Println(maps)
	// maps := map[string]string{"name": "yug", "age": "18"}
	// ok := maps["name"]
	// if ok == "yug" {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("all not ok")
	// }
	m1 := map[string]string{"name": "yug", "age": "18"}
	m2 := map[string]string{"name": "yug", "age": "18"}
	fmt.Println(maps.Equal(m1, m2))
}
