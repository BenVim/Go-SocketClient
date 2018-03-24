package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"bufio"
)

func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {

	//if the reader has a WriteTo method, use it to do the copy.
	//Avoids an allocation and a copy.

	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}

	// similarly , if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	if buf == nil {
		buf = make([]byte, 32*1024)
	}

	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err

}

func main() {

	fmt.Println("fuck client!!")
	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("访问公网IP地址是：", conn.RemoteAddr().String())
	fmt.Printf("client addr and port:%v\n", conn.LocalAddr())

	fmt.Println("conn.LocalAddr() 所对应的数据类型是", reflect.TypeOf(conn.LocalAddr().String()))
	fmt.Println("conn.RemoteAddr().String() 所对应的数据类型是", reflect.TypeOf(conn.RemoteAddr().String()))

	n, err := conn.Write([]byte("It avoids these problems of duplication and fission.\n"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("向服务端发送的数据大小是：", n)



	r :=bufio.NewReader(conn)
	for{
		line, err := r.ReadString('\n')
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(line)
	}



	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来
	conn.Close()                 // 断开TCP

}
