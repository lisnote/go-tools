package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var (
		endpoint      = flag.String("endpoint", "", "S3 endpoint URL (e.g., http://localhost:9000)")
		bucket        = flag.String("bucket", "", "S3 bucket name")
		pathFlag      = flag.String("path", "", "Destination directory/prefix in bucket (optional)")
		region        = flag.String("region", "us-east-1", "AWS region")
		accessKey     = flag.String("access-key", "", "AWS Access Key ID")
		secretKey     = flag.String("secret-key", "", "AWS Secret Access Key")
		pathStyle     = flag.Bool("path-style", true, "Use path-style addressing (http://endpoint/bucket/key)")
		replaceBefore = flag.String("replace-before", "", "String to replace in output URL")
		replaceAfter  = flag.String("replace-after", "", "Replacement string in output URL")
	)
	flag.Parse()

	files := flag.Args()

	if *endpoint == "" || *bucket == "" || len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: typora-s3-image-uploader -endpoint <url> -bucket <name> [options] <file> [file...]")
		fmt.Fprintln(os.Stderr, "\nRequired flags:")
		fmt.Fprintln(os.Stderr, "  -endpoint    S3 endpoint URL")
		fmt.Fprintln(os.Stderr, "  -bucket      S3 bucket name")
		fmt.Fprintln(os.Stderr, "\nArguments:")
		fmt.Fprintln(os.Stderr, "  file         Local file(s) to upload")
		fmt.Fprintln(os.Stderr, "\nOptional flags:")
		fmt.Fprintln(os.Stderr, "  -path        Destination directory in bucket")
		fmt.Fprintln(os.Stderr, "  -region      AWS region (default: us-east-1)")
		fmt.Fprintln(os.Stderr, "  -access-key  AWS Access Key ID")
		fmt.Fprintln(os.Stderr, "  -secret-key  AWS Secret Access Key")
		fmt.Fprintln(os.Stderr, "  -path-style  Use path-style addressing (default: true)")
		os.Exit(1)
	}

	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error: file not accessible: %s: %v\n", file, err)
			os.Exit(1)
		}
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(*region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(*accessKey, *secretKey, ""),
		),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to load config: %v\n", err)
		os.Exit(1)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(*endpoint)
		o.UsePathStyle = *pathStyle
	})

	pathPrefix := strings.TrimRight(*pathFlag, "/")

	for _, file := range files {
		key := filepath.Base(file)
		if pathPrefix != "" {
			key = pathPrefix + "/" + key
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to open %s: %v\n", file, err)
			os.Exit(1)
		}

		_, err = client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(*bucket),
			Key:    aws.String(key),
			Body:   f,
		})
		f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: upload failed for %s: %v\n", file, err)
			os.Exit(1)
		}

		url := buildURL(*endpoint, *bucket, key, *pathStyle)
		if *replaceBefore != "" {
			url = strings.ReplaceAll(url, *replaceBefore, *replaceAfter)
		}
		fmt.Println(url)
	}
}

func buildURL(endpoint, bucket, key string, pathStyle bool) string {
	endpoint = strings.TrimRight(endpoint, "/")
	key = strings.TrimLeft(key, "/")

	if pathStyle {
		return fmt.Sprintf("%s/%s/%s", endpoint, bucket, key)
	}

	host := endpoint
	scheme := ""
	if strings.HasPrefix(endpoint, "http://") {
		scheme = "http://"
		host = strings.TrimPrefix(endpoint, "http://")
	} else if strings.HasPrefix(endpoint, "https://") {
		scheme = "https://"
		host = strings.TrimPrefix(endpoint, "https://")
	}
	return fmt.Sprintf("%s%s.%s/%s", scheme, bucket, host, key)
}
