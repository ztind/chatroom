# run steps
1.Run the server first
- cd network
- go run server.go
  
2.Then run two clients
- cd client
- go run client.go  
- go run client.go 
  
3.Message sending format<br/>
eg: client1(127.0.0.1:55996) and client1(127.0.0.1:55997)<br/>
message format: ip:port#msg<br/>
- client1 send message to client2:<br/>
127.0.0.1:55997#hello,i`m client1<br/>

- client2 send message to client1:<br/>
127.0.0.1:55996#hello,i`m client2<br/>
