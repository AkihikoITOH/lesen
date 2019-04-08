package util

import (
	"crypto/md5"
	"fmt"
	"os/user"
)

const (
	LesenDirFragment = "/.lesen"
	LogPathFragment  = "/lesen.log"
	DiskvDirFragment = "/diskv"
)

func LesenDir() string {
	return homeDir() + LesenDirFragment
}

func LogPath() string {
	return LesenDir() + LogPathFragment
}

func DiskvDir() string {
	return LesenDir() + DiskvDirFragment
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func MD5Sum(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum([]byte{}))
}
