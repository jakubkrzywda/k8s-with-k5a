# Demo App - Go

## Basic

Build the Node.js container image:

```bash
sudo podman build -t nodejsdemoapp .
```

Scan the image for vulnerabilities:

```bash
trivy image nodejsdemoapp
```

Run the Go container:

```bash
sudo podman run --rm -p 8080:8080 nodejsdemoapp
```

Visit http://localhost:8080
