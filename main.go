package main

import (
	"fmt"
	"github.com/yeziyitao/utils/shell"
)

func main(){
	fmt.Println(shell.Exec("ls -lh"))
	fmt.Println(shell.Exec("lss -lh"))
}
