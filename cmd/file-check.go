package cmd

import (
	"fmt"
	"os"
	"strings"
)

const fm = `---
title: "title"
date: "date"
description: "description"
---
`

func extensionCheck(filename string) (err error) {
	if strings.HasSuffix(filename, ".md") != true {
		return fmt.Errorf("markdown file required")
	}
	return nil
}

func fileEdit(filename string) (err error) {
	fn := filename
	err = extensionCheck(fn)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	str := ""
	data := []byte(str)
	cnt, err := f.Read(data)
	if err != nil {
		return err
	}
	text := fm + string(data[:cnt])

	return fmt.Errorf(text)
}
