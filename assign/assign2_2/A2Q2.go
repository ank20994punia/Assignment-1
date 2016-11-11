package A2Q2


func frequency_calculator(words []string) map[string]int{
word_frequency:=make(map[string]int)
for _,word:=range words{
       _,value:=word_frequency[word]
      
        if value==true{
              word_frequency[word]+=1
}else{
              word_frequency[word]=1
}
}
return word_frequency
}





