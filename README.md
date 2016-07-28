# getforkedgopkg
getforkedgopkg will get a forked Go package, set up upstream and create a symbolic link.

\# go run sample.go -githubUsername junhsieh

sample.go:

```go
package main

import (
	"flag"
	"github.com/junhsieh/getforkedgopkg"
	"log"
)

func main() {
	flag.Parse()

	repoArr := []getgopkg.Repo{
		getgopkg.Repo{
			UpstreamUsername: "go-sql-driver",
			Name:             "mysql",
		},
	}

	for _, r := range repoArr {
		if err := r.Run(); err != nil {
			log.Printf("%v", err.Error())
		}
	}
}
```

The commands will be run:

/bin/bash -c '/bin/cd /home/jun/work/src/github.com/junhsieh/mysql && /bin/git clone git@github.com:junhsieh/mysql.git'

/bin/bash -c '/bin/cd /home/jun/work/src/github.com/junhsieh/mysql && /bin/git remote add upstream https://github.com/go-sql-driver/mysql.git'

/bin/ln -s /home/jun/work/src/github.com/junhsieh/mysql /home/jun/work/src/github.com/go-sql-driver/mysql

