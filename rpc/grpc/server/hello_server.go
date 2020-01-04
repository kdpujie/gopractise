/**
@description grpc helloworld server示例程序
**/
package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"pujie.org/rpc/grpc/entry"
)

type simpileServer struct {
}

func (this *simpileServer) GetStatus(ctx context.Context, in *entry.Aggregation) (*entry.Aggregation, error) {
	r := in.GetRequest()
	//fmt.Printf("server-r: id=%s, uid=%s, name=%s, ip=%s \n",
	//	r.GetId(), r.GetUid(), r.GetName(), r.GetIp())
	r.Name = r.Name + "-您好"
	res := &entry.Response{}
	res.Id = r.Id
	res.Message = "hello," + r.GetName()
	res.Status = "pass"
	in.Response = res
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", entry.HostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	entry.RegisterGreeterServer(s, &simpileServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Printf("server listen on %s \n", entry.HostPort)
	}
}
