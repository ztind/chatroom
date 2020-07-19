package main

import (
	"fmt"
	"io"
	"my_code/net_sty/chatroom/utils"
	"net"
	"strings"
)

/**
基于tcp协议：
server实现多个client端链接的信息中转，从而实现client间的通信
 */
var (
	mqChan = make(chan *utils.Message,1000) //消息队列channel,message format: ip:port#msg
	onlineConnMap = make(map[string]net.Conn)//存储每个客户端(ip:port)对应的链接
)

func main(){
	server,err := net.ResolveTCPAddr("tcp","127.0.0.1:8080")
	checkError(err)
	socket,err := net.ListenTCP("tcp",server)
	checkError(err)
	defer socket.Close()
	//消费信息
	go consumeMsg()
	//建立链接
	for{
		conn,err := socket.Accept() //阻塞等待client端链接
		ip_port := conn.RemoteAddr().String()
		onlineConnMap[ip_port] = conn
		fmt.Println("build conn : ",ip_port)
		go handlerConn(ip_port,conn,err)
	}
}
func checkError(err error){
	if err!=nil{
		panic(err)
	}
}
func handlerConn(ip_port string,conn net.Conn,err error){
	defer func() {
		//关闭通道，移除map里的ip链接映射
		conn.Close()
		delete(onlineConnMap,ip_port)
		//打印剩余在线链接信息
		for v := range onlineConnMap{
			fmt.Println("online conn info: ",v)
		}
	}()
	checkError(err)
	//read阻塞，循环读取通道数据
	for{
		data := make([]byte,1024)
		n,err := conn.Read(data)
		if n == 0 && err == io.EOF {
			break
		}
		if err!=nil {
			break
		}
		paraseMsg(conn,string(data[0:n]))
	}
}
func paraseMsg(conn net.Conn,msg string){
	//#字号拆封分消息
	content := strings.Split(msg,"#")
	if len(content) > 1{
		ip_port := content[0]
		msg := content[1]
		destIp :=  strings.Split(ip_port,":")[0]
		destPort :=  strings.Split(ip_port,":")[1]
		message := utils.NewMessage(conn.RemoteAddr().String(),destIp,destPort,msg)
		//将读取的数据存入channel
		mqChan <- message
	}else {
		conn.Write([]byte("消息发送失败,请以目标ip:port#msg格式发送"))
	}
}
func consumeMsg(){
	for{
		select {
		case msg := <- mqChan:
			//channel管道有信息
			dest_ip_port := msg.DestIPPort()
			//通过ip:port拿到对应客户度的conn链接管道
			if len(onlineConnMap) > 0 {
				if conn,ok:= onlineConnMap[dest_ip_port];ok{
					//写入消息，回应
					conn.Write([]byte(msg.GetMeg()))
				}else {
					//自己回应自己对方ip:port没有
					if localConn,ok:= onlineConnMap[msg.SrcIPPort()];ok{
						err_msg := dest_ip_port+"没有在线，请检查ip和port是否正确！"
						localConn.Write([]byte(err_msg))
					}
				}
			}else {
				fmt.Println("所有的客户端链断开！")
				break
			}
		}
	}
}
