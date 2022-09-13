package helper

import "fmt"

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
