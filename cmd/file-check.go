package cmd

import (
	"fmt"
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
	return nil
}

func extensionCheck(filename string) (err error) {
	if strings.HasSuffix(filename, ".md") != true {
		return fmt.Errorf("markdown file required")
	}
	return nil
}

func fileEdit(filename string) (err error) {

	return nil
}
