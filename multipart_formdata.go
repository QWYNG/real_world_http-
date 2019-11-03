package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffor bytes.Buffer
	writer := multipart.NewWriter(&buffor)
	writer.WriteField("name", "Michael jackson")

	fileWriter, err := writer.CreateFormFile("text", "test.txt")
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffor)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
