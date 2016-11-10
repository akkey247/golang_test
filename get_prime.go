/* 課題 初級編2 素数判定 */

package main

import (
    "fmt"
)

func main() {
  result := getPrime(10)
  //配列ループの書式(for キー, 値 := range 配列名 {})
  for i, v := range result {
    fmt.Print(fmt.Sprintf("%d: %d\n", i, v))
  }
}

//0～nまでの素数の配列を返す
//関数の書式(func 関数名(引数名 引数型) 返り値型{})
func getPrime(n int) []int {
  var isPrime bool
  var primeLists []int
  //通常ループの書式
  for i := 0; i < n; i++ {
    isPrime = true
    for j := 2; j < i; j++ {
      // 割り切れた数が存在したらアウト
      if i % j == 0 {
        isPrime = false;
        break;
      }
      if (isPrime) {
        primeLists = append(primeLists, i)
        break
      }
    }
  }
  return primeLists
}
