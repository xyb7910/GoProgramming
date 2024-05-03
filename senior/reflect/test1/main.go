package test1

import "encoding/json"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		ID:   1,
		Name: "ypb",
		Age:  18,
	}
	s, _ := json.Marshal(u)
	jsonStr := string(s)
	println(jsonStr)
}
