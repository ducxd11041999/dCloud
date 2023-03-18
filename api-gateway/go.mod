module my-project/dCloud/api-gateway

go 1.19

replace my-project/dCloud/protobuf => ../protobuf

require (
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.53.0
	my-project/dCloud/protobuf v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
