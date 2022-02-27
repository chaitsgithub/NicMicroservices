.PHONY: swagger
swagger:
	GO111MODULE=off swagger generate spec -o ./swagger/swagger.yaml -c server -c handlers -c main -c docs -c models --scan-models

.PHONY: swagger-server
swagger-serve:
	swagger serve -F=swagger ./swagger/swagger.yaml

.PHONY: swagger-docs
swagger-docs:
	swagger generate markdown -f ./swagger/swagger.yaml --output ./swagger/swagger.md


# check_install:
# 	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

# swagger: check_install
# 	swagger generate spec -o ./swagger.yaml -c server -c handlers -c main -c docs -c models --scan-models

# swagger-serve: swagger
# 	swagger serve -F=swagger ./swagger.yaml

# swagger-docs: swagger-serve
# 	swagger generate markdown -f ./swagger.yaml --output ./swagger.md