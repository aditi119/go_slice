package main

import ( "fmt"

       "sort" )

func main(){
var y string
var t int
s := make([]int,3,100)
for {

YAY: fmt.Print("\nEnter Integer:")

fmt.Scan(&t)

s = append(s,t)

sort.Ints(s)

fmt.Print("\nSorted slice:", s)

fmt.Print("\nPress X to exit Or Y/y to continue:")

fmt.Scan(&y)

if y=="X" {goto a }else{goto YAY}

a:break

}


}