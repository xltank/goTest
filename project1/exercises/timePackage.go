package main

import (
	"fmt"
)

func main() {
	/*t := time.NewTicker(time.Millisecond * 100)
	  c := 0
	  for now := range t.C{
	      c ++
	      fmt.Println(now)
	      if c > 10 {
	          t.Stop()
	          break
	      }
	  }*/

	data := 1
	for i := 0; i < 5; i++ {
		go func() {
			t := data + 1
			data = t
		}()
	}
	fmt.Println(data)
}
