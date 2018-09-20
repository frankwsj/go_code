package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Person []struct {
		Age      int      `json:"age"`
		Name     string   `json:"name"`
		Email    string   `json:"email"`
		PhoneNum []string `json:"phoneNum"`
	} `json:"person"`
}

var jsonstr = `
	  {
           "person": [{
              "name": "piao",
              "age": 30,
              "email": "piaoyunsoft@163.com",
              "phoneNum": [
                  "13974999999",
                  "13984999999"
              ]
           }, {

              "name": "aaaaa",
              "age": 20,
              "email": "aaaaaa@163.com",
              "phoneNum": [
                  "13974998888",
                  "13984998888"
              ]
           }, {
              "name": "bbbbbb",

              "age": 10,

              "email": "bbbbbb@163.com",

              "phoneNum": [

                  "13974997777",

                  "13984997777"

              ]

           }]

       }`

func main() {
	s := Student{}
	json.Unmarshal([]byte(jsonstr), &s)
	fmt.Println(s.Person[0].Name)
	fmt.Println(s.Person[1].Age)

}
