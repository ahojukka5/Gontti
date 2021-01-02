package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Gontti container util")
	tag := os.Args[1]
	fileURL := "https://github.com/" + tag + "/archive/master.zip"
	fmt.Println("Fetching tag" + tag + " from url " + fileURL)

	err := DownloadFile("master.zip", fileURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded file: " + fileURL)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
