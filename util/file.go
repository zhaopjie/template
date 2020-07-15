/**
 * @Author: duke
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2020/6/18 10:45 上午
 */

package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// SaveFile 储存文件到本地
func SaveFile(fpath string, in io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}
	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer out.Close()     // nolint: errcheck
	err = out.Chmod(0644) //更改文件的权限
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", fpath, err)
	}
	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", fpath, err)
	}
	return nil
}
