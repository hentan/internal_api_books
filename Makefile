LOCAL_BIN := $(shell pwd)/bin

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1

generate-go-book-service-api-proto:
# комментарии для себя на будущее
# mkdir -p src/go/pkg/grpc/book_service создали директорию для хранения сгенерированных файлов
# proto_path api/grpc/protobuf/book_service указали путь к директории с исходниками .proto файлов
# --go_out=src/go/pkg/grpc/book_service сгенерирует Go файлы для описанных типов данных и сообщений.
# --go_opt=paths=source_relative указали, что путь к исходникам будет относительно директории с исходниками
# --go-grpc_out=src/go/pkg/grpc/book_service генерирует GRPC серверный и клиентский код для описанных типов данных и сообщений.
# --plugin=protoc-gen-go=bin/protoc-gen-go указали пути до плагина protoc-gen-go
# protoc-gen-go: основной генератор Go кода.
# protoc-gen-go-grpc: генератор GRPC сервера и клиента.
# protoc-gen-doc: генератор документации (не стандартный протобаф плагин).
# --doc_out=api/grpc/protobuf/book_service генерирует документацию в формате markdown
	mkdir -p src/go/pkg/grpc/book_service
	GOBIN=$(LOCAL_BIN) protoc \
  --proto_path api/grpc/protobuf/book_service \
  --go_out=src/go/pkg/grpc/book_service \
  --go_opt=paths=source_relative \
  --go-grpc_out=src/go/pkg/grpc/book_service \
  --go-grpc_opt=paths=source_relative \
  --plugin=protoc-gen-go=bin/protoc-gen-go \
  --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
  --plugin=protoc-gen-doc=bin/protoc-gen-doc \
  --doc_out=api/grpc/protobuf/book_service \
  --doc_opt=markdown,README.md,source_relative \
  api/grpc/protobuf/book_service/*.proto