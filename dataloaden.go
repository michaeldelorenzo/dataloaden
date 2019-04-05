package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vektah/dataloaden/pkg/generator"
)

func main() {
	keyType := flag.String("keys", "int", "what type should the keys be")
	slice := flag.Bool("slice", false, "this dataloader will return slices")
	loaderDir := flag.String("dir", "", "directory to store the generated loaders")

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if loaderDir == nil || *loaderDir == "" {
		loaderDir = &wd
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if err := generator.Generate(flag.Arg(0), *keyType, *slice, *loaderDir, wd); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
