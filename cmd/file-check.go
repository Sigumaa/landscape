package cmd

import (
	"fmt"
	"os"
	"strings"
)

const text = `---
title: "title"
date: "date"
description: "description"
---`

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

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}
