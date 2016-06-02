# extract-imports

A program to extract `import` directives of given go files.

# Install

```
git clone
go install
extract-imports...
```

# Usage

```
extract-imports [-h|--help] [-e path,path] source_path

Extract imports directives of the given source directory

To extract directives of each file of the directory 'src'
  extract-imports src/

To exclude some directories of the process
  extract-imports -e path_to_exclude src/

To exclude multiple directories, seperate them by a comma
  extract-imports -e path_to_exclude,path_to_exclude2 src/

Paths are always resolved against cwd.

The resulting output prints [file]: [import]

```

### Example

```
$ bin/extract-imports main.go
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: os
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: fmt
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: flag
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: strings
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: regexp
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: path/filepath
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: go/parser
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/main.go: go/token

$ bin/extract-imports -e main.go main.go
No files found!

$ bin/extract-imports -e fixtures/sub fixtures/
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: os
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: fmt
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: flag
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: strings
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: regexp
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: path/filepath
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: go/parser
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/a.go: go/token

$ bin/extract-imports -e fixtures/a.go fixtures/
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: os
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: fmt
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: flag
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: strings
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: regexp
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: path/filepath
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: go/parser
/home/mh-cbon/goprojects/github.com/mh-cbon/extract-imports/fixtures/sub/b.go: go/token

$ bin/extract-imports -e fixtures/a.go,fixtures/sub/ fixtures/
No files found!
```
