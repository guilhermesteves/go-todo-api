# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.8.0/protoc-3.8.0-osx-x86_64.zip

# Unzip
unzip protoc-3.8.0-osx-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

rm -rf protoc3
rm -rf protoc-3.8.0-osx-x86_64.zip

GIT_TAG="v1.2.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go

sudo mv "$(go env GOPATH)"/bin/protoc-gen-go /usr/local/bin/
