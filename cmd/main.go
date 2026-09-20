package main

import(
	"os"

	"task-tracker/internal/cli"
)

func main(){
	if err := cli.Run(); err != nil{
		os.Exit(1)
	}
}