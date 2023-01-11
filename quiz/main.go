package main
import (
	"fmt"
	"flag"

)
func main(){
	fmt.Println("hello world")
	fileName := flag.String("csv", "problems.csv", "file name of the csv file")
	flag.Parse()
	_ = fileName
}