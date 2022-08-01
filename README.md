# deptree

show golang dependence like tree

## How to get it

```shell
➜ go install github.com/vc60er/deptree@latest
➜ export PATH=$PATH:$GOPATH/bin
```

## Usage

### params

```
Usage: ./deptree [OPTIONS] [go-mod-graph-output-file]
OPTIONS:
  -d int
    	max depth of dependence (default 3)
```

### example

```shell
➜  redis git:(master) go mod graph | deptree -d 3
package: github.com/go-redis/redis/v9
dependence tree:

┌── github.com/cespare/xxhash/v2@v2.1.2
├── github.com/dgryski/go-rendezvous@v0.0.0-20200823014737-9f7001d12a5f
├── github.com/fsnotify/fsnotify@v1.4.9
│    └── golang.org/x/sys@v0.0.0-20191005200804-aed5e4c7ecf9
├── github.com/nxadm/tail@v1.4.8
│    ├── github.com/fsnotify/fsnotify@v1.4.9
│    │    └── golang.org/x/sys@v0.0.0-20191005200804-aed5e4c7ecf9
│    └── gopkg.in/tomb.v1@v1.0.0-20141024135613-dd632973f1e7
├── github.com/onsi/ginkgo@v1.16.5
│    ├── github.com/go-task/slim-sprig@v0.0.0-20210107165309-348f09dbbbc0
│    │    ├── github.com/davecgh/go-spew@v1.1.1
│    │    └── github.com/stretchr/testify@v1.5.1
│    │         └── ...
│    ├── github.com/nxadm/tail@v1.4.8
│    │    ├── github.com/fsnotify/fsnotify@v1.4.9
│    │    │    └── ...
│    │    └── gopkg.in/tomb.v1@v1.0.0-20141024135613-dd632973f1e7
│    ├── github.com/onsi/gomega@v1.10.1
│    │    ├── github.com/golang/protobuf@v1.4.2
│    │    │    └── ...
│    │    ├── github.com/onsi/ginkgo@v1.12.1
│    │    │    └── ...
│    │    ├── golang.org/x/net@v0.0.0-20200520004742-59133d7f0dd7
│    │    │    └── ...
│    │    ├── golang.org/x/xerrors@v0.0.0-20191204190536-9bdfabe68543
│    │    └── gopkg.in/yaml.v2@v2.3.0
│    │         └── ...
│    ├── golang.org/x/sys@v0.0.0-20210112080510-489259a85091
│    └── golang.org/x/tools@v0.0.0-20201224043029-2b0845dc783e
│         ├── github.com/yuin/goldmark@v1.2.1
│         ├── golang.org/x/mod@v0.3.0
│         │    └── ...
│         ├── golang.org/x/net@v0.0.0-20201021035429-f5854403a974
│         │    └── ...
│         ├── golang.org/x/sync@v0.0.0-20201020160332-67f06af15bc9
│         └── golang.org/x/xerrors@v0.0.0-20200804184101-5ec99f83aff1
├── github.com/onsi/gomega@v1.19.0
│    ├── github.com/golang/protobuf@v1.5.2
│    │    ├── github.com/google/go-cmp@v0.5.5
│    │    │    └── ...
│    │    └── google.golang.org/protobuf@v1.26.0
│    │         └── ...
│    ├── github.com/onsi/ginkgo/v2@v2.1.3
│    │    ├── github.com/go-task/slim-sprig@v0.0.0-20210107165309-348f09dbbbc0
│    │    │    └── ...
│    │    ├── github.com/google/pprof@v0.0.0-20210407192527-94a9f03dee38
│    │    │    └── ...
│    │    ├── github.com/onsi/gomega@v1.17.0
│    │    │    └── ...
│    │    ├── golang.org/x/sys@v0.0.0-20210423082822-04245dca01da
│    │    └── golang.org/x/tools@v0.0.0-20201224043029-2b0845dc783e
│    │         └── ...
│    ├── golang.org/x/net@v0.0.0-20220225172249-27dd8689420f
│    │    ├── golang.org/x/sys@v0.0.0-20211216021012-1d35b9e2eb4e
│    │    ├── golang.org/x/term@v0.0.0-20210927222741-03fcf44c2211
│    │    │    └── ...
│    │    └── golang.org/x/text@v0.3.7
│    │         └── ...
│    ├── gopkg.in/yaml.v2@v2.4.0
│    │    └── gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
│    ├── golang.org/x/sys@v0.0.0-20211216021012-1d35b9e2eb4e
│    ├── golang.org/x/text@v0.3.7
│    │    └── golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
│    └── google.golang.org/protobuf@v1.26.0
│         ├── github.com/golang/protobuf@v1.5.0
│         │    └── ...
│         └── github.com/google/go-cmp@v0.5.5
│              └── ...
├── golang.org/x/net@v0.0.0-20220225172249-27dd8689420f
│    ├── golang.org/x/sys@v0.0.0-20211216021012-1d35b9e2eb4e
│    ├── golang.org/x/term@v0.0.0-20210927222741-03fcf44c2211
│    │    └── golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
│    └── golang.org/x/text@v0.3.7
│         └── golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
├── golang.org/x/sys@v0.0.0-20211216021012-1d35b9e2eb4e
├── golang.org/x/text@v0.3.7
│    └── golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
├── gopkg.in/tomb.v1@v1.0.0-20141024135613-dd632973f1e7
└── gopkg.in/yaml.v2@v2.4.0
     └── gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405

```

