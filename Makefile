run: 
	go run cmd/go-jwt.go

push:
	git add . && git commit -m "$(commit)" && git push origin main