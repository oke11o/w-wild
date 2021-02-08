module github.com/oke11o/w-wild/internal/users

go 1.14

require (
	cloud.google.com/go/firestore v1.4.0
	firebase.google.com/go v3.13.0+incompatible
	github.com/go-chi/chi v4.1.0+incompatible
	github.com/go-chi/render v1.0.1
	github.com/oke11o/w-wild/internal/common v0.0.0-20210128115539-5ba76e9472c0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.5.0
	google.golang.org/api v0.36.0
	google.golang.org/grpc v1.33.2
)

replace github.com/oke11o/w-wild/internal/common => ../common/
