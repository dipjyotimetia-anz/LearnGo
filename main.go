package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learnGo/proto"
	"net"
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

//func main() {
//	fmt.Println("Hello world!")
//
//	elliot := &Person{
//		Name:                 "Elliot",
//		Age:                  24,
//		Followers: &SocialFollowers{
//			Youtube:              120,
//			Twitter:              250,
//		},
//	}
//
//	data, err := proto.Marshal(elliot)
//	if err != nil {
//		log.Fatal("Marshaling error")
//	}
//	fmt.Println(data)
//
//	newElliot := &Person{}
//	err = proto.Unmarshal(data, newElliot)
//	if err != nil {
//		log.Fatal("Unmarshall error", err)
//	}
//	fmt.Println(newElliot.GetAge())
//	fmt.Println(newElliot.GetName())
//	fmt.Println(newElliot.Followers.GetTwitter())
//	fmt.Println(newElliot.Followers.GetYoutube())
//}

type Server struct {
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &Server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *Server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *Server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
