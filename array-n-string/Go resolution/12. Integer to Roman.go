package main

import (
	"fmt"
)

func intToRoman(num int) string {
    ones := []string{"","I","II","III","IV","V","VI","VII","VIII","IX"}
    tens := []string{"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"}
    hunds := []string{"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"}
    thous := []string{"","M","MM","MMM"}

    return thous[num/1000] + hunds[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}

func main() {
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
