run:
	go run app/alexandria-server/main.go

live:
	curl -il http://localhost:3000/v1/liveness