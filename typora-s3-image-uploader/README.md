# Typora S3 Image Uploader

A lightweight CLI tool for uploading files to S3-compatible storage services.

## Usage

```bash
typora-s3-image-uploader -endpoint <url> -bucket <name> [options] <file> [file...]
```

### Required Flags

| Flag | Description |
|------|-------------|
| `-endpoint` | S3 endpoint URL, e.g. `http://localhost:9000` |
| `-bucket` | S3 bucket name |

### Arguments

| Argument | Description |
|----------|-------------|
| `file` | Local file(s) to upload |

### Optional Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-path` | - | Destination directory/prefix in bucket |
| `-region` | `us-east-1` | AWS region |
| `-access-key` | - | AWS Access Key ID |
| `-secret-key` | - | AWS Secret Access Key |
| `-path-style` | `true` | Use path-style addressing (`http://endpoint/bucket/key`) |
| `-replace-before` | - | String to replace in output URL |
| `-replace-after` | - | Replacement string in output URL |

## Examples

**MinIO / Local S3:**
```bash
typora-s3-image-uploader \
  -endpoint http://localhost:9000 \
  -bucket mybucket \
  -path uploads \
  -access-key minioadmin \
  -secret-key minioadmin \
  ./data.csv
```

**AWS S3:**
```bash
typora-s3-image-uploader \
  -endpoint https://s3.ap-northeast-1.amazonaws.com \
  -bucket mybucket \
  -path backups \
  -region ap-northeast-1 \
  -access-key AKIA... \
  -secret-key secret... \
  -path-style=false \
  ./archive.zip
```

**Cloudflare R2 with custom domain:**
```bash
typora-s3-image-uploader \
  -endpoint https://<account-id>.r2.cloudflarestorage.com \
  -bucket mybucket \
  -access-key <access-key> \
  -secret-key <secret-key> \
  -replace-before "<account-id>.r2.cloudflarestorage.com/mybucket" \
  -replace-after "oss.example.com" \
  ./image.png
```

## Build

```bash
go build -o typora-s3-image-uploader .
```
