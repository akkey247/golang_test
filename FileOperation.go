/* お題その① ファイル作成 */

package main

import (
    "fmt"
    "os"
)

func main() {
    if err := os.Mkdir("hoge", 0777); err != nil {
        fmt.Println(err)
    }
}