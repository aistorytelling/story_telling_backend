
#!/bin/bash
RUN_NAME=
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o output/bin/${RUN_NAME} -buildvcs=false
