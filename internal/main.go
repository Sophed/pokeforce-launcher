package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/adrg/xdg"
)

const CHECKSUM_URL = "https://client.pokeforce.org/repository/checksum_linux.md5"
const CLIENT_URL = "https://client.pokeforce.org/repository/Linux.zip"

const CHECKSUM_FILE = "checksum"
const CLIENT_ZIP = "client.zip"
const CLIENT_FILE = "PokeForce.x86_64"

func main() {
	dataPath := xdg.DataHome + "/pokeforce"
	_, err := os.Stat(dataPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("failed to read data directory:\n", err)
		}
		run(Init(dataPath))
		return
	}
	if FetchChecksum() == CurrentChecksum(dataPath+"/"+CHECKSUM_FILE) {
		println("no new updates")
		run(dataPath + "/" + CLIENT_FILE)
		return
	}
	run(update(dataPath))
}

func update(path string) string {
	StoreChecksum(path)
	DownloadClient(path)
	err := os.Remove(path + "/" + CLIENT_FILE)
	if err != nil {
		log.Fatal("failed to remove old client:\n", err)
	}
	client, err := ExtractClient(path)
	if err != nil {
		log.Fatal("failed to extract client:\n", err)
	}
	os.Chmod(client, os.ModePerm)
	return client
}

func run(client string) {
	println("launching pokeforce...")
	exec.Command(client).Run()
}
