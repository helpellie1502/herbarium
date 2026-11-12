package main

import (
	"fmt"
	"log"

	"github.com/helpellie1502/herbarium/internal/project"
)

func main() {

	root := "/home/tsvet/code/herbarium"

	p, err := project.Scan(root, []string{".git"}, []string{"go.mod"})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(p.Tree()); i++ {
		res, _ := p.TEntry(i, "DIR", "FILE", false)
		fmt.Println(res)
	}

}
