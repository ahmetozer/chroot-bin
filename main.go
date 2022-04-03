package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	// Get execute name
	binName := filepath.Base(os.Args[0])

	// Default dir
	ChrootDIR := "/apps/" + binName
	if path := os.Getenv("CHROOTBIN_DIR_DEFAULT"); path != "" {
		ChrootDIR = path + binName
	}

	//Check custom dir is avaible
	if DIR := os.Getenv("CHROOTBIN_DIR"); DIR != "" {
		paths := strings.Split(DIR, ";")
		for i := range paths {
			kv := strings.Split(paths[i], "=")
			if kv[0] == binName {
				ChrootDIR = kv[1]
			}
		}
	}

	binPath := ""
	if path := os.Getenv("CHROOTBIN_PATH"); path != "" {
		paths := strings.Split(path, ";")
		for i := range paths {
			kv := strings.Split(paths[i], "=")
			if kv[0] == binName {
				binPath = kv[1]
			}
		}
	}

	if err := syscall.Chroot(ChrootDIR); err != nil {
		log.Fatalf("sys chroot: %s\n", err)
	}

	if err := os.Chdir("/"); err != nil {
		log.Fatalf("chdir: %s\n", err)
	}

	if binPath == "" {
		var err error
		binPath, err = exec.LookPath(binName)

		if err != nil {
			log.Fatalf("bin path: %s\n", err)
		}
	}

	//switch to new process
	env := append(os.Environ(), "CHROOTBIN_EXEC=TRUE")
	if err := syscall.Exec(binPath, os.Args, env); err != nil {
		log.Fatalf("exec err: %s\n", err)
	}
	log.Fatal("this is not expected")
}
