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

func fileOpe(filename string) (err error) {
	fn := filename
	err = extensionCheck(fn)
	if err != nil {
		return err
	}
	err = fileEdit(fn)
	if err != nil {
		return err
	}
	return nil
}

func extensionCheck(filename string) (err error) {
	if strings.HasSuffix(filename, ".md") != true {
		return fmt.Errorf("markdown file required")
	}
	return nil
}

func fileEdit(filename string) (err error) {
	fn := filename
	data, err := fileCopy(fn)
	if err != nil {
		return err
	}

	text := fm + string(data)
	err = os.WriteFile(fn, []byte(text), 0666)
	if err != nil {
		return err
	}
	return nil
}

func fileCopy(filename string) (data []byte, err error) {
	fn := filename
	data, err = os.ReadFile(fn)
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}