package cmd

import (
	"fmt"
	"os"
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

	f, err := os.Create(fn)
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}
