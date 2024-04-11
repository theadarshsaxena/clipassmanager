package functions

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func TestAtottoClipboardWorking() error {

	stringToPaste := "myTestString"
	err := clipboard.WriteAll(stringToPaste)
	if err != nil {
		return err
	}

	content, err := clipboard.ReadAll()
	if err != nil {
		return err
	}
	if content != stringToPaste {
		return fmt.Errorf("expected \"%s\", got \"%s\"", stringToPaste, content)
	}
	return nil
}

func CopyToClipboard(s string) error {
	return clipboard.WriteAll(s)
}
