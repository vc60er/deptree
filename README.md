# deptree

show golang dependence like tree

## How to get it

```shell
➜ go get -u github.com/vc60er/deptree
➜ export PATH=$PATH:$GOPATH/bin
```

## Usage

```shell
➜ go mod graph | deptree
package: github.com/json-iterator/go
dependence tree:

┌── github.com/davecgh/go-spew@v1.1.1
├── github.com/google/gofuzz@v1.0.0
├── github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
├── github.com/modern-go/reflect2@v1.0.2
└── github.com/stretchr/testify@v1.3.0
     ├── github.com/davecgh/go-spew@v1.1.0
     ├── github.com/pmezard/go-difflib@v1.0.0
     └── github.com/stretchr/objx@v0.1.0
```
