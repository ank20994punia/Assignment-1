package main

import(
"fmt"
"strings"
"bufio"
"os"
"sort"
)

func frequency_calculator(words []string) map[string]int{
word_frequency:=make(map[string]int)
for _,word:=range words{
       _,value:=word_frequency[word]
      //fmt.Println(word_frequency[word])
        if value==true{
              word_frequency[word]+=1
}else{
              word_frequency[word]=1
}
}
return word_frequency
}

func main(){
 
reader :=bufio.NewReader(os.Stdin)
n := map[int][]string{}
        
//var inputString string
fmt.Println("Enter the sentence :")
inputString,_ :=reader.ReadString('\n')
//fmt.Scanln(&inputString)

words:=strings.Fields(inputString)

/*fmt.Println(words[0])
fmt.Println(words[1])
fmt.Println(words[2])*/

frequency:=frequency_calculator(words)

/*for key,value :=range frequency{
     fmt.Println(key,value)
}*/


        var a []int
        for key, value := range frequency {
                n[value] = append(n[value], key)
        }
        for key := range n {
                a = append(a, key)
        }
        sort.Sort(sort.Reverse(sort.IntSlice(a)))
        for _, key := range a {
                for _, s := range n[key] {
                        fmt.Printf("%s - %d\n", s, key)
                }
        }
}
