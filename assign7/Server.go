package main

import(
"fmt"
"net"
"errors"
)

func main(){
	ln,err:=net.Listen("tcp",":8082")

	
   	if err!=nil{
      		panic(err)
   	}

	

   	defer ln.Close()

   	for{
		
     		conn,err:=ln.Accept()
		
	     	if err!=nil{
	       		panic(err)
		}
		
	
         go HandleRequest(conn)
}
      
}

func HandleRequest(conn net.Conn) error{
  //Make a buffer to hold incoming data.
buf := make([]byte, 1024)
 for{
   
  // Read the incoming connection into the buffer.
      err1 :=Handle(conn,buf)
      if err1 != nil {

      return errors.New("error")
  }
    
}
return nil
}


func Handle(conn net.Conn, buf []byte) error{
   bytes, err := conn.Read(buf)

   if err != nil {

      fmt.Println("Error reading:", err.Error())
  }
  // Send a response back to person contacting us.
   message := string(buf[:bytes-2])
   res:=Message(message)
   _,err2:=conn.Write([]byte(res))
if err2 != nil {

      fmt.Println("Error", err.Error())
  }

   
return nil
}

func Message(message string) string{

   if message=="Hello"{

   message="hi\n"

   }else if message=="Can I know your name ?"{

   message="Yes sure , myself Ankita\n"

   }else if message=="City"{

   message ="Delhi\n"

   }else{

   message ="Invalid\n"
}
return message
}



