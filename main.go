package main

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
)

var sess = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("north virginia"),
}))

type Event struct {
	File     []byte `json:"file"`
	FileName string `json:"file_name"`
}

type Response struct {
	Message string   `json:"message"`
	Files   []string `json:"files"`
}

func UploadFileToS3Handler(w http.ResponseWriter, r *http.Request) {
	var request Event

	file := request.File
	fileName := request.FileName

	svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket: aws.String("name-bucket"),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(file),
	}

	_, err := svc.PutObject(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload file successfully"))
}

func ListFilesInS3Hadler(w http.ResponseWriter, r *http.Request) {
	svc := s3.New(sess)

	params := &s3.ListObjectsInput{
		Bucket: aws.String("name-bucket"),
	}

	resp, err := svc.ListObjects(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileList []string
	for _, item := range resp.Contents {
		fileList = append(fileList, *item.Key)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Files en S3: \n" + strings.Join(fileList, "\n")))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", UploadFileToS3Handler).Methods("POST")
	r.HandleFunc("/list", ListFilesInS3Hadler).Methods("GET")

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
