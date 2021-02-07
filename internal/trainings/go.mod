module github.com/oke11o/w-wild/internal/trainings

go 1.14

require (
	cloud.google.com/go/firestore v1.4.0
	github.com/deepmap/oapi-codegen v1.5.0
	github.com/go-chi/chi v4.1.0+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0
	github.com/oke11o/w-wild/internal/common v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	google.golang.org/api v0.39.0
)

replace github.com/oke11o/w-wild/internal/common => ../common/
