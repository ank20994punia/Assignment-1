package	 slice


func Frequency_count(userInput int) []int{
digits :=make([]int,10)
i:=userInput

if(i<0){
  i=-1*i
}
for i>0{

digits[i%10]++

i=i/10
}

return digits
}



