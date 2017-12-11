package storage

import (
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/taku-k/polymerase/pkg/base"
	"github.com/taku-k/polymerase/pkg/polypb"
)

type PhysicalStorage interface {
	Create(name string) (io.WriteCloser, error)
	CreateBackup(key polypb.Key, name string) (io.WriteCloser, error)
	Move(src string, dest polypb.Key) error
	Delete(name string) error
	DeleteBackup(key polypb.Key) error
	FullBackupStream(key polypb.Key) (io.Reader, error)
	IncBackupStream(key polypb.Key) (io.Reader, error)
	LoadXtrabackupCP(key polypb.Key) base.XtrabackupCheckpoints
	Walk(f func(path string, info os.FileInfo, err error) error) error
}

type DiskStorage struct {
	backupsDir string
}

func (s *DiskStorage) Create(name string) (io.WriteCloser, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *DiskStorage) CreateBackup(key polypb.Key, name string) (io.WriteCloser, error) {
	return s.Create(filepath.Join(s.backupsDir, string(key), name))
}

func (s *DiskStorage) Move(src string, dest polypb.Key) error {
	p := path.Join(s.backupsDir, string(dest))
	if err := os.MkdirAll(p, 0755); err != nil {
		return err
	}
	if err := os.Remove(p); err != nil {
		return err
	}
	return os.Rename(src, p)
}

func (s *DiskStorage) Delete(name string) error {
	return os.RemoveAll(name)
}

func (s *DiskStorage) DeleteBackup(key polypb.Key) error {
	return os.RemoveAll(filepath.Join(s.backupsDir, string(key)))
}

func (s *DiskStorage) FullBackupStream(key polypb.Key) (io.Reader, error) {
	return os.Open(filepath.Join(s.backupsDir, string(key), "base.tar.gz"))
}

func (s *DiskStorage) IncBackupStream(key polypb.Key) (io.Reader, error) {
	return os.Open(filepath.Join(s.backupsDir, string(key), "inc.xb.gz"))
}

func (s *DiskStorage) LoadXtrabackupCP(key polypb.Key) base.XtrabackupCheckpoints {
	return base.LoadXtrabackupCP(filepath.Join(s.backupsDir, string(key), "xtrabackup_checkpoints"))
}

func (s *DiskStorage) Walk(f func(path string, info os.FileInfo, err error) error) error {
	return filepath.Walk(s.backupsDir, f)
}

// TODO: implement it
type MemStorage struct {
}
