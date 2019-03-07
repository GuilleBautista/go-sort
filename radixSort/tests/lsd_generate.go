package main

import (
	"fmt"
	lsd "github.com/commonhater/go-sort/radixSort"
	"math/rand"
	"time"
	"flag"
)

func main(){
	var array []int
	
	totalPtr := flag.Int("total", 0, "number of elements generated")
	maxPtr := flag.Int("max", 0, "max number generated")
	
	flag.Parse()

	if( *totalPtr!=0 && *maxPtr!=0){
		for i:=0;i<*totalPtr;i++{
			array=append(array, rand.Intn(*maxPtr))
		}
	}else{
		for i:=0;i<10000000;i++{
			array=append(array, rand.Intn(1000000))
		}
	}
	//fmt.Println(array)

	start:=time.Now()
	anotherArray:=lsd.Radix(array[:])
	t:=time.Now()

	fmt.Println(anotherArray)
	fmt.Println(t.Sub(start))
}
