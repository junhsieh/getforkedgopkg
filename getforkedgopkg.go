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
	githubUsernamePtr = flag.String("githubUsername", "junhsieh", "GitHub Username")
	goPath            = os.Getenv("GOPATH")
	githubSrc         = goPath + "/" + githubSrcRelative
)

// Repo ...
type Repo struct {
	UpstreamUsername string
	Name             string
	GoGetURL         string
}

// GetOriginSSH ...
func (r *Repo) GetOriginSSH() string {
	return githubSSH + ":" + *githubUsernamePtr + "/" + r.Name + ".git"
}

// GetUpstreamHTTPS ...
func (r *Repo) GetUpstreamHTTPS() string {
	return githubHTTPS + "/" + r.UpstreamUsername + "/" + r.Name + ".git"
}

// GetSrcDir ...
func (r *Repo) GetSrcDir() string {
	return githubSrc
}

// GetOriginDir ...
func (r *Repo) GetOriginDir() string {
	return r.GetSrcDir() + "/" + *githubUsernamePtr
}

// GetUpstreamDir ...
func (r *Repo) GetUpstreamDir() string {
	return r.GetSrcDir() + "/" + r.UpstreamUsername
}

// GetCMDBash ...
func (r *Repo) GetCMDBash(cmd string) string {
	return bashBin + " -c '" + cmd + "'"
}

// GetCMDClone ...
func (r *Repo) GetCMDClone() string {
	return r.GetCMDBash(cdBin + " " + r.GetOriginDir() + "/" + r.Name + " && " + gitBin + " clone " + r.GetOriginSSH())
}

// GetCMDRemoteAdd ...
func (r *Repo) GetCMDRemoteAdd() string {
	return r.GetCMDBash(cdBin + " " + r.GetOriginDir() + "/" + r.Name + " && " + gitBin + " remote add upstream " + r.GetUpstreamHTTPS())
}

// GetCMDLink ...
func (r *Repo) GetCMDLink() string {
	return lnBin + " -s " + r.GetOriginDir() + "/" + r.Name + " " + r.GetUpstreamDir() + "/" + r.Name
}

// Run ...
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
