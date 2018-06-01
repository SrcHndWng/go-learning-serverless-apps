account = 1234567890

build:
	env GOOS=linux go build -o bin/images/get functions/images/get/main.go
	env GOOS=linux go build -o bin/images/post functions/images/post/main.go functions/images/post/url.go
	env GOOS=linux go build -o bin/images/update functions/images/update/main.go
	env GOOS=linux go build -o bin/images/id/get functions/images/id/get/main.go
	
	serverless deploy --account ${account}
