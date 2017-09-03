echo "Building project and creating binary"
docker run --rm -v "$PWD":/go/src/github.com/Codigami/newsletter-engine -w /go/src/github.com/Codigami/newsletter-engine -e 'CGO_ENABLED=0' -e 'GOOS=linux' golang:1.8 /bin/bash -c "go build -a -v -ldflags '-w'"

echo "Creating container with the binary"
docker build -t newsletter-engine .
