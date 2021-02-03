package lib

import (
	"bufio"
	"fmt"
	"os"
)

type CmdReader struct{
	reader *bufio.Reader
}

func NewCmdReader()*CmdReader{
	reader := bufio.NewReader(os.Stdin)
	return &CmdReader{reader: reader}
}

func (c* CmdReader)ReadCmd()string{
	lineBytes, _, _ := c.reader.ReadLine()
	return string(lineBytes)
}

func (c* CmdReader)WriteCmd(data string){
	fmt.Println(data)
}