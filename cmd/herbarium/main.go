package main

import (
	"fmt"
	"log"

	"github.com/helpellie1502/herbarium/internal/config"
	"github.com/helpellie1502/herbarium/internal/project"
)

func main() {

	root := "/home/tsvet/code/herbarium"

	p, err := project.Scan(root, []string{".git"}, []string{"go.mod"}, []string{".md"})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(p.Tree()); i++ {
		entry := p.Tree()[i]
		fmt.Println(entry.Path(), entry.IsDir(), entry.IsIgnored())
	}

	fmt.Println(config.Path())

}
