module gocrontab

go 1.13

require (
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.uber.org/zap v1.13.0 // indirect
)

require google.golang.org/grpc v1.26.0

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
