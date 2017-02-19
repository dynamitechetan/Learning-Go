package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	all := flag.Bool("a", false, "Include all files")
	ftype := flag.Bool("F", false, "Show file type suffix")
	//other potential options:
	//-l follow symlinks, default is to not
	//-P include matches to pattern
	//-I exclude matches to pattern
	flag.Parse()
	args := flag.Args()
	var fileFilter func(os.FileInfo) bool
	if !*all {
		fileFilter = defaultFilter
	}
	if len(args) > 0 {
		for _, dir := range args {
			tree(dir, fileFilter, *ftype)
		}
	} else {
		tree(".", fileFilter, *ftype)
	}
}

func tree(dir string, filter func(os.FileInfo) bool, ftype bool) {
	subtree(dir, filter, ftype, "")
}

func defaultFilter(info os.FileInfo) bool {
	name := info.Name()
	if strings.HasPrefix(name, ".") || strings.HasSuffix(name, "~") || strings.HasPrefix(name, "#") {
		return true
	}
	return false
}

type Info struct {
	Name  string
	IsDir bool
	Mode  os.FileMode
}

func readDir(dir string, filter func(os.FileInfo) bool) ([]Info, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	result := make([]Info, 0, len(infos))
	for _, info := range infos {
		if filter == nil || !filter(info) {
			result = append(result, Info{Name: info.Name(), IsDir: info.IsDir(), Mode: info.Mode()})
		}
	}
	return result, nil
}

func subtree(dir string, filter func(os.FileInfo) bool, ftype bool, indent string) {
	infos, err := readDir(dir, filter)
	if err == nil {
		last := len(infos) - 1
		spacer := "├──"
		max := len(infos)
		for i := 0; i < max; i++ {
			info := infos[i]
			name := info.Name
			if i == last {
				spacer = "└──"
			}
			if ftype {
				fmt.Printf("%s%s %s%s\n", indent, spacer, name, fileTypeSuffix(info))
			} else {
				fmt.Printf("%s%s %s\n", indent, spacer, name)
			}
			if info.IsDir {
				var nextIndent string
				if i+1 < max {
					nextIndent = indent + "│   "
				} else {
					nextIndent = indent + "    "
				}
				subtree(filepath.Join(dir, name), filter, ftype, nextIndent)
			}
		}
	}
}

func fileTypeSuffix(info Info) string {
	switch info.Mode & os.ModeType {
	case os.ModeDir:
		return "/"
	case os.ModeSymlink:
		return "@"
	case os.ModeSocket:
		return "="
	case os.ModeNamedPipe:
		return "|"
	}
	perms := info.Mode & os.ModePerm
	if (perms & 0111) != 0 {
		return "*"
	}
	return ""
}