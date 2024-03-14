// This the main node of the isred engine that performs caching, persistence, has a buffer etc..

package main

import (
	"fmt"
	"isred/Engine/Cacher"
)

func main() {
	kvs := Cacher.NewKeyValueStore()

	kvs.Set("allo", "allo")
	kvs.Set("uncle", "roger")
	b1, _ := kvs.Exists("qwerty")
	b2, _ := kvs.Exists("allo")
	fmt.Println(b1, b2)
	val, _ := kvs.Get("uncle")
	fmt.Println(val)

	fmt.Println("Before Deletion Exists 'allo':", b2)
	kvs.Delete("allo")
	b3, _ := kvs.Exists("allo")
	fmt.Println("After Deletion Exists 'allo':", b3)

}
