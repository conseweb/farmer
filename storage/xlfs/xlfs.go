/// Xinlian FS
package xlfs

import (
	"io"

	"github.com/hyperledger/fabric/storage"
	"golang.org/x/net/context"
)

type Driver struct {
}

var _ storage.StorageDriver = &Driver{}

func NewDriver() *Driver {
	d := &Driver{}

	return d
}

func (d *Driver) Name() string {
	return "xinlian filesystem"
}

func (d *Driver) Init(ctx context.Context) error {
	return nil
}

func (d *Driver) GetContent(ctx context.Context, path string) ([]byte, error) {
	return nil, nil
}

func (d *Driver) PutContent(ctx context.Context, path string, content []byte) error {
	return nil
}

func (d *Driver) Reader(ctx context.Context, path string) (io.ReadCloser, error) {
	return nil, nil
}

func (d *Driver) Writer(ctx context.Context, path string, isAppend bool) (io.WriteCloser, error) {
	return nil, nil
}

func (d *Driver) Stat(ctx context.Context, path string) (storage.FileInfo, error) {
	return nil, nil
}

func (d *Driver) List(ctx context.Context, path string) ([]storage.FileInfo, error) {
	return nil, nil
}

func (d *Driver) Mkdir(ctx context.Context, path string) error {
	return nil
}

func (d *Driver) Move(ctx context.Context, sourcePath string, destPath string) error {
	return nil
}

func (d *Driver) Delete(ctx context.Context, path string) error {
	return nil
}
