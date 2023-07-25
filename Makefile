check_install: 
	which swagger || (go get github.com/go-swagger/go-swagger && echo "you need to install swagger-codegen as well") 

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models