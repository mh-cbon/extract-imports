package main

import (
  "os"
  "fmt"
  "flag"
  "strings"
  "regexp"
  "path/filepath"
  "go/parser"
  "go/token"
)

var cmd = struct {
  UsageLine string
  Short string
  Long string
  Flag *flag.FlagSet
}{
  UsageLine: "extract-imports [-h|--help] [-e path,path] source_path",
  Short:     "Extract imports directives of the given source directory",
  Long: `
To extract directives of each file of the directory 'src'
  extract-imports src/

To exclude some directories of the process
  extract-imports -e path_to_exclude src/

To exclude multiple directories, seperate them by a comma
  extract-imports -e path_to_exclude,path_to_exclude2 src/

Paths are always resolved against cwd.

The resulting output prints [file]: [import]
  `,
  Flag: flag.NewFlagSet("main", flag.ExitOnError),
}

func main() {

  var exclude = cmd.Flag.String("e", "", "List of paths to exclude")
  var help = cmd.Flag.Bool("h", false, "Show help")
  var h = cmd.Flag.Bool("help", false, "Show help")

  cmd.Flag.Usage = func () {
    fmt.Println(cmd.UsageLine)
    fmt.Println("")
    fmt.Println(cmd.Short)
    fmt.Println(cmd.Long)
  }

  cmd.Flag.Parse(os.Args[1:])

  var tail = cmd.Flag.Args()

  if *help || *h || len(tail)==0 {
    cmd.Flag.Usage()
    os.Exit(2)
  }

  if _, err := os.Stat(tail[0]); os.IsNotExist(err) {
    fmt.Fprintf(os.Stderr, "This path does not exists: %s\n", tail[0])
    os.Exit(2)
  }

  var excludes = strings.Split(*exclude, ",")
  if len(excludes[0])==0 && len(*exclude)>0 {
    excludes[0] = *exclude
  }

  var okExcludes = make([]string, 0, len(excludes))
  for _, element := range excludes {
    if len(element)>0 {
      var aSrc, err = filepath.Abs(element)
      if err != nil {
        fmt.Fprintf(os.Stderr, "%s", err)
        os.Exit(1)
      }
      okExcludes = append(okExcludes, aSrc)
    }
  }
  var aSrc, err = filepath.Abs(tail[0])
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
    os.Exit(1)
  }


  files, err := walkFiles(aSrc, okExcludes)
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
    os.Exit(1)
  }

  found := false
  for f := range files {
    i, err := parseImports(f.path)
    if err != nil {
      fmt.Fprintf(os.Stderr, "%s", err)
      os.Exit(1)
    }
    for _,importStr := range i.imports {
      found = true
      fmt.Printf("%s: %s\n", i.file, importStr)
    }
  }

  if found==false {
    fmt.Fprintf(os.Stderr, "No files found!\n")
    os.Exit(1)
  }

}


type FileItem struct {
  path string
  stats os.FileInfo
}

func walkFiles (path string, excludes []string) (chan FileItem, error) {
  out := make(chan FileItem)

  stats, err := os.Stat(path)
  if err != nil {
    close(out)
    return out, err
  }

  go func(){
    if stats.IsDir() {
      filepath.Walk(path, func (path string, stats os.FileInfo, err error) error {
        if stats.IsDir()==false {
          if isExcluded(path, excludes) == false {
          	out <- FileItem{path: path, stats: stats}
          }
        }
        return nil
      })
      close(out)
    } else {
      if isExcluded(path, excludes) == false {
      	out <- FileItem{path: path, stats: stats}
      }
      close(out)
    }
  }()

  return out, err
}

func isExcluded (path string, excludes []string) bool {
  ret := false
  for _, element := range excludes {
    if len(path)>=len(element) {
      if path[0:len(element)]==element {
        ret = true
        break
      }
    }
  }
  return ret
}

type ParsedImport struct {
  file string
  imports []string
}

func parseImports (path string) (ParsedImport, error) {

  ret := ParsedImport{file: path}

  fset := token.NewFileSet() // positions are relative to fset
  f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
  if err == nil {
    r, _ := regexp.Compile("^\".+\"$")
    for _, s := range f.Imports {
      v := s.Path.Value
      if r.MatchString(v) {
        v = v[1:len(v)-1]
      }
      ret.imports = append(ret.imports, v)
    }
  }

  return ret, err
}
