#repositories
mockgen -source internal/repositories/contract.go -destination mocks/repositories/contract.go
#pkg
mockgen -source pkg/storage/storage.go -destination mocks/pkg/storage/storage.go