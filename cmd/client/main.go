package main

import (
    "context"
    "flag"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "github.com/nishenghan/go-web-crawler/proto"
)

func main() {
    addr := flag.String("addr", "localhost:50051", "server address")
    url := flag.String("url", "", "URL to crawl")
    timeout := flag.Duration("timeout", 10*time.Second, "timeout for request")
    flag.Parse()

    if *url == "" {
        log.Fatalf("please provide -url")
    }

    ctx, cancel := context.WithTimeout(context.Background(), *timeout)
    defer cancel()

    conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewCrawlerClient(conn)
    resp, err := client.FetchAndParse(ctx, &pb.CrawlRequest{Url: *url})
    if err != nil {
        log.Fatalf("error during FetchAndParse: %v", err)
    }

    log.Printf("Fetched %s, found %d links:", resp.GetUrl(), len(resp.GetLinks()))
    for _, link := range resp.GetLinks() {
        log.Println(link)
    }
}
