package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/schollz/progressbar/v3"
)

func DownloadClient(path string) {
	out, err := os.Create(path + "/" + CLIENT_ZIP)
	if err != nil {
		log.Fatal("failed to create client:\n", err)
	}
	defer out.Close()

	resp, err := http.Get(CLIENT_URL)
	if err != nil {
		log.Fatal("failed to create client:\n", err)
	}
	defer resp.Body.Close()

	downloadBar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	_, err = io.Copy(io.MultiWriter(out, downloadBar), resp.Body)
	if err != nil {
		log.Fatal("failed to copy client data:\n", err)
	}
}

func ExtractClient(dest string) (string, error) {
	source := (dest + "/" + CLIENT_ZIP)
	println("extracting client...")
	read, err := zip.OpenReader(source)
	if err != nil {
		return "", err
	}
	for _, file := range read.File {
		if file.Mode().IsDir() {
			continue
		}
		open, err := file.Open()
		if err != nil {
			return "", err
		}
		name := path.Join(dest, file.Name)
		os.MkdirAll(path.Dir(name), os.ModeDir)
		create, err := os.Create(name)
		if err != nil {
			return "", err
		}
		defer create.Close()
		create.ReadFrom(open)
	}
	read.Close()
	err = os.Remove(dest + "/" + CLIENT_ZIP)
	if err != nil {
		log.Fatal("failed to remove archive:\n", err)
	}
	return dest + "/" + CLIENT_FILE, nil
}
