package main

import "fmt"

type instance struct {
	InstansceId string
	Status      string
	Ip          struct{ ip1, ip2 string }
}

func (s *instance) Setstatus(status string) {
	s.Status = status
}
func main() {
	name := &instance{}
	name.InstansceId = "wangsijun"
	name.Status = "normal"
	name.Ip.ip1 = "war"
	name.Ip.ip2 = "star"
	name.Setstatus("error")
	//name.InstansceId = "wangsijun"
	fmt.Println(name)

}
