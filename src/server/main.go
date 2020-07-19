package main

import (
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

const addr = "127.0.0.1:9090"

func init()  {
	InitBigSetIf()
}

func main()  {

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket: ",err)
		os.Exit(1)
	}
	handler := &GenericServiceHandler{}
	processor := myGeneric.NewTGenericServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", addr)
	server.Serve()
}