package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"bitbucket.com/Eldius/ansible-wrapper-go/config"
	"bitbucket.com/Eldius/ansible-wrapper-go/logger"
)

/*
GetCommandExecutionEnvVars generates the env vars to execute
commands
*/
func GetCommandExecutionEnvVars(cfg *config.AppConfig) []string {
	sysPath, _ := os.LookupEnv("PATH")
	newPath := fmt.Sprintf("PATH=%s:%s", cfg.GetPyenvBinFolder(), sysPath)
	workspace := cfg.WorkspaceFolder()
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)
}

/*
ExecuteWithEnv executes a command
*/
func ExecuteWithEnv(command string, execArgs []string, cfg *config.AppConfig, path string, env []string) {
	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path:   command,
		Args:   execArgs,
		Env:    env,
		Stdout: l,
		Stderr: l,
		Dir:    path,
	}

	if err := cmd.Run(); err != nil {
		log.Println("Failed to execute command")
		log.Println(err.Error())
		return
	}
	log.Println("Command finished with success.")
	//cmd.Wait()
	return
}

/*
ExecuteScript executes an ansible playbook
*/
func ExecuteScript(s ScriptTemplate, cfg config.AppConfig) {
	if tmp, err := RenderTemplate(s, PlaybookParams{
		Name: "test.yml",
		Workspace: cfg.WorkspaceFolder(),
	}); err == nil {
		_ = tmp.Close()
		ExecuteWithEnv("/bin/bash", []string{"/bin/bash", "-c", tmp.Name()}, &cfg, cfg.WorkspaceFolder(), GetCommandExecutionEnvVars(&cfg))
		log.Println(tmp.Name())
	} else {
		log.Println("Failed to execute script...")
		log.Panic(err.Error())
	}
}
