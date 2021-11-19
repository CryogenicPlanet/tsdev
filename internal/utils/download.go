package utils

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFile(link string, path string) error {

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	return nil
}

func DownloadArchive(link string, dirName string) error {

	tmpFile := dirName + "/temp.zip"
	out, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	unzip(tmpFile, dirName)

	err = os.Remove(tmpFile)
	if err != nil {
		return err
	}
	return nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
