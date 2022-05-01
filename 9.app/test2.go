package main

import (
	"fmt"
)

type Example struct {
	Id   []int
	Name []string
}

func (data *Example) AppendExample(id int, name string) {
	data.Id = append(data.Id, id)
	data.Name = append(data.Name, name)
}

var MyMap map[string]Example

func main() {

	var object Example
	object.AppendExample(1, "SomeText")

	MyMap = make(map[string]Example)
	MyMap["key1"] = object
	fmt.Println(MyMap)
}
