# Demo App - .NET

## Basic

Build the .NET container image:

```bash
sudo podman build -t dotnetdemoapp .
```

Scan the image for vulnerabilities:

```bash
trivy image dotnetdemoapp
```

Run the .NET container:

```bash
sudo podman run --rm -p 5001:5000 dotnetdemoapp
```

Visit http://localhost:5001/Environment

## Advanced

Build the .NET container image:

```bash
sudo podman build -t dotnetdemoapp-advanced -f Dockerfile_advance .
```

Scan the image for vulnerabilities:

```bash
trivy image dotnetdemoapp-advanced
```

Run the .NET container:

```bash
sudo podman run --rm -p 8080:8080 dotnetdemoapp-advanced
```

Visit http://localhost:8080/Environment