package minio

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioCfg struct {
	Host            string
	Port            string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
}

func NewConfig(host, port, accessKeyID, secretAccessKey, bucketName string) *MinioCfg {
	return &MinioCfg{
		Host:            host,
		Port:            port,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		BucketName:      bucketName,
	}
}

func New(ctx context.Context, cfg *MinioCfg) *minio.Client {
	endpoint := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	accessKeyID := cfg.AccessKeyID
	secretAccessKey := cfg.SecretAccessKey

	log.Println("minio client init")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("failed to connect to minio; err: %v", err)
	}

	log.Printf("%#v\n", minioClient)

	bucketName := cfg.BucketName

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("already own %s\n", bucketName)
		} else {
			log.Fatalf("failed to create bucket %s; err: %v", bucketName, err)
		}
	} else {
		log.Printf("successfully created %s\n", bucketName)
	}

	return minioClient
}
