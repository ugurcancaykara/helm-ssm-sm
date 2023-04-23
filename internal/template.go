package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"ssm": ssmFunc,
	"sm":  smFunc,
}

func ssmFunc(path string) any {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	value := getSSMParam(path)
	return value
}

func smFunc(key string) string {
	//value := getSecret(key)
	return key
}

func processTemplate(tpl *template.Template, valueFile string, ssmFlag bool, smFlag bool, verbose bool) {
	if valueFile == "" {
		log.Println("File path must be provided")
		return
	}

	content, err := ioutil.ReadFile(valueFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create template with specified function map

	tpl = template.Must(template.New("").Funcs(fm).Parse(string(content)))
	//Execute template and output result to stdout

	var buf bytes.Buffer

	if err := tpl.Execute(&buf, nil); err != nil {
		log.Fatal(err)
		return
	}
	if verbose {
		fmt.Println(string(buf.Bytes()))
	}

	err = WriteFile(valueFile, string(buf.Bytes()))
	if err != nil {
		log.Fatal(err)
	}

}

func WriteFile(targetFilePath string, content string) error {
	return ioutil.WriteFile(targetFilePath, []byte(content), 0755)
}
