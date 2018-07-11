export GOPATH=`pwd`
# go get github.com/gin-gonic/gin
gofmt -w src/
go test -v -cover -coverprofile=coverage.out  ./...
# go install main
