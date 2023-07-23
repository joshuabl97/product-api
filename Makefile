swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

get go-swagger: 
	GO111MODULE=off go get github.com/go-swagger/go-swagger
