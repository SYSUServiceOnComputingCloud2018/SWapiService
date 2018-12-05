package main

import (
	"fmt"
	"reflect"

	"github.com/peterhellberg/swapi"
)

func main() {
	c := swapi.DefaultClient

	tt, _ := c.Person(1)
	t := reflect.TypeOf(tt)
	v := reflect.ValueOf(tt)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			fmt.Printf("%s %s = %v -tag:%s \n",
				t.Field(i).Name,
				t.Field(i).Type,
				v.Field(i).Interface(),
				t.Field(i).Tag)
		}
	}
}
