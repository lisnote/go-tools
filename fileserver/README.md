# fileserver

A lightweight Go file server with Basic Auth password protection for quick local directory sharing.

## Features

- Serve static files over HTTP
- Basic Auth username/password authentication
- Customizable port, directory, and credentials via CLI flags
- Single binary, no dependencies

## Usage

```bash
fileserver [options]
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `-u` | `admin` | Username |
| `-p` | `123456` | Password |
| `-port` | `8080` | Listening port |
| `-dir` | `./` | Directory to share |

### Examples

Start with defaults:

```bash
./fileserver
```

Specify directory, port, and credentials:

```bash
./fileserver -dir "./" -port 8080 -u admin -p 123456
```

Then open `http://localhost:8080` in your browser and enter the credentials.

## Notes

- The default password is weak; change it before exposing to the public internet
- Traffic is unencrypted; use within a trusted network for sensitive files
