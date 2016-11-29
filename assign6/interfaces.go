package main

import(
"fmt"
)

type Planet interface{
Name() 
Mass() 
}

type PlanetInfo struct{
   pName string
   pMass uint64
}

func(p PlanetInfo) Name() {
   
   fmt.Printf("Name : %s " , p.pName)
}

func(p PlanetInfo) Mass() {
    
    fmt.Printf("Mass :%d " ,p.pMass)
}

func Print(p Planet){
   p.Name()
   p.Mass()
}

func main(){

  earth :=&PlanetInfo{"Earth" , 34000000}
  jupiter :=&PlanetInfo{"Jupiter" , 2500000}
  mars :=&PlanetInfo{"Mars", 1698000}
  venus :=&PlanetInfo{"Venus" , 45679990}
  
   Print(earth)
   fmt.Println("\n")
   Print(jupiter)
   fmt.Println("\n")
   Print(mars)
   fmt.Println("\n")
   Print(venus)
   fmt.Println("\n")

}
