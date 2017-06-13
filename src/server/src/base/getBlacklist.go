package base

import (
	"archive/zip"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"

	"github.com/linkit360/xmp-api/src/structs"
)

func GetBlacklist(id_provider int) (string, error) {
	data := make([]xmp_api_structs.Blacklist, 0)
	prep := make([]string, 0)

	// Get campaigns by services ids
	db.Select("msisdn").Where("id_provider = ?", id_provider).Find(&data)
	for _, blacklist := range data {
		prep = append(prep, strconv.Itoa(blacklist.Msisdn))
	}
	text := strings.Join(prep, "\n")

	// ZIP
	filename := strconv.Itoa(id_provider) + "_" + time.Now().Format("20060102150405")
	defer os.Remove(filename)

	zipfile, err := os.Create(filename)
	if err != nil {
		log.Error(err)
	}
	defer zipfile.Close()

	w := zip.NewWriter(zipfile)
	f, err := w.Create("blacklist")
	if err != nil {
		log.Error(err)
	}

	_, err = f.Write([]byte(text))
	if err != nil {
		log.Error(err)
	}

	err = w.Close()
	if err != nil {
		log.Error(err)
	}

	// AWS S3
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			cfgAws.Id,
			cfgAws.Secret,
			"",
		),
	})
	if err != nil {
		log.Error("Cannot init S3 Session ", err)
	}

	svc := s3.New(sess)
	ctx := context.Background()

	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String("xmp-blacklist"),
		Key:    aws.String(filename),
		Body:   zipfile,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
	}

	return filename, nil
}
