package main

import (
    "context"
    "flag"
    "log"
    "net"
    "net/http"

    "github.com/PuerkitoBio/goquery"
    "google.golang.org/grpc"
    pb "github.com/nishenghan/go-web-crawler/proto"
)

type server struct {
    pb.UnimplementedCrawlerServer
}

func (s *server) FetchAndParse(ctx context.Context, req *pb.CrawlRequest) (*pb.CrawlResult, error) {
    url := req.GetUrl()
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    var links []string
    doc.Find("a[href]").Each(func(i int, sel *goquery.Selection) {
        href, _ := sel.Attr("href")
        links = append(links, href)
    })

    return &pb.CrawlResult{Url: url, Links: links}, nil
}

func main() {
    port := flag.String("port", "50051", "server port")
    flag.Parse()

    lis, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterCrawlerServer(s, &server{})

    log.Printf("server listening on %s", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
