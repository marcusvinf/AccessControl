package terminalactions

import (
	"fmt"
)

const timeFormat = "2006-01-02 15:04:05"

func (t *TerminalData) PersonToTerminal(img []byte) error {
	x := "dale"
	test, err := compressBase64Image(x)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	fmt.Println(timeFormat)
	return nil
}
