package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    pb "protof-test/proto"
)

func main() {
    conn, err := grpc.Dial(":9527", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("dial error: %v\n", err)
    }
    defer func() {
        _ = conn.Close()
    }()
    
    // 实例化 UserInfoService 微服务的客户端
    client := pb.NewUserInfoServiceClient(conn)
    
    // 调用服务
    req := new(pb.UserRequest)
    req.Name = "Happysir"
    resp, err := client.GetUserInfo(context.Background(), req)
    if err != nil {
        log.Fatalf("resp error: %v\n", err)
    }
    
    fmt.Printf("Recevied: %v\n", resp)
}