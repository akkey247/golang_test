/* 課題 初級編2 素数判定 */

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  n := 10
  var data []int
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < n; i++ {
    data = append(data, rand.Intn(100))
  }
  QuickSort(0, n-1, &data)
  for j, v := range data {
    fmt.Print(fmt.Sprintf("%d: %d\n", j, v))
  }
}

//n個のランダム数を並び替えた配列を返す
func QuickSort(bottom int, top int, data *[]int) {
  if bottom >= top {
    return
  }
  for i, v := range data {
    data[i] = 10
  }
}

/*
function QuickSort($bottom=0,$top,&$data){
    if($bottom >= $top){
        return;
    }
    //先頭の値を「適当な値」とする
    $div = $data[$bottom];

    //$data[$bottom]番目より大きい値を後ろに持っていく
    for($lower = $bottom, $upper = $top; $lower < $upper;){
        //$lower番目の値が適当に選択した値以下の限り
        while($lower <= $upper && $data[$lower] <= $div){
            $lower++; //最終的な$lowerの値は$divよりも大きい値の配列番号になる
        }
        //$upper番目の値が適当に選択した値より大きい限り
        while($lower <= $upper && $data[$upper] > $div){
            $upper--;//最終的な$lowerの値は$divよりも小さい値の配列番号になる
        }
        //もし$divより大きな値が、$divより小さい値よりも前に有った場合、順番を入れ替える
        if($lower < $upper){
            $temp = $data[$lower];
            $data[$lower] = $data[$upper];
            $data[$upper] = $temp;
        }
    }
    //最初に選択した値を中央に移動
    $temp = $data[$bottom];
    $data[$bottom] = $data[$upper];
    $data[$upper] = $temp;
    QuickSort($bottom, $upper -1, $data);
    QuickSort($upper + 1, $top, $data);
}

$sort = array();
$N = 10000;
//ソート準備
for($i = 0; $i < $N; $i++){
    $sort[$i] = rand();
}

QuickSort(0, $N-1, $sort);

var_dump($sort);
*/