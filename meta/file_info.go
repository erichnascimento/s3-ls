package meta

import (
	"time"

	"github.com/dustin/go-humanize"
)

// FileInfo is the information of File
type FileInfo struct {
	Name         string
	LastModified time.Time
	Owner        string
	Size         uint64
}

// HumanizeSize return a humanized file size info
func (*f) HumanSize() string {
	return humanize.Bytes(f.Size)
}
