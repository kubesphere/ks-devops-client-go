swagger-client-gen:
	# download swagger: go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	rm -rf client models && swagger generate client -f openapi-spec/swagger.json
	go mod tidy
	# test if client packages have compile issue
	go test -run=DOES-NOT-EXIST ./client