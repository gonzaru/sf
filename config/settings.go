// by Gonzaru
// Distributed under the terms of the GNU General Public License v3

package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const ProgName = "sf"

var (
	SFLog    = fmt.Sprintf("%s/%s-sf.log", tmpDir, userName)
	tmpDir   = os.TempDir()
	userName = getUserName()
	Term     = "xterm"
	TermArgs = []string{"-e"}
)

// ProgExt returns the program associated by their extension
func ProgExt(file string) (map[string]interface{}, error) {
	var prg string
	var prgArgs []string
	var useTerm bool
	prgOpts := make(map[string]interface{})
	ext := filepath.Ext(file)
	switch ext {
	case ".pdf":
		prg = "mupdf"
		prgArgs = append(prgArgs, "-A", "8")
	case ".aif", ".avi", ".cda", ".mid", ".midi", ".mkv", ".mov", ".mp3", ".mp4", ".mpg", ".mpeg", ".ogg", ".wav",
		".wma", ".wmv":
		prg = "gorum"
		prgArgs = []string{}
	case ".doc", ".docx", ".odt", ".ppt", ".pptx", ".rtf", ".xls", ".xlsx":
		prg = "soffice"
	case ".txt", ".c", ".conf", ".cpp", ".css", ".go", ".h", ".htm", ".html", ".ini", ".js", ".json", ".log", ".md",
		".php", ".pl", ".py", ".rb", ".sh", ".sql", ".tmp", ".yaml", ".yml", ".vim", ".xhtml", ".xml":
		// prg = "gvim"
		// prgArgs = append(prgArgs, "--servername", ProgName, "--remote-silent")
		prg = "vim"
		prgArgs = []string{}
		useTerm = true
	case ".bmp", ".gif", ".ico", ".jpg", ".jpeg", ".png", ".svg", ".tif", ".tiff":
		prg = "geeqie"
		prgArgs = []string{}
	default:
		fi, err := os.Lstat(file)
		if os.IsNotExist(err) {
			return prgOpts, fmt.Errorf("progExt: error: '%s' no such file or directory\n", file)
		} else if err != nil {
			return prgOpts, err
		}
		if !fi.IsDir() && fi.Mode().IsRegular() {
			prg = "vim"
			prgArgs = []string{}
			useTerm = true
		}
	}
	prgOpts["name"] = prg
	prgOpts["args"] = prgArgs
	prgOpts["useTerm"] = useTerm
	return prgOpts, nil
}

// getUserName
func getUserName() string {
	usc, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usc.Username
}
