package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffor bytes.Buffer
	writer := multipart.NewWriter(&buffor)
	writer.WriteField("name", "Michael jackson")

	part := make(textproto.MIMEHeader)
	part.Set("content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpeg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.jpeg")
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
