package fs

import (
	"context"
	"encoding/json"
	"os"

	"github.com/shurcooL/webdavfs/vfsutil"
	"golang.org/x/net/webdav"
)

// jsonEncodeFile encodes v into file at path, overwriting or creating it.
func jsonEncodeFile(ctx context.Context, fs webdav.FileSystem, path string, v interface{}) error {
	f, err := fs.OpenFile(ctx, path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(v)
}

// jsonDecodeFile decodes contents of file at path into v.
func jsonDecodeFile(ctx context.Context, fs webdav.FileSystem, path string, v interface{}) error {
	f, err := vfsutil.Open(ctx, fs, path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}
