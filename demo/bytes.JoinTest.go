package main

import (
	"strings"
	"fmt"
	"bytes"
)

func main()  {
	str:=[]string{"hello","world","!"}
	//func Join(a []string, sep string) string {
	result:=strings.Join(str,"+")
	fmt.Printf("%v\n",result)
	result=strings.Join(str,"")
	fmt.Printf("%v\n",result)

	//func Join(s [][]byte, sep []byte) []byte {
	tmp:=[][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("!"),
	}
	res:=bytes.Join(tmp,[]byte{})
	fmt.Printf("%s\n",res)
}
