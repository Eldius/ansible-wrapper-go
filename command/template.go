package command

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gobuffalo/packr"
)

/*
ScriptTemplate script
*/
type ScriptTemplate string

const (
	// ExecuteAnsiblePlaybook ansible script
	ExecuteAnsiblePlaybook ScriptTemplate = "ansible/execute_playbook.sh"
)

var box = packr.NewBox("./templates")

/*
Template template as string
*/
func (s *ScriptTemplate)Template() string {
	return string(*s)
}

/*
RenderTemplate render template to file
*/
func RenderTemplate(script ScriptTemplate) (tmpFile *os.File, err error) {
	var templateStr string
	if templateStr, err = box.FindString(script.Template()); err != nil {
		log.Printf("Failed to load template file '%v'", script)
	} else {
		if tmpFile, err = ioutil.TempFile("", "execute_ansible.sh"); err != nil {
			log.Printf("Failed to load template file '%v'", script)
		} else {
			tmpFile.WriteString(templateStr)
			os.Chmod(tmpFile.Name(), os.ModePerm)
			return
		}
	}
	return
}
