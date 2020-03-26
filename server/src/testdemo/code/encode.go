package main
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 使用encoding/binary
func main(){
	print("this is test result:\n")

	type A struct {
		// should be exported member when read back from buffer
		One int32
		Two int32
	}

	var a A


	a.One = int32(1)
	a.Two = int32(2)

	buf := new(bytes.Buffer)
	fmt.Println("a's size is ",binary.Size(a))
	binary.Write(buf,binary.LittleEndian,a)
	fmt.Println("after write ，buf is:",buf.Bytes())

	var aa A

	binary.Write(buf,binary.LittleEndian,a)
	binary.Read(buf,binary.LittleEndian,&aa)
	fmt.Println("after aa is ",aa)
	f := fmt.Sprintf("%s",[]byte{1,2,3})
	fmt.Print(f)

}