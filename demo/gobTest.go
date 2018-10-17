package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"fmt"
)

type Person struct {
	Name string
	Age int
}
func main()  {

	var xiaoming Person
	xiaoming.Name="小明"
	xiaoming.Age=10
	//编码的数据放到buffer中
	var buffer bytes.Buffer
	//func NewEncoder(w io.Writer) *Encoder {
	//使用gob进行序列化（编码）得到字节流
		//1、定义一个编码器
		//2、使用编码器编码
	Encoder:=gob.NewEncoder(&buffer)
	err:=Encoder.Encode(&xiaoming)
	if err!=nil{
		log.Panic("编码出错")
		os.Exit(1)
	}
	fmt.Printf("编码后的小明:%v\n",buffer.Bytes())

	//使用gob进行反序列化（解码）得到Person结构体
		//1、定义一个解码器
		//2、使用解码器解码
		var daMing Person
	//func NewDecoder(r io.Reader) *Decoder {
	decoder:=gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	err=decoder.Decode(&daMing)
	if err!=nil{
		log.Panic("解码出错")
		os.Exit(1)
	}
	fmt.Printf("解码后:%v\n",daMing)

}
