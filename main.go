package main

import (
	"flag"
	"log"
)

func main(){
	var name string
	flag.StringVar(&name,"name","go语言编程之旅","帮助信息")
	flag.StringVar(&name,"n","go语言编程之旅","帮助信息")
	flag.Parse()
	log.Printf("name %s",name)

}

