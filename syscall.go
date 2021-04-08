// +build !darwin,!freebsd,!netbsd

package main

import "os"
import "syscall"

func atime(info os.FileInfo) syscall.Timespec {
	return sysStat(info).Atim
}
