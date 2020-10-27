package main

import (
	"bytes"
	"fmt"
)

func main(){
	/**
	*0-9;48-57
	*A-Z;65-90
	*a-z;97-
	 */
	a := []byte("a")//ascii编码0-255
	fmt.Println(a)//[97]
	 b := []byte("b")
	fmt.Println(b)//[98]
	c := []byte("c")
	fmt.Println(c)//[99]

	m := []byte("0")
	fmt.Println(m)//[48]

	rs := bytes.Join([][]byte{a,b,c},[]byte{})
	fmt.Println("拼接后的结果:",rs)//[97,98,99]

	rs1 := bytes.Join([][]byte{a,b,c},m)
	//[97,48,98,48,99]
	fmt.Println("按照m规则拼接后的结果是:",rs1)
}
