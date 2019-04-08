package backend

import (
	"os"

	"github.com/peterbourgon/diskv"

	"github.com/AkihikoITOH/lesen/util"
)

type DiskvBackend struct {
	base *diskv.Diskv
}

func (be *DiskvBackend) Write(key string) error {
	return be.base.Write(hash(key), []byte{})
}

func (be *DiskvBackend) Read(key string) bool {
	_, err := be.base.Read(hash(key))
	return err == nil
}

func (be *DiskvBackend) Erase(key string) error {
	return be.base.Erase(key)
}

func NewDiskvBackend(baseDir string) *DiskvBackend {
	_ = os.Mkdir(baseDir, os.ModePerm)
	return &DiskvBackend{
		diskv.New(diskv.Options{
			BasePath:     baseDir,
			Transform:    func(_ string) []string { return []string{} },
			CacheSizeMax: 1024 * 1024,
		}),
	}
}

func hash(str string) string {
	return util.MD5Sum(str)
}
