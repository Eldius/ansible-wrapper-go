package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	workspace := cfg.Workspace
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)
}

/*
ExecuteWithEnv executes a command
*/
func ExecuteWithEnv(command string, execArgs []string, cfg *config.AppConfig, path string, env []string)  {
	path, err := filepath.Abs("./command/files/source.sh")
	if err != nil {
		log.Println("Failed to parse source file path")
		log.Println(err.Error())
		return
	}
	extraFile, err := os.Open(path)
	if err != nil {
		log.Println("Failed to opem source file")
		log.Println(err.Error())
		return
	}
	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path: path,
		Args: execArgs,
		Env:  env,
		Stdout: l,
		Stderr: l,
		Dir: cfg.Workspace,
		ExtraFiles: []*os.File{extraFile},
	}

	if err := cmd.Run(); err != nil {
		log.Println("Failed to execute command")
		log.Println(err.Error())
		return
	}
	log.Println("Command finished with success.")
	cmd.Wait()
}

/*
Execute executes a command
*/
func Execute(command string, path string, cfg *config.AppConfig)  {
	ExecuteWithEnv(command, []string{}, cfg, path, GetCommandExecutionEnvVars(cfg))
}

/*
ExecuteWithArgs executes a command with args
*/
func ExecuteWithArgs(command string, args []string, path string, cfg *config.AppConfig)  {
	log.Println("executing:", command, args, "at", path)
	ExecuteWithEnv(command, args, cfg, path, GetCommandExecutionEnvVars(cfg))
}

/*
ExecutePyenvCmd executes a command with args
*/
func ExecutePyenvCmd(args []string, cfg *config.AppConfig)  {
	log.Println("executing: pyenv", strings.Join(args, " "))
	execArgs := []string{"pyenv"}
	execArgs = append(execArgs, args...)
	ExecuteWithEnv(cfg.GetPyenvBinFolder() + "/pyenv", execArgs, cfg, cfg.GetPyenvBinFolder(), GetCommandExecutionEnvVars(cfg))
}
