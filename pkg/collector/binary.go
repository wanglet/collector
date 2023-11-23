package collector

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type Binary struct {
	Url       string `yaml:"url"`
	Sha256sum string `yaml:"sha256sum"`
	Arch      string `yaml:"arch"`
	Filename  string `yaml:"filename"`
	directory string
	filePath  string
}

func (b *Binary) Validate() error {
	if b.Url == "" {
		return fmt.Errorf("binary url is required")
	}
	if b.Sha256sum == "" {
		return fmt.Errorf("binary sha256sum is required")
	}
	if b.Arch == "" {
		return fmt.Errorf("binary arch is required")
	}
	return nil
}

func (b *Binary) String() string {
	return fmt.Sprintf("Binary{Url: %s, Sha256sum: %s, Arch: %s, Filename: %s}", b.Url, b.Sha256sum, b.Arch, b.Filename)
}

func (b *Binary) prepare(outputPath string) error {
	// 创建目录和赋值directory、filePath变量
	b.directory = path.Join(outputPath, "binaries")
	if _, err := os.Stat(b.directory); os.IsNotExist(err) {
		err := os.MkdirAll(b.directory, 0755)
		if err != nil {
			return err
		}
	}
	if b.Filename == "" {
		b.Filename = path.Base(b.Url)
	}
	b.filePath = path.Join(b.directory, b.Filename)

	return nil
}

func (b *Binary) download() error {
	fmt.Println("Downloading", b.Url, "to", b.filePath)
	resp, err := http.Get(b.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile(b.filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (b *Binary) verify() error {
	f, err := os.Open(b.filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	s := sha256.New()

	if _, err := io.Copy(s, f); err != nil {
		return err
	}

	sha256sum := fmt.Sprintf("%x", s.Sum(nil))

	if sha256sum != b.Sha256sum {
		return fmt.Errorf("sha256sum mismatch: %s != %s", sha256sum, b.Sha256sum)
	}

	return nil
}

func (b *Binary) Collect(ctx context.Context, outputPath string) error {
	if err := b.prepare(outputPath); err != nil {
		return err
	}

	// 文件是否已经下载，并且checksum正确. 否则删除重新下载
	if _, err := os.Stat(b.filePath); err == nil {
		if err = b.verify(); err == nil {
			// 文件存在并且checksum正确
			fmt.Println("File", b.filePath, "already exists and is valid")
			return nil
		} else {
			// 删除文件重新下载
			fmt.Println(err)
			if err = os.Remove(b.filePath); err != nil {
				return err
			}
		}
	}

	if err := b.download(); err != nil {
		return err
	}

	if err := b.verify(); err != nil {
		return err
	}
	return nil
}
