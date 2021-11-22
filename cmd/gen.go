package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

//go:embed template/*
var templateDir embed.FS

type configNode struct {
	Name       string
	KeyName    string
	FieldName  string
	StructName string
	Children   map[string]*configNode
}

func printNode(n *configNode, indent string) {
	fmt.Println(indent, n.Name, ":", n.KeyName)
	if len(n.Children) > 0 {
		indent = strings.ReplaceAll(indent, "--", "  ")
		indent += "|--"
		for _, v := range n.Children {
			printNode(v, indent)
		}
	}
}

func firstLower(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}

func generate(dir string) error {
	root := &configNode{Name: "Root", Children: make(map[string]*configNode)}
	all := viper.AllKeys()
	sort.Strings(all)
	for _, k := range all {
		// skip imports key
		if k == "pp_imports" {
			continue
		}
		current := root
		ss := strings.Split(k, ".")
		for i, p := range ss {
			if current.Children[p] != nil {
				current = current.Children[p]
			} else {
				StructName := ""
				KeyName := ""
				FieldName := strings.Title(p)
				FieldName = strings.ReplaceAll(FieldName, "_", "")

				if i == len(ss)-1 {
					KeyName = strings.Join(ss[:i+1], ".")
				} else {
					v := strings.Join(ss[:i+1], " ")
					v = strings.Title(v)
					v = strings.ReplaceAll(v, " ", "")
					StructName = v
					StructName = strings.ReplaceAll(StructName, "_", "")
				}

				new := &configNode{
					Name:       strings.Title(p),
					FieldName:  FieldName,
					StructName: firstLower(StructName),
					KeyName:    KeyName,
					Children:   make(map[string]*configNode),
				}
				current.Children[p] = new
				current = new
			}
		}
	}

	fmt.Println("")
	fmt.Println(">>>> CONFIG LIST <<<<")
	printNode(root, "|--")

	funcMap := template.FuncMap{
		"ToTitle": strings.Title,
	}

	tmpl, err := template.New("config.tmpl").Funcs(funcMap).ParseFS(templateDir, "template/config.tmpl")
	if err != nil {
		return err
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, root)

	if err != nil {
		return err
	}

	var buf []byte
	if buf, err = format.Source(b.Bytes()); err != nil {
		return err
	}

	if err = ioutil.WriteFile(dir+"/config.go", buf, 0644); err != nil { //nolint
		return err
	}
	return nil
}

func viperLoad(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	var dir string
	if fi.IsDir() {
		dir = path
	} else {
		dir = filepath.Dir(path)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	viper.AddConfigPath(dir)
	for _, file := range files {
		fileName := file.Name()
		ext := filepath.Ext(fileName)
		if len(ext) < 2 {
			continue
		}
		if !stringInSlice(ext[1:], viper.SupportedExts) {
			continue
		}

		viper.SetConfigFile(dir + "/" + fileName)
		err := viper.MergeInConfig()
		if err != nil {
			return "", err
		}
	}
	return dir, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
