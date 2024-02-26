package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gocloud.dev/blob"

	_ "gocloud.dev/blob/azureblob"
)

func main() {
	//	if len(os.Args) != 3 {
	// 		log.Fatal("usage: upload bucket_url file")
	//	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env: %s", err)
	}

	bucketURL := os.Args[1]

	// initialize context
	ctx := context.Background()

	// open connection
	bucket, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		log.Fatalf("failed setup bucket: %s", err)
	}
	defer bucket.Close()

	for _, filePath := range os.Args[2:] {
		// read
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("failed to read file: %s", err)
		}

		// write & upload
		blobWriter, err := bucket.NewWriter(ctx, filePath, nil)
		if err != nil {
			log.Fatalf("failed to obtain writer: %s", err)
		}
		_, err = blobWriter.Write(fileContent)
		if err != nil {
			log.Fatalf("failed to write to bucket: %s", err)
		}
		if err := blobWriter.Close(); err != nil {
			log.Fatalf("failed to close writer: %s", err)
		} else {
			log.Println("upload successful!")
		}
	}
}
