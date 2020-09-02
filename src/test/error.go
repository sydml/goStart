package main

import (
	"fmt"
)

func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDeividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dDate := DivideError{
			dividee: varDeividee,
			divider: varDivider,
		}
		errorMsg = dDate.Error()
		return
	} else {
		return varDeividee / varDivider, ""
	}

}
