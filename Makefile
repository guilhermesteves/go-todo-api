#**********************************************************************************************************************************
#  Common
#**********************************************************************************************************************************

ensure-docker:
		./scripts/ensure_docker.sh

#**********************************************************************************************************************************
#  API
#**********************************************************************************************************************************

build-api:
		go build -o ./build/bin/api ./cmd/api/api.go

run-api:
		cd ./cmd/api/; go run api.go

test-api:
		go test -coverprofile=coverage.out  -v -coverpkg=./pkg/... ./test/...