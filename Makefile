run :
	go run cmd/api/main.go


wire:
	cd pkg/di && wire


swag: 
	swag init -g pkg/api/server.go -o ./cmd/api/docs