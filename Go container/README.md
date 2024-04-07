# Demo App - Go

## Basic

Build the Go container image:

```bash
sudo podman build -t golangdemoapp .
```

Scan the image for vulnerabilities:

```bash
trivy image golangdemoapp
```

Run the Go container:

```bash
sudo podman run --rm -p 8080:8080 golangdemoapp
```

Visit http://localhost:8080/hello

## Advanced

Build the Go container image:

```bash
sudo podman build -t golangdemoapp-advanced -f Dockerfile_advance .
```

Scan the image for vulnerabilities:

```bash
trivy image golangdemoapp-advanced
```

Run the Go container:

```bash
sudo podman run --rm -p 8080:8080 golangdemoapp-advanced
```

Visit http://localhost:8080/hello