package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"github.com/go-git/go-git/v5"
)

var directories = [5]string{"lua", "pack", "pack/colors/start", "pack/plugins/start", "pack/syntax/start"}

func main() {
	arg := strings.ToLower(os.Args[1])
	installDir := os.Args[2]

	switch arg {
		case "install": Install(installDir) // Will install all of the individual components
		case "grab": Grab() // Used for me to collect all of the configuration files and what nots, you probably don't need this
		case "gen-file-structure": GenFileStructure("./")
		default:
			fmt.Println("Uh oh, command not found! Did you mean (install <install-dir>, grab, gen-file-structure)")
	}
}

func Install(installDir string) {
	root := installDir

	err := os.Mkdir(root, os.ModePerm)
	if err != nil {
		fmt.Errorf("unable to generate directory: %w", err)
	}

	err = copyFile("./init.vim", root + "init.vim")
	if err != nil {
		fmt.Errorf("unable to copy file: %w", err)
	}

	for _, dir := range directories[:2] {
		fmt.Printf("Copying: %s -> %s\n", dir, root)
		err := copyDirectory(dir, root)
		if err != nil {
			fmt.Errorf("unable to copy directory: %w", err)
		}
	}
}

func Grab() {
	GenFileStructure("./")

	homeDir, err := os.UserHomeDir()
    if err != nil {
		fmt.Errorf("unable to find home directory: %w", err)
    }
	fileRoot := homeDir + "/.config/nvim/"
	
	files := [2]string{"init.vim", "lua/statusbar.lua"}

	for _, file := range files {
		fmt.Printf("Copying: %s -> %s\n", fileRoot + file, file)
		err := copyFile(fileRoot + file, "./" + file)
		if err != nil {
			fmt.Errorf("unable to copy file: %w", err)
		}
	}
}

func GenFileStructure(root string) {
	for _, dir := range directories {
		fmt.Printf("Generating: %s\n", dir)
		err := os.MkdirAll(root + dir, os.ModePerm)
		if err != nil {
			fmt.Errorf("unable to generate directory: %w", err)
		}
	}
}

func copyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}

	defer r.Close()
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()
	w.ReadFrom(r)

	return nil
}

func copyDirectory(src, dst string) error {
    cmd := exec.Command("cp", "-R", src, dst)
	err := cmd.Run()
	if err != nil {
		return err
	}
	
	return nil
}
