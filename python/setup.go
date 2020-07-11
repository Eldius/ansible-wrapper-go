package python

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

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

/*
SetupPython set's up the Python environment
*/
func SetupPython(cfg *config.AppConfig) {
	if runtime.GOOS == "linux" {
		log.Println("Cloning pyenv...")
		if repo, err := clone(pyenvRepo, cfg.GetPyenvFolder()); err != nil {
			log.Panic(err.Error())
		} else {
			repo.Fetch(&git.FetchOptions{
				RemoteName: git.DefaultRemoteName,
				Progress:   os.Stdout,
			})
		}
		if repo, err := clone(
			pyenvVirtualenvRepo,
			filepath.Join(cfg.GetPyenvFolder(), "plugins", "pyenv-virtualenv"),
		); err != nil {
			log.Panic(err.Error())
		} else {
			repo.Fetch(&git.FetchOptions{
				RemoteName: git.DefaultRemoteName,
				Progress:   os.Stdout,
			})
		}

	} else if runtime.GOOS == "windows" {
		log.Println("[not implemented yet] Cloning pyenv-win...")
		os.Exit(1)
	}
}
