package main

import (
	"fmt"

	"git.code.oa.com/going/going/json"
)

// A test object
type A struct {
	A  int
	AA int    `json:"aa,string"`
	S  string `json:"s"`
	SS int    `json:"ss,int"`
	B  bool
}

var a = &A{
	A:  1,
	AA: 2,
	S:  "ss",
	B:  true,
}
var s = []byte("{\"a\":11, \"aa\":\"22\"}")

func Example() {
	e := json.Unmarshal([]byte(s), a)

	a.S = "bbbb"
	a.SS = 1234567
	b, e := json.Marshal(a)
	unm := make(map[string]interface{})
	_ = json.Unmarshal(b, &unm)

	fmt.Println(e)
	fmt.Println(a)
	fmt.Printf("%T %v\n", unm["ss"], unm["ss"])
	fmt.Println(string(b))
	//output:
	//hello json
}

func main() {
	//jsonStr := `{"number":1234567}`
	//result := make(map[string]interface{})
	//err := json.Unmarshal([]byte(jsonStr), &result)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(result["number"].(float64))

	Example()
}
