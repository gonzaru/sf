// by Gonzaru
// Distributed under the terms of the GNU General Public License v3

package main

import "log"

// local packages
import (
	"github.com/gonzaru/sf/sf"
	"github.com/gonzaru/sf/utils"
)

// main sf
func main() {
	if errSl := sf.SetLog(); errSl != nil {
		utils.ErrPrint(errSl)
		log.Fatal(errSl)
	}
	go sf.SignalHandler()
	if errSf := sf.Run(); errSf != nil {
		utils.ErrPrint(errSf)
		log.Fatal(errSf)
	}
}
