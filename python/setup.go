package python

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"bitbucket.com/Eldius/ansible-wrapper-go/command"
	"bitbucket.com/Eldius/ansible-wrapper-go/config"
	"github.com/go-git/go-git/v5"
)

const (
	pyenvRepo           = "https://github.com/pyenv/pyenv.git"
	pyenvVirtualenvRepo = "https://github.com/pyenv/pyenv-virtualenv.git"
)

func clone(repo string, dest string) (r *git.Repository, err error) {
	r, err = git.PlainClone(dest, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})
	return
}

func pull(repo string, dest string) (r *git.Repository, err error) {
	r, err = git.PlainOpen(dest)
	return
}

func isInstalled(path string) bool {
	i, err := os.Stat(path)
	if err != nil {
		log.Println("Failed to verify previous installation.")
		return false
	}
	return i.IsDir()
}

/*
SetupPython set's up the Python environment
*/
func SetupPython(cfg *config.AppConfig) {
	_ = os.MkdirAll(cfg.WorkspaceFolder(), os.ModePerm)
	if runtime.GOOS == "linux" {
		if isInstalled(cfg.GetPyenvFolder()) {
			log.Println("Pyenv alread installed. We will update Pyenv version.")
			if r, err := pull("", pyenvRepo); err == nil {
				r.Fetch(&git.FetchOptions{
					Force:      true,
					RemoteName: git.DefaultRemoteName,
				})
			}
		} else {
			log.Println("Cloning pyenv...")
			if repo, err := clone(pyenvRepo, cfg.GetPyenvFolder()); err != nil {
				log.Println("Failed to install pyenv...")
				log.Panic(err.Error())
			} else {
				repo.Fetch(&git.FetchOptions{
					RemoteName: git.DefaultRemoteName,
					Progress:   os.Stdout,
				})
			}
		}
		pyenvVirtualenvPath := filepath.Join(cfg.GetPyenvFolder(), "plugins", "pyenv-virtualenv")
		if isInstalled(pyenvVirtualenvPath) {
			log.Println("Pyenv Virtualenv alread installed. We will update Pyenv Virtualenv version.")
			if r, err := pull("", pyenvRepo); err == nil {
				r.Fetch(&git.FetchOptions{
					Force:      true,
					RemoteName: git.DefaultRemoteName,
				})
			}
		} else {
			log.Println("Cloning pyenv-virtualenv...")
			if repo, err := clone(
				pyenvVirtualenvRepo,
				pyenvVirtualenvPath,
			); err != nil {
				log.Println("Failed to install pyenv-virtualenv...")
				log.Panic(err.Error())
			} else {
				repo.Fetch(&git.FetchOptions{
					RemoteName: git.DefaultRemoteName,
					Progress:   os.Stdout,
				})
			}
		}
		command.ExecuteScript(command.SetupPythonEnv, *cfg)
	} else if runtime.GOOS == "windows" {
		log.Println("[not implemented yet] Cloning pyenv-win...")
		os.Exit(1)
	}
}
