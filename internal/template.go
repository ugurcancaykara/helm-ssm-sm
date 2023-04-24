package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"text/template"
)

var fm template.FuncMap

func ProcessTemplate(tpl *template.Template, valueFile string, ssmFlag bool, smFlag bool, verbose bool, dryRun bool) {
	if valueFile == "" {
		log.Println("File path must be provided")
		return
	}

	content, err := ioutil.ReadFile(valueFile)
	if err != nil {
		log.Fatal(err)
	}

	var fm template.FuncMap

	switch {
	case ssmFlag && smFlag:
		fm = template.FuncMap{
			"ssm": ssmFunc,
			"sm":  smFunc,
		}
	case smFlag:
		fm = template.FuncMap{
			"sm": smFunc,
			"ssm": func(args ...string) string {
				if len(args[0]) < 1 {
					return "error: you didn't provide Secrets Manager secret name, you might be put empty string" + args[0]
				}

				if len(args[1]) < 1 {
					return "error: you didn't provide region" + args[0]
				}
				return "error: you didn't provide --ssm flag despite you've used within the provided values file: " + args[0]
			},
		}
	case ssmFlag:
		fm = template.FuncMap{
			"ssm": ssmFunc,
			"sm": func(args ...string) string {
				if len(args[0]) < 1 {
					return "error: you didn't provide SSM Parameter key, you might be put empty string" + args[0]
				}

				if len(args[1]) < 1 {
					return "error: you didn't provide region" + args[0]
				}
				return "error: you didn't provide --sm flag despite you've used within the provided values file: " + args[0]
			},
		}
	default:
		fm = template.FuncMap{
			"ssm": func(args ...string) string {
				if len(args[0]) < 1 {
					return "error: you didn't provide SSM Parameter key, you may put empty string" + args[0]
				}

				if len(args[1]) < 1 {
					return "error: you didn't provide region" + args[0]
				}
				return "error: you didn't provide --ssm flag despite you've used within the provided values file: " + args[0]
			},
			"sm": func(args ...string) string {
				if len(args[0]) < 1 {
					return "error: you didn't provide SSM Parameter key, you may put empty string" + args[0]
				}

				if len(args[1]) < 1 {
					return "error: you didn't provide region" + args[0]
				}
				return "error: you didn't provide --sm flag despite you've used within the provided values file: " + args[0]
			},
		}
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

	if !dryRun {
		err = WriteFile(valueFile, string(buf.Bytes()))
		if err != nil {
			log.Fatal(err)
		}
		return
	} else {
		log.Println(string(buf.Bytes()))
	}

}

func WriteFile(targetFilePath string, content string) error {
	return ioutil.WriteFile(targetFilePath, []byte(content), 0755)
}
