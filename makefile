docker: 
	docker build --tag lynx/crudapp . 
	docker run --network="host" lynx/crudapp
swagger: 
	swag init -d ./cmd/api -o ./docs --parseDependency
native:
	go build -o app/app ./cmd/api/main.go 
	./app/app