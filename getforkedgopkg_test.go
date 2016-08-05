package getforkedgopkg

import (
	"flag"
	"log"
	"testing"
)

func TestNotYet(t *testing.T) {
}

func ExampleMain() {
	flag.Parse()

	repoArr := []Repo{
		Repo{
			UpstreamUsername: "go-sql-driver",
			Name:             "mysql",
		},
	}

	for _, r := range repoArr {
		if err := r.Run(); err != nil {
			log.Printf("%v", err.Error())
		}
	}

	// Output:
	// /bin/bash -c '/bin/cd /home/jun/work/src/github.com/junhsieh/mysql && /bin/git clone git@github.com:junhsieh/mysql.git'
	// /bin/bash -c '/bin/cd /home/jun/work/src/github.com/junhsieh/mysql && /bin/git remote add upstream https://github.com/go-sql-driver/mysql.git'
	// /bin/ln -s /home/jun/work/src/github.com/junhsieh/mysql /home/jun/work/src/github.com/go-sql-driver/mysql
}
