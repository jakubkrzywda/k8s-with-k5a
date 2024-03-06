# Demo App - Java Spring Boot

## Basic

Build the Java app using Maven:

```bash
./mvnw package
```

Build the Java container image:

```bash
podman build -t springio/gs-spring-boot-docker .
```

Scan the image for vulnerabilities:

```bash
trivy image springio/gs-spring-boot-docker
```

Run the Java container:

```bash
podman run -p 8080:8080 springio/gs-spring-boot-docker
```

Visit http://localhost:8080

## Advanced

Build the Java container image:

```bash
sudo podman build -t springio/gs-spring-boot-docker-advanced -f Dockerfile_advance .
```

Scan the image for vulnerabilities:

```bash
trivy image springio/gs-spring-boot-docker-advanced
```

Run the Java container:

```bash
sudo podman run --rm -p 8080:8080 springio/gs-spring-boot-docker-advanced
```

Visit http://localhost:8080