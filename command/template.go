package command

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/gobuffalo/packr"
)

/*
ScriptTemplate script
*/
type ScriptTemplate string

const (
	// ExecuteAnsiblePlaybook ansible script
	ExecuteAnsiblePlaybook ScriptTemplate = "ansible/execute_playbook.sh"
	// SetypePythonEnv setup Python and Ansible
	SetupPythonEnv ScriptTemplate = "ansible/setup_python_env.sh"
)

var box = packr.NewBox("./templates")

/*
Template template as string
*/
func (s *ScriptTemplate) Template() string {
	return string(*s)
}

/*
RenderTemplate render template to file
*/
func RenderTemplate(script ScriptTemplate, params PlaybookParams) (tmpFile *os.File, err error) {
	var templateStr string
	if templateStr, err = box.FindString(script.Template()); err != nil {
		log.Printf("Failed to load template file '%v'", script)
	} else {
		var tmpl *template.Template
		tmpl, err = template.New("test").Parse(templateStr)
		if err != nil {
			return
		}
		if tmpFile, err = ioutil.TempFile("", "execute_ansible.sh"); err != nil {
			log.Printf("Failed to load template file '%v'", script)
		} else {
			var tpl bytes.Buffer
			err = tmpl.Execute(&tpl, params)
			if err != nil {
				return
			}
			_, err = tmpFile.Write(tpl.Bytes())
			if err = tmpl.Execute(&tpl, params); err != nil {
				return
			}
			os.Chmod(tmpFile.Name(), os.ModePerm)
			return
		}
	}
	return
}
