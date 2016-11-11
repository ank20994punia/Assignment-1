package A2Q2

import ("testing"
         "fmt"
         "reflect")

func TestFrequencyCount(t *testing.T) {
        
     
        
       input:=[]string{"go","go","went","go","gone","go"}


       fmt.Println(input)
       
       output:=map[string]int{
           "go" :4,
           "gone" :1,
           "went" :1,
       }
       frequency:=frequency_calculator(input)
      if(reflect.DeepEqual(output,frequency)==false){
          t.Fatal("failed")}
      //if(reflect.DeepEqual(output,frequency)==true){
         // t.Fatal("OK")}
       

}



