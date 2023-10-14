package mykube

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	_ "k8s.io/kubectl/pkg/cmd/cp"
	_ "unsafe"
)

// linkname exposes private packages

//go:linkname cpMakeTar k8s.io/kubectl/pkg/cmd/cp.makeTar
func cpMakeTar(srcPath, destPath string, writer io.Writer) error

//go:linkname cpStripPathShortcuts k8s.io/kubectl/pkg/cmd/cp.stripPathShortcuts
func cpStripPathShortcuts(p string) string

// https://github.com/ica10888/client-go-helper/blob/master/pkg/kubectl/cp.go
func untarAll(reader io.Reader, destDir string, prefix string) error {
	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		if !strings.HasPrefix(header.Name, prefix) {
			return fmt.Errorf("tar contents corrupted")
		}

		// basic file information
		mode := header.FileInfo().Mode()
		// TODO remove trailing slash
		// header.Name is a name of the REMOTE file
		destFileName := filepath.Join(destDir, header.Name[len(prefix):])

		baseName := filepath.Dir(destFileName)
		if err := os.MkdirAll(baseName, 0755); err != nil {
			return err
		}
		if header.FileInfo().IsDir() {
			if err := os.MkdirAll(destFileName, 0755); err != nil {
				return err
			}
			continue
		}

		if mode&os.ModeSymlink != 0 {
			fmt.Fprintf(os.Stderr, "warning: skipping symlink: %q -> %q\n", destFileName, header.Linkname)
			continue
		}

		outFile, err := os.Create(destFileName)
		if err != nil {
			return err
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, tarReader); err != nil {
			return err
		}
	}
	return nil
}
