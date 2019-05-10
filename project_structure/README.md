# project_structure

## pre-req

Typically you would start a folder to work on your go source.

The command below temporarily will set GOPATH to that folder so that it can
become the workspace.

```
export GOPATH:$PWD
```

This will build the binary

```
go build
```

This will install any packages into lib

```
go install
```

After you've run the commands above, expect the following tree on disk:

```
.
├── bin
│   └── hello
├── pkg
│   └── linux_amd64
│       └── github.com
│           └── gsvolt
│               └── gotour
│                   └── project_structure
│                       └── string.a
├── README.md
└── src
    └── github.com
        └── gsvolt
            └── gotour
                └── project_structure
                    ├── hello
                    │   └── hello.go
                    └── string
                        └── string.go

14 directories, 5 files
```

## ref:

- https://www.youtube.com/watch?v=XCsL89YtqCs


