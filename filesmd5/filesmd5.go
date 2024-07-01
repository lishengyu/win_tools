// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func fileExists(file string) bool {
	fi, err := os.Stat(file)
	if err == nil {
		if !fi.IsDir() {
			return true
		}
	}
	return false
}

func IsDir(file string) bool {
	fi, err := os.Stat(file)
	if err == nil {
		if fi.IsDir() {
			return true
		}
	}
	return false
}

func getHumanSize(size int64) string {
	if size > 1024 {
		return fmt.Sprintf("%.2f(Kb)\n", float64(size)/1024)
	} else if size > 1024*1024 {
		return fmt.Sprintf("%.2f(Mb)\n", float64(size)/1024/1024)
	}

	return fmt.Sprintf("%d(b)\n", size)
}

func getFileMd5(file string) string {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	defer fd.Close()

	md5h := md5.New()
	io.Copy(md5h, fd)
	return fmt.Sprintf("%x", md5h.Sum(nil))
}

func fileInfo(file string) string {
	fi, err := os.Stat(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}

	var buf string
	buf += fmt.Sprintf("文件名称：%v\r\n", file)
	buf += fmt.Sprintf("文件大小：%v\r\n", getHumanSize(fi.Size()))
	buf += fmt.Sprintf("文件属性：%v\r\n", fi.Mode())
	buf += fmt.Sprintf("文件时间：%v\r\n", fi.ModTime().Format("2006-01-02 15:04:05"))
	buf += fmt.Sprintf("文件MD5：%v\r\n", getFileMd5(file))
	return buf
}

func getFileInfo(path string) string {
	var buf string
	if IsDir(path) {
		err := filepath.WalkDir(path, func(dir string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("filepath walk failed:%v\n", err)
				return err
			}

			if !d.IsDir() {
				buf += fileInfo(dir)
				buf += "\r\n"
			}

			return nil
		})

		if err != nil {
			fmt.Printf("filepath walk failed:%v\n", err)
			return ""
		}
	} else {
		buf += fileInfo(path)
	}

	return buf
}

func getFilesInfo(files []string) string {
	var buff string
	for _, file := range files {
		buff += getFileInfo(file)
		buff += "\r\n"
	}

	return buff
}

func main() {
	var textEdit *walk.TextEdit
	MainWindow{
		Title:   "Md5计算器(syli)",
		MinSize: Size{320, 240},
		Layout:  VBox{},
		OnDropFiles: func(files []string) {
			textEdit.SetText(getFilesInfo(files))
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &textEdit,
				ReadOnly: true,
				HScroll:  true,
				VScroll:  true,
				Text:     "Drop files here, from windows explorer...",
			},
		},
	}.Run()
}
