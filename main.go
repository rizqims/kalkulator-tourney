package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("masukkan ekspresi: ")
	sc.Scan()
	res := sc.Text()

	expFnum := `^(\d+)`
	expExp := `([+\-/*])`
	expSnum := `(\d+)$`

	resFnum, _ := regexp.Compile(expFnum)
	resexp, _ := regexp.Compile(expExp)
	resSnum, _ := regexp.Compile(expSnum)

	ressFnum := resFnum.FindString(res)
	ressexp := resexp.FindString(res)
	ressSnum := resSnum.FindString(res)


	convFnum, _ := strconv.Atoi(ressFnum)
	convSnum, _ := strconv.Atoi(ressSnum)

	fmt.Println("hasilnya adalah", proses(convFnum, convSnum, ressexp))
}

func proses(a int, b int, op string) int{
	switch op{
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	default:
		return 0
	}
}