package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	x := sc.Text()
	
	sc.Scan()
	y := sc.Text()

	fmt.Println(x, y)

}
