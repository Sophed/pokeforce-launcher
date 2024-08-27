package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Init(path string) string {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatal("failed to create directory:\n", err)
	}
	println("created data directory: " + path)
	StoreChecksum(path)
	DownloadClient(path)
	client, err := ExtractClient(path)
	if err != nil {
		log.Fatal("failed to extract client:\n", err)
	}
	os.Chmod(client, os.ModePerm)
	return client
}

func FetchChecksum() string {
	res, err := http.Get(CHECKSUM_URL)
	if err != nil {
		log.Fatal("error fetching checksum:\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("unknown error occured")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func StoreChecksum(path string) {
	data := []byte(FetchChecksum())
	err := os.WriteFile(path+"/"+CHECKSUM_FILE, data, 0644)
	if err != nil {
		log.Fatal("failed to write checksum file:\n", err)
	}
	println("saved checksum")
}

func CurrentChecksum(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("failed to read checksum:\n", err)
	}
	return string(dat)
}
