// ipc_test.go

package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	method_ := "ECHO: " + method
	res := &Response{method_, params}

	return res
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("From Client1", "test1")
	resp2, _ := client2.Call("From Client2", "test2")

	res1 := resp1.Code + " " + resp1.Body
	res2 := resp2.Code + " " + resp2.Body

	if res1 != "ECHO: From Client1 test1" || res2 != "ECHO: From Client2 test2" {
		t.Error("IpcClient.Call Faild. resp1:", res1, ", resp2:", res2)
	}

	client1.Close()
	client2.Close()
}

