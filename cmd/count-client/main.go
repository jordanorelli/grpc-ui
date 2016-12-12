package main

import (
    "fmt"
    "flag"
    "time"
    "log"
    "os"
    "context"

    "github.com/therecipe/qt/widgets"
    "github.com/jordanorelli/grpc-ui/lib/count"
    "google.golang.org/grpc"
)

var (
    info_log *log.Logger
    error_log *log.Logger
)

func bg(label *widgets.QLabel) {
    info_log.Println("bg goroutine started")
    defer info_log.Println("bg goroutine exited")

    time.Sleep(5*time.Second)
    conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
    if err != nil {
        label.SetText(fmt.Sprintf("error: %v", err))
        return
    }

    req := count.NextRequest{Name: "example-param"}
    client := count.NewCountClient(conn)
    for range time.Tick(time.Second) {
        reply, err := client.Next(context.TODO(), &req)
        if err != nil {
            label.SetText(fmt.Sprintf("error: %v", err))
            continue
        }
        label.SetText(fmt.Sprintf("count: %v", reply.GetVal()))
    }

    label.SetText(fmt.Sprintf("Connection: %v", conn))
}

func main() {
    var options struct {
        infoLogPath string
        errorLogPath string
    }
    flag.StringVar(&options.infoLogPath, "info-log", "count-client.log", "file path of an info log")
    flag.StringVar(&options.errorLogPath, "error-log", "count-client.log", "file path of an error log")
    flag.Parse()

    {
        f, err := os.OpenFile(options.infoLogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
        if err != nil {
            panic(err)
        }
        info_log = log.New(f, "i ", log.LstdFlags)
    }

    {
        f, err := os.OpenFile(options.errorLogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
        if err != nil {
            panic(err)
        }
        error_log = log.New(f, "e ", log.LstdFlags)
    }

    widgets.NewQApplication(len(os.Args), os.Args)

    label := widgets.NewQLabel(nil, 0)
    label.SetText("this is a label")
    label.Show()

    info_log.Println("spawning bg goroutine")
    go bg(label)

    info_log.Println("starting qapplication")
    code := widgets.QApplication_Exec()
    info_log.Printf("QApplication exited with code: %d\n", code)
    os.Exit(code)
}