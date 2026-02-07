# http-server

A simple HTTP tool with two modes: **echo server** and **file server**.

## Installation

```bash
go install github.com/doraemonkeys/http-server@latest
```

## Usage

### Echo Server

Start an HTTP echo server that returns request details as JSON:

```bash
http-server echo [-port 6688] [-ip 0.0.0.0]
```

| Flag    | Default   | Description            |
|---------|-----------|------------------------|
| `-port` | `6688`    | Port to listen on      |
| `-ip`   | `0.0.0.0` | IP address to listen on |

### File Server

Start an HTTP file server that serves a directory:

```bash
http-server file [-port 8080] [-ip 0.0.0.0] [-d .]
```

| Flag    | Default   | Description            |
|---------|-----------|------------------------|
| `-port` | `8080`    | Port to listen on      |
| `-ip`   | `0.0.0.0` | IP address to listen on |
| `-d`    | `.`       | Directory to serve     |
