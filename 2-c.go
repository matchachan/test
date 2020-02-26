package main
import "fmt"

func main(){
  var bubble [3]int
  for i:=0;i<3;i++{
    fmt.Scanf("%d",&bubble[i])
  }
  for i:=0;i<3;i++{
    fmt.Printf("%d\n",bubble[i])
  }
}
