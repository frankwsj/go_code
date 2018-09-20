package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var jsonStr = `

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
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	list()
}

func list() {

	js, err := simplejson.NewJson([]byte(jsonStr))

	if err != nil {

		panic(err.Error())

	}

	personArr, err := js.Get("person").Array()

	fmt.Println(len(personArr))

	// 遍历

	for i, _ := range personArr {

		//fmt.Println(i, v)

		person := js.Get("person").GetIndex(i)

		name := person.Get("name").MustString()

		age := person.Get("age").MustInt()

		email := person.Get("email").MustString()

		fmt.Printf("name=%s, age=%d, email=%s\n", name, age, email)

		// 读取手机号

		phoneNumArr, _ := person.Get("phoneNum").Array()

		for ii, vv := range phoneNumArr {

			fmt.Println(ii, vv)

		}

	}

}
