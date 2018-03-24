package main

import (
	"fmt"
	"net"
	"log"
	"reflect"
	"io"
)

func main()  {

	fmt.Println("fuck client!!")
	addr:="127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("访问公网IP地址是：", conn.RemoteAddr().String())
	fmt.Printf("client addr and port:%v\n", conn.LocalAddr())

	fmt.Println("conn.LocalAddr() 所对应的数据类型是", reflect.TypeOf(conn.LocalAddr().String()))
	fmt.Println("conn.RemoteAddr().String() 所对应的数据类型是", reflect.TypeOf(conn.RemoteAddr().String()))


	n, err :=conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("向服务端发送的数据大小是：", n)

	buf:=make([]byte, 1024) //定义一个切片的长度是1024.

	n, err = conn.Read(buf) // 接收到的内容大小

	if err != nil && err != io.EOF{
		log.Fatal(err)
	}

	fmt.Println(string(buf[:n])) //将接受的内容都读取出来
	conn.Close() // 断开TCP


}
