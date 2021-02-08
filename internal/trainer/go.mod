module github.com/oke11o/w-wild/internal/trainer

go 1.14

require (
	cloud.google.com/go/firestore v1.4.0
	github.com/deepmap/oapi-codegen v1.5.0
	github.com/go-chi/chi v4.1.0+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.4.3
	github.com/oke11o/w-wild/internal/common v0.0.0-20210207104356-b472c9e58f10
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.5.0
	google.golang.org/api v0.39.0
	google.golang.org/grpc v1.35.0
)

replace github.com/oke11o/w-wild/internal/common => ../common/
