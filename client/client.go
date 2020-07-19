package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

//客户端
func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8080")
	checkError(err)
	defer conn.Close()//关闭链接
	go readMsg(conn)

	for{
		//监听输入，将输入发送给服务端
		fmt.Println("input>>")
		r := bufio.NewReader(os.Stdin)
		data,_,_ := r.ReadLine()
		conn.Write(data)
		if strings.ToUpper(string(data)) == "EXIT" {
			break //跳出循环，结束程序
		}
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func readMsg(conn net.Conn){
	for{
		data := make([]byte,1024)
		n,err := conn.Read(data)
		if n==0 && err==io.EOF {
			break
		}
		if err !=nil{
			break
		}
		fmt.Println("receive : ",n,err,string(data[0:n]))
	}
}