package main

import (
    "fmt"
    "os"

    "github.com/therecipe/qt/widgets"
)

func main() {
    fmt.Println("Hey from VSCode")
    widgets.NewQApplication(len(os.Args), os.Args)
    widgets.QApplication_Exec()
}