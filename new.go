package main

import (
	"context"
	"flag"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func new() {
	name := flag.Arg(1)
	folder := flag.Arg(2)

	if name == "" {
		newHelp()
		return
	}

	if folder == "" {
		folder = "./" + name
	}

	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           "https://github.com/VOLIX-dev/leopard-base.git",
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewBranchReferenceName("master"),
		Progress:      os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.RemoveAll(path.Join(folder, ".git"))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				file, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				file = []byte(strings.Replace(string(file), "project_name", name, -1))

				err = os.WriteFile(path, file, 0644)
				return err
			}

			return nil
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	err = exec.CommandContext(context.TODO(), "go", "mod", "verify").Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func newHelp() {
	fmt.Println(`Usage: leopard new <name> [folder]`)
}
