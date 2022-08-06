package cmd

import (
	"fmt"
	"strings"
)

func extensionCheck(filename string) (err error) {
	if strings.HasSuffix(filename, ".md") != true {
		return fmt.Errorf("markdown file required")
	}
	return nil
}

func fileCheck(filename string) (err error) {
	fn := filename
	err = extensionCheck(fn)
	if err != nil {
		return err
	}
	return nil
}
