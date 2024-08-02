# O que é gRPC: Diferença entre REST API e Caso de Uso em Python com Golang

# Protocol buffers vs JSON
arquivos binários são menores que arquivos de texto(JSON)
processo de serializaçao é mais leve que JSON (converter classe pra json)

Fonts:
https://grpc.io/docs/what-is-grpc/introduction/
https://protobuf.dev/
https://www.youtube.com/watch?v=QyyMjF764mo

apt install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=. --go-grpc_out=. ./proto/hello.proto  