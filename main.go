// by Gonzaru
// Distributed under the terms of the GNU General Public License v3

package main

import (
	"fmt"
	"log"
	"os"
)

// local packages
import (
	"github.com/gonzaru/sf/config"
	"github.com/gonzaru/sf/sf"
	"github.com/gonzaru/sf/utils"
)

// help shows help information
func help() {
	fmt.Print("Usage:\n")
	fmt.Printf("  %s                # opens an interactive menu\n", config.ProgName)
	fmt.Printf("  %s /path/to/file  # opens the local file\n", config.ProgName)
}

// main sf
func main() {
	if errSl := sf.SetLog(); errSl != nil {
		utils.ErrPrint(errSl)
		log.Fatal(errSl)
	}
	args := os.Args[1:]
	lenArgs := len(args)
	if lenArgs > 0 {
		if lenArgs == 1 {
			file := args[0]
			if errSp := sf.Spawn(file); errSp != nil {
				utils.ErrPrint(errSp)
				log.Fatal(errSp)
			}
		} else {
			help()
			os.Exit(1)
		}
	} else {
		go sf.SignalHandler()
		if errSf := sf.Run(); errSf != nil {
			utils.ErrPrint(errSf)
			log.Fatal(errSf)
		}
	}
}
