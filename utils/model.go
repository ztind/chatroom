package utils

//发送消息model
type Message struct {
	srcIpPort string //本机所在局域网ip:本机程序端口
	destIp string //ip:目标网络地址
	destPort string  //port:目标程序即端口
	msg string //消息
}

func NewMessage(srcIpPort string,destIp string,destPort string,msg string)*Message{
	message := &Message{
		srcIpPort:srcIpPort,
		destIp:destIp,
		destPort:destPort,
		msg:msg,
	}
	return message
}
//获取目标ip:port格式
func (this *Message) DestIPPort()string{
	return this.destIp+":"+this.destPort
}
//获取本机ip:port格式
func (this *Message) SrcIPPort()string{
	return this.srcIpPort
}
//获取消息
func (this *Message)GetMeg()string{
	return this.msg
}