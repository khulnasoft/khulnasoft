package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"mime"
	"os"
	"path"
	"strings"

	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/s3"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func filebase64OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		return base64.StdEncoding.EncodeToString(fileData[:])
	} else {
		panic(err.Error())
	}
}

func filebase64sha256OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		hashedData := sha256.Sum256([]byte(fileData))
		return base64.StdEncoding.EncodeToString(hashedData[:])
	} else {
		panic(err.Error())
	}
}

func sha1Hash(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		encoded := base64.StdEncoding.EncodeToString([]byte("haha business"))
		tmpVar0, _ := base64.StdEncoding.DecodeString(encoded)
		decoded := string(tmpVar0)
		_ = strings.Join([]string{
			encoded,
			decoded,
			"2",
		}, "-")
		// tests that we initialize "var, err" with ":=" first, then "=" subsequently (Go specific)
		_, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{}, nil)
		if err != nil {
			return err
		}
		_, err = aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{}, nil)
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		_ = bucket.ID().ApplyT(func(id string) (khulnasoft.String, error) {
			return khulnasoft.String(base64.StdEncoding.EncodeToString([]byte(id))), nil
		}).(khulnasoft.StringOutput)
		_ = bucket.ID().ApplyT(func(id string) (khulnasoft.String, error) {
			value, _ := base64.StdEncoding.DecodeString(id)
			return khulnasoft.String(value), nil
		}).(khulnasoft.StringOutput)
		secretValue := khulnasoft.ToSecret("hello").(khulnasoft.StringOutput)
		_ = khulnasoft.Unsecret(secretValue).(khulnasoft.StringOutput)
		currentStack := ctx.Stack()
		currentProject := ctx.Project()
		workingDirectory := func(cwd string, err error) string {
			if err != nil {
				panic(err)
			}
			return cwd
		}(os.Getwd())
		fileMimeType := mime.TypeByExtension(path.Ext("./base64.txt"))
		// using the filebase64 function
		_, err = s3.NewBucketObject(ctx, "first", &s3.BucketObjectArgs{
			Bucket:      bucket.ID(),
			Source:      khulnasoft.NewStringAsset(filebase64OrPanic("./base64.txt")),
			ContentType: khulnasoft.String(fileMimeType),
			Tags: khulnasoft.StringMap{
				"stack":   khulnasoft.String(currentStack),
				"project": khulnasoft.String(currentProject),
				"cwd":     khulnasoft.String(workingDirectory),
			},
		})
		if err != nil {
			return err
		}
		// using the filebase64sha256 function
		_, err = s3.NewBucketObject(ctx, "second", &s3.BucketObjectArgs{
			Bucket: bucket.ID(),
			Source: khulnasoft.NewStringAsset(filebase64sha256OrPanic("./base64.txt")),
		})
		if err != nil {
			return err
		}
		// using the sha1 function
		_, err = s3.NewBucketObject(ctx, "third", &s3.BucketObjectArgs{
			Bucket: bucket.ID(),
			Source: khulnasoft.NewStringAsset(sha1Hash("content")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
