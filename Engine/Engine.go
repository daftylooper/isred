// This the main node of the isred engine that performs caching, persistence, has a buffer etc..

package Engine

import (
	"fmt"
	"isred/Buffer"
	"isred/Engine/Cacher"
	"strings"
	"time"
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

func EngineLoop(kvs *Cacher.KeyValueStore, buf *Buffer.Buffer, replybuf *Buffer.Buffer) {
	fmt.Println("Starting Up Engine Loop!")
	for {
		query, _ := buf.GetCommand()
		// fmt.Println("->", query)
		// if query not empty( if empty queue dequeued )
		if query != "" {
			// fmt.Println("engine received query:", query)
			switch query, _ := ParseQuery(query); query.operation {
			case "GET":
				out, err := kvs.Get(query.key)
				fmt.Println(err)
				replybuf.PushCommand(out)
			case "SET":
				err := kvs.Set(query.key, query.value)
				fmt.Println(err)
			case "SSET":
				err := kvs.SafeSet(query.key, query.value)
				fmt.Println(err)
			case "EXISTS":
				outbool, err := kvs.Exists(query.key)
				fmt.Println(err)
				out := ""
				if outbool {
					out = query.key + " EXISTS!"
				} else {
					out = query.key + " DOESN'T EXIST!"
				}
				replybuf.PushCommand(out)
			case "DELETE":
				outbool, err := kvs.Delete(query.key)
				fmt.Println(err)
				out := ""
				if outbool {
					out = query.key + " DELETED!"
				} else {
					out = query.key + " COULDN'T DELETE!"
				}
				replybuf.PushCommand(out)
			case "VIEW":
				out := kvs.View()
				fmt.Println(out)
			}
		}
		// doesn't work without thisss??? why??
		time.Sleep(10 * time.Millisecond)
	}
}
