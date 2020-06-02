package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

//import (
//	"learnGo/controllers"
//	"net/http"
//)
//
//func main() {
//
//	controllers.RegisterControllers()
//	_ = http.ListenAndServe(":3000", nil)
//}

func main() {
	fmt.Println("Hello world!")

	elliot := &Person{
		Name:                 "Elliot",
		Age:                  24,
		Followers: &SocialFollowers{
			Youtube:              120,
			Twitter:              250,
		},
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshaling error")
	}
	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("Unmarshall error", err)
	}
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.Followers.GetTwitter())
	fmt.Println(newElliot.Followers.GetYoutube())
}
