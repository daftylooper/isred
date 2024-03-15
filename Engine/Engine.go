// This the main node of the isred engine that performs caching, persistence, has a buffer etc..

package main

import (
	"bufio"
	"fmt"
	"isred/Engine/Cacher"
	"os"
	"strings"
)

type Query struct {
	operation string
	key       string
	value     string
}

// returns operator, key-value
func ParseQuery(query string) (*Query, error) {
	words := strings.Fields(query)
	fmt.Println(words)
	if len(words) == 3 {
		return &Query{words[0], words[1], words[2]}, nil
	} else if len(words) == 2 {
		return &Query{words[0], words[1], ""}, nil
	} else if len(words) == 1 {
		return &Query{words[0], "", ""}, nil
	} else {
		return nil, fmt.Errorf("INVALID QUERY")
	}
}

func EngineLoop(kvs *Cacher.KeyValueStore) {
	for {
		query := ""
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your query: ")
		if scanner.Scan() {
			query += scanner.Text()
		} else {
			fmt.Println("Error reading input:", scanner.Err())
		}
		fmt.Println("--->", query)

		// replace fmt.Println with something that pushes it to a module that handles cli
		switch query, _ := ParseQuery(query); query.operation {
		case "GET":
			out, err := kvs.Get(query.key)
			fmt.Println(out, err)
		case "SET":
			err := kvs.Set(query.key, query.value)
			fmt.Println(err)
		case "SSET":
			err := kvs.SafeSet(query.key, query.value)
			fmt.Println(err)
		case "EXISTS":
			out, err := kvs.Exists(query.key)
			fmt.Println(out, err)
		case "DELETE":
			out, err := kvs.Delete(query.key)
			fmt.Println(out, err)
		case "VIEW":
			out := kvs.View()
			fmt.Println(out)
		}
	}
}

func main() {
	kvs := Cacher.NewKeyValueStore()

	// kvs.Set("allo", "allo")
	// kvs.Set("uncle", "roger")
	// b1, _ := kvs.Exists("qwerty")
	// b2, _ := kvs.Exists("allo")
	// fmt.Println(b1, b2)
	// val, _ := kvs.Get("uncle")
	// fmt.Println(val)

	// fmt.Println("Before Deletion Exists 'allo':", b2)
	// kvs.Delete("allo")
	// b3, _ := kvs.Exists("allo")
	// fmt.Println("After Deletion Exists 'allo':", b3)

	EngineLoop(kvs)

}
