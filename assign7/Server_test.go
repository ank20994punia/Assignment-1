package main

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
)

//Struct for Conn type
type conn struct{
buf []byte
r int
w int
}

func (c *conn)Read(b []byte) (n int , err error){

//if c.r=0 and c.w=0 , it means buffer i empty
    if c.r==c.w{
           bufferEmpty:=errors.New("errBufferEmpty")
           return 0,bufferEmpty
        }
//Copy the value of buffer to byte array b
     n=copy(b,c.buf[c.r:c.w])
//increment the read of connection by n values
     c.r=c.r+n
//set the values of read and write to 0 after the read and write operation has been completed
     if c.r==c.w{
         c.w=0
         c.r=0
}
     return n,nil
}

func (b *conn) Write(p []byte) (n int, err error) {
	// Slide existing data to beginning.
	if b.r > 0 && len(p) > len(b.buf)-b.w {
		copy(b.buf, b.buf[b.r:b.w])
		b.w -= b.r
		b.r = 0
	}

	// Write new data.
	n = copy(b.buf[b.w:], p)
	b.w += n
	if n < len(p) {
		err = errors.New("errWriteFull")
	}
	return n, err
}

//Function for close connection
func (c *conn) Close() error {

	return nil
}

func (c *conn) LocalAddr() net.Addr {
	
	return nil
}
func (c *conn) RemoteAddr() net.Addr {

	return nil
}
func (c *conn) SetDeadline(t time.Time) error {
	
	return nil
}
func (c *conn) SetReadDeadline(t time.Time) error {
	
	return nil
}
func (c *conn) SetWriteDeadline(t time.Time) error {
	
	return nil
}

//Testing for Read() Function
func TestRead(t *testing.T) {

	b := make([]byte, 3)
	b[0] = 0x00
	b[1] = 0x01
	
	c := &conn{buf: b, r: 0, w: 3}

	//expectederr := errors.New("errReadEmpty")
	count, err := c.Read(b)
	if err != nil && count != 3 {
		t.Fatal("failed to read ")
	}
	//fmt.Println("read count is : \n", rCount)
	fmt.Println(b)
}

//Testing of Write() Function
func TestWrite(t *testing.T) {

	b := make([]byte, 3)
	b[0] = 0x00
	b[1] = 0x01
	
	c := &conn{buf: b, r: 0, w: 3}

	count, err := c.Read(b)
	if err != nil && count != 3 {
		t.Fatal("failed to read ")
	}
	fmt.Println(b)

	wCount, err := c.Write(b)
	if err != nil && wCount != 3 {
		t.Fatal("failed to write ")
	}
	fmt.Println(c.buf)
}

//Testing for Check function (gives reponse to a message)
func TestHandleRequest(t *testing.T){

       b := make([]byte, 30)

	c := &conn{buf: b, r: 0, w: 30}

	handleRequest := Handle(c,b)
	if handleRequest != nil {
		t.Fatal("failed to Handle Operation ")
	}else{
         fmt.Println("Connection Successfull")
        }
}

func TestMessage(t *testing.T){

	handleRequest := Message("Hello")
	if handleRequest != "hi" {
		t.Fatal("failed to Handle Operation ")
	}else{
         fmt.Println("Connection Successfull")
        }
}

//func TestHandleRequest(t *testing.T){

        
      
        


/*func TestRequestFailure(t *testing.T) {

	b := make([]byte, 20)

	c := &conn{buf: b, r: 0, w: 0}

	handle := HandleRequest(c)
	if handle == nil {
		t.Fatal("Failure")
	}

}*/
