package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
         "io"
	"net/http"
        "encoding/json"
         "time"
)

type Lang struct {
    Name string
    Url string
    Bytes  int64
    Time time.Duration
}

func(l Lang) crawl(c chan Lang){
 t:=time.Now()
 res,err:=http.Get(l.Url)
 if err!=nil{
  fmt.Println(err)
  }
  defer res.Body.Close()

  body,_:=ioutil.ReadAll(res.Body)
  l.Bytes,_ = io.Copy(ioutil.Discard,bytes.NewReader(body))

  l.Time=time.Since(t)
 
  fmt.Println("\nJson formatted output :")
  jFormat, _ := json.Marshal(l)
  fmt.Printf(string(jFormat)+"\n")
   
   c<-l
}
func main(){

 t:=time.Now()
 getCount:=make(chan Lang)
 var count,countBytes int64

 w := Lang{Name:"Python",Url:"https://www.python.org/"}
 w1 := Lang{Name:"Ruby",Url:"https://www.ruby-lang.org/en/"}
 w2 := Lang{Name:"Golang",Url:"https://golang.org/"}

 go w.crawl(getCount)
 go w1.crawl(getCount)
 go w2.crawl(getCount)
 

 count=0
 countBytes=0
 for a:=range getCount{
 countBytes+=a.Bytes
 count++
 if count==3{
close(getCount)
}
}
fmt.Println(time.Since(t))
fmt.Println("Total number of Bytes read : ",countBytes)

 
}
