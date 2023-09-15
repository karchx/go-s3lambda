package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"s3lambda-api/aws"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
)

func uploadCSVToS3(filename string, fileContent []byte) error {
	svc, err := aws.CreateS3Session()
	if err != nil {
		return err
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.StringAws("gos3lambda-test"),
		Key:    aws.StringAws(filename),
		Body:   bytes.NewReader(fileContent),
	})

	return err
}

func listCSVFilesInS3() ([]string, error) {
	svc, err := aws.CreateS3Session()
	if err != nil {
		return nil, err
	}

	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.StringAws("gos3lambda-test"),
	})
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, item := range result.Contents {
		fileNames = append(fileNames, *item.Key)
	}
	return fileNames, nil
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, metadata, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error upload file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileExtension := filepath.Ext(metadata.Filename)
	if fileExtension != ".csv" {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	fileContent, err := io.ReadAll(file)

	err = uploadCSVToS3(metadata.Filename, fileContent)
	if err != nil {
		http.Error(w, "Error upload file", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Upload file")
}

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	fileNames, err := listCSVFilesInS3()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(fileNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadFileHandler).Methods("POST")
	r.HandleFunc("/list", listFilesHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server in port 8080")
	http.ListenAndServe(":8080", nil)
}
