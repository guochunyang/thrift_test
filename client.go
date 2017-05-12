package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"thrift_test/gen-go/tutorial"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "8080"))
	if err != nil {
		fmt.Println("err ....", err)
		return
	}

	useTransport := transportFactory.GetTransport(transport)
	client := tutorial.NewCalculatorServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Println("err ..........")
		return

	}

	defer transport.Close()

	ret, err := client.Add(1, 2)
	fmt.Println("ret: ", ret)
	for i:= 0; i < 10; i++ {
		client.Ping()
	}
}

