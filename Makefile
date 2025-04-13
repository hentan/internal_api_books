LOCAL_BIN := $(shell pwd)/bin

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1

# комментарии для себя на будущее
# mkdir -p src/go/pkg/grpc/books создали директорию для хранения сгенерированных файлов
# proto_path api/grpc/protobuf/books указали путь к директории с исходниками .proto файлов
# --go_out=src/go/pkg/grpc/books сгенерирует Go файлы для описанных типов данных и сообщений.
# --go_opt=paths=source_relative указали, что путь к исходникам будет относительно директории с исходниками
# --go-grpc_out=src/go/pkg/grpc/books генерирует GRPC серверный и клиентский код для описанных типов данных и сообщений.
# --plugin=protoc-gen-go=bin/protoc-gen-go указали пути до плагина protoc-gen-go
# protoc-gen-go: основной генератор Go кода.
# protoc-gen-go-grpc: генератор GRPC сервера и клиента.
# protoc-gen-doc: генератор документации (не стандартный протобаф плагин).
# --doc_out=api/grpc/protobuf/books генерирует документацию в формате markdown
generate-go-book-service-api-proto:
	mkdir -p gen/go

	protoc \
        -Iapi/grpc/protobuf \
        --go_out=gen/go \
        --go_opt=paths=source_relative \
        --go-grpc_out=gen/go \
        --go-grpc_opt=paths=source_relative \
        --doc_out=api/grpc/protobuf \
        --doc_opt=html,index.html \
        --plugin=protoc-gen-go=bin/protoc-gen-go \
        --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
        --plugin=protoc-gen-doc=bin/protoc-gen-doc \
        api/grpc/protobuf/authors/authors.proto \
        api/grpc/protobuf/books/books.proto \
        api/grpc/protobuf/combines/combine.proto