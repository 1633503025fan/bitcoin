package main

import (
	"os"
	"fmt"
)

func main()  {
	args:=os.Args
	length:=len(args)
	fmt.Printf("length:%d\n",length)
	for i,cmd:=range args{
		fmt.Printf("agrs[%d]=%s\n",i,cmd)
	}
}
