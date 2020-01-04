package imp

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"learn.com/rpc/thrift/protocol/thriftrpc"
	"log"
	//cu "ksyun.com/commons/util"
)

type BookServiceImpl struct {
}

func (service *BookServiceImpl) ReadBook(name string, work *thriftrpc.Work) (string, error) {
	//num := cu.RandomInt(50)
	//time.Sleep(200 * time.Millisecond)
	return "hello," + name, nil
}

func StartServer(hostPost string) error {
	serverTransport, err := thrift.NewTServerSocket(hostPost)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	handler := &BookServiceImpl{}
	processor := thriftrpc.NewBookServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	fmt.Println("thrift server in", hostPost)
	err = server.Serve()
	if err != nil {
		return err
	}
	return nil
}
