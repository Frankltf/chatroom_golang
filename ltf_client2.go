package main

import (
	"net"
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8080")
	CheckError3(err)
	defer conn.Close()
	//conn.Write([]byte("hello beifengwang"))
	//fmt.Println("has sent the message")
	go Messagesend2(conn)
	buf:=make([]byte,1024)
	for{
		numOfBytes,err:=conn.Read(buf)
		CheckError3(err)
		fmt.Println("receive server message content:"+string(buf[0:numOfBytes]))
	}
	fmt.Println("client program end!")
}
func Messagesend2(conn net.Conn)  {
	var input string
	for{
		reader:=bufio.NewReader(os.Stdin)
		data,_,_:=reader.ReadLine()
		input=string(data)
		if strings.ToUpper(input)=="exit"{
			conn.Close()
			break
		}
		_,err:=conn.Write([]byte(input))
		if err!=nil{
			conn.Close()
			fmt.Println("client connect failure:"+err.Error())
			break
		}
	}
}
func CheckError3 (err error)  {
	if err!=nil{
		panic(err)
	}
}