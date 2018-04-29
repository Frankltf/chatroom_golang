package main
import (
	"net"
	"fmt"
	"strings"
)
var onlineConns=make(map[string]net.Conn)
var messageQueue=make(chan string,1000)
var quitChan=make(chan bool)
func main() {
	listen_socker,err:=net.Listen("tcp","127.0.0.1:8080")
	CheckError(err)
	defer listen_socker.Close()
	fmt.Println("server is runing")
	go consumermesage()
	for{
		conn,err:=listen_socker.Accept()
		CheckError(err)
		addr:=fmt.Sprintf("%s",conn.RemoteAddr())
		onlineConns[addr]=conn
		for i:=range onlineConns{
			fmt.Print(i)
			fmt.Print("\n")
		}
		go ProcessInfo(conn)
	}
}
func CheckError (err error)  {
	if err!=nil{
		panic(err)
	}
}
func ProcessInfo(conn net.Conn)  {
	buf:=make([]byte,1024)
	defer conn.Close()
	for{
		numOfBytes,err:=conn.Read(buf)
		if err!=nil{
			break
		}
		if numOfBytes != 0{
			message:=string(buf[0:numOfBytes])
			messageQueue<-message
			remoteAddr:=conn.RemoteAddr()
			fmt.Print(remoteAddr)
			fmt.Printf("Has received this message:%s\n",string(buf))
		}
	}
}
func consumermesage()  {
	for{
		select {
			case message:=<-messageQueue:
				//对消息进行解析
				doProcessMessage(message)
			case <-quitChan:
				break

		}
	}
}
func doProcessMessage(message string)  {
	contens:=strings.Split(message,"#")
	if len(contens)>1{
		addr:=contens[0]
		sendMessage:=contens[1]
		addr=strings.Trim(addr," ")
		if conn,ok:=onlineConns[addr];ok{
			_,err:=conn.Write([]byte(sendMessage))
			if err!=nil{
				fmt.Print("shi bai")
			}
		}
	}
}