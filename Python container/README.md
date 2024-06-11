# Demo App - Python

## Basic

Build the Python container image:

```bash
podman build -t python-docker-basic -f Dockerfile_basic .
```

Scan the image for vulnerabilities:

```bash
trivy image python-docker-basic
```

Run the Python container:

```bash
podman run --rm -p 8080:8080 python-docker-basic
```

Visit http://localhost:8080/quote

## Advanced

### Multi-stage

Build the Python container image:

```bash
podman build -t python-docker-advanced -f Dockerfile_advance .
```

Scan the image for vulnerabilities:

```bash
trivy image python-docker-advanced
```

Run the Python container:

```bash
podman run --rm -p 8080:8080 python-docker-advanced
```

Visit http://localhost:8080/quote

### Multi-stage from distroless

Build the Python container image:

```bash
podman build -t python-docker-distroless -f Dockerfile_distroless .
```

```bash
trivy image python-docker-distroless
```

Run the Python container:

```bash
sudo podman run --rm -p 8080:8080 python-docker-distroless
```

Visit http://localhost:8080/quote

## Compare the size of container images

```bash
podman images | grep python
```
