package main

import(
"fmt"
"io/ioutil"
"io"
"net/http"
"time"
"encoding/json"
"sync"
)

type Lang struct{
Name string
URL string
Bytes int64
Time time.Duration
}

func(w *Lang) crawl(Print func(w *Lang),wg *sync.WaitGroup){
   
    defer wg.Done()
    now:=time.Now()
    resp,_ :=http.Get(w.URL)
    w.Bytes,_=io.Copy(ioutil.Discard,resp.Body)
    t:=time.Since(now)
    w.Time=t
    Print(w)
}

func Print(w *Lang){
  fmt.Printf("\nOutput in GO format for " + w.Name +":")
  fmt.Printf("%v \n",w)
  js,_:=json.Marshal(w)
  fmt.Printf("Output in json format for " + w.Name +":")
  fmt.Printf(string(js) + "\n")
}

func main(){
   
    var wg sync.WaitGroup
    now:=time.Now()
    w:=&Lang{Name: "Python" ,URL: "https://www.python.org/" }
    w1:=&Lang{Name: "Ruby" ,URL: "https://www.ruby-lang.org/en/" }
    w2:=&Lang{Name: "Golang" ,URL: "https://golang.org/" }

    wg.Add(3)
    go w.crawl(Print,&wg)
    go w1.crawl(Print,&wg)
    go w2.crawl(Print,&wg)
    wg.Wait()
    t:=time.Since(now)
    fmt.Println("\nTotal time taken to crawl the data from all the sites :  ",t)
    
}


