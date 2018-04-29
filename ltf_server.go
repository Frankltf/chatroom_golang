package main
import (
	"net"
	"fmt"
)
func main() {
	listen_socker,err:=net.Listen("tcp","127.0.0.1:8080")
	CheckError(err)
	defer listen_socker.Close()
	fmt.Println("server is runing")
	for{
		conn,err:=listen_socker.Accept()
		CheckError(err)
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
			remoteAddr:=conn.RemoteAddr()
			fmt.Print(remoteAddr)
			fmt.Printf("Has received this message:%s\n",string(buf))
		}
	}
}