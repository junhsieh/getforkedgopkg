package getgopkg

import (
	"errors"
	"flag"
	"log"
	"os"
)

const (
	bashBin string = "/bin/bash"
	cdBin   string = "/bin/cd"
	lnBin   string = "/bin/ln"
	gitBin  string = "/bin/git"
)

const (
	githubSSH         string = "git@github.com"
	githubHTTPS       string = "https://github.com"
	githubSrcRelative string = "src/github.com"
)

var (
	githubUsernamePtr *string = flag.String("githubUsername", "junhsieh", "GitHub Username")
	goPath            string  = os.Getenv("GOPATH")
	githubSrc         string  = goPath + "/" + githubSrcRelative
)

type Repo struct {
	UpstreamUsername string
	Name             string
	GoGetURL         string
}

func (r *Repo) GetOriginSSH() string {
	return githubSSH + ":" + *githubUsernamePtr + "/" + r.Name + ".git"
}

func (r *Repo) GetUpstreamHTTPS() string {
	return githubHTTPS + "/" + r.UpstreamUsername + "/" + r.Name + ".git"
}

func (r *Repo) GetSrcDir() string {
	return githubSrc
}

func (r *Repo) GetOriginDir() string {
	return r.GetSrcDir() + "/" + *githubUsernamePtr
}

func (r *Repo) GetUpstreamDir() string {
	return r.GetSrcDir() + "/" + r.UpstreamUsername
}

func (r *Repo) GetCMDBash(cmd string) string {
	return bashBin + " -c '" + cmd + "'"
}

func (r *Repo) GetCMDClone() string {
	return r.GetCMDBash(cdBin + " " + r.GetOriginDir() + "/" + r.Name + " && " + gitBin + " clone " + r.GetOriginSSH())
}

func (r *Repo) GetCMDRemoteAdd() string {
	return r.GetCMDBash(cdBin + " " + r.GetOriginDir() + "/" + r.Name + " && " + gitBin + " remote add upstream " + r.GetUpstreamHTTPS())
}

func (r *Repo) GetCMDLink() string {
	return lnBin + " -s " + r.GetOriginDir() + "/" + r.Name + " " + r.GetUpstreamDir() + "/" + r.Name
}

func (r *Repo) Run() error {
	if r.UpstreamUsername == *githubUsernamePtr {
		return errors.New("GitHub username and upstream username could not be the same")
	}

	// clone
	log.Printf("%v", r.GetCMDClone())

	// add remote
	log.Printf("%v", r.GetCMDRemoteAdd())

	// add symbolic link
	log.Printf("%v", r.GetCMDLink())

	return nil
}
