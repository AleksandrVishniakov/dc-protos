all: gen-auth gen-orch gen-daemon


gen-auth:
	protoc -I proto proto/auth/v1/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-orch:
	protoc -I proto proto/orchestrator/v1/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-daemon:
	protoc -I proto proto/daemon/v1/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative