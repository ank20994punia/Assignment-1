package slice

import ("testing"
         
         "reflect")

func TestDoubleAndPrint(t *testing.T) {
        
     
        
       input:=121234
       
       input_2:=-123
       output:=[]int{0,2,2,1,1,0,0,0,0,0}
       
       output_2:=[]int{0,1,1,1,0,0,0,0,0,0}
       frequency:=Frequency_count(input)
       
       if(reflect.DeepEqual(output,frequency)==false){
          t.Fatal("failed")}

       
       frequency_2:=Frequency_count(input_2)
       if(reflect.DeepEqual(output_2,frequency_2)==false){
          t.Fatal("failed")}

      
       

}



