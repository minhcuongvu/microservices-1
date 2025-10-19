module github.com/minhcuongvu/microservices-1/internal/services/product_service

go 1.25.3

replace github.com/minhcuongvu/microservices-1/internal/pkg => ../../pkg

require github.com/minhcuongvu/microservices-1/internal/pkg v0.0.0-00010101000000-000000000000

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
