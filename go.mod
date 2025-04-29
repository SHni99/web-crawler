module github.com/nishenghan/go-web-crawler

go 1.23

toolchain go1.23.1

require (
	github.com/PuerkitoBio/goquery v1.8.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)

replace github.com/nishenghan/go-web-crawler/proto => ./proto
