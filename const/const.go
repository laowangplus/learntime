package main

import "fmt"

const (
	i=7
	j
	k
)
// i j k分别等于多少

func main()  {
	defer func() {
		         fmt.Println("defer one")
		     }()
	     defer func() {
		         fmt.Println("defer two")
		     }()
	     defer func() {
		         if info := recover(); info != nil {
			             fmt.Println("catch: ", info)
			         }
		         fmt.Println("defer three")
		     }()
	    panic("panic here")
}
