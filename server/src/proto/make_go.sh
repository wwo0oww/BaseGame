protoc --proto_path=./  --go_out=.. ./proto/net/*.proto
protoc --proto_path=./proto --proto_path=./ --go_out=. ./proto/rpc/*.proto