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

springio/gs-spring-boot-docker-advanced (alpine 3.19.1)

Total: 2 (UNKNOWN: 0, LOW: 2, MEDIUM: 0, HIGH: 0, CRITICAL: 0)

Java (jar)

Total: 16 (UNKNOWN: 0, LOW: 0, MEDIUM: 7, HIGH: 8, CRITICAL: 1)



Run the Java container:

```bash
sudo podman run --rm -p 8080:8080 springio/gs-spring-boot-docker-advanced
```

Visit http://localhost:8080

## Buildpacks

Build the Java container image using buildpacks:

```bash
pack build springio/gs-spring-boot-buildpack -B paketobuildpacks/builder-jammy-tiny
```

Scan the image for vulnerabilities:

```bash
trivy image springio/gs-spring-boot-buildpack
```

Run the Java container:

```bash
sudo podman run --rm -p 8080:8080 springio/gs-spring-boot-buildpack
```

Visit http://localhost:8080
