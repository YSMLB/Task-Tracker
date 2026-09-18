package main

import(
	"os"
)

func main(){
	if err := cli.Execute(); err != nil{
		os.Exit(1)
	}
}