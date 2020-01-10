// 	參考網址

package main

import(
	"flag"
	"fmt"
)

// 透過宣告 flag.String(), flag.Int() or flag.Bool() 來宣告需要使用的型別
// flag with default para name "gerry" and default age "20"
// e.g. Input_pstrname = flag.String("name","gerry","input ur name")
// 在此處宣告一個變數Input_pstrname存放flag "name" 
var Input_pstrName = flag.String("name","gerry","input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_falgvar int

func Init(){
	flag.IntVar(&Input_falgvar, "flagname", 1234, "help message for flagname")
}

// 夾帶flag的使用方法
// 1. -flag xxx (使用空格間隔要放置的參數，flag前面放置一個 - 號) ; only support for non-bool variable
// 2. --flag xxx (使用空格間隔要放置的參數，flag前面放置兩個 -- 號) ; 
// 3. -flag=xxx (使用等號再放置參數，flag前面放置一個 - 號) ; 
// 4. --flag=xxx (使用等號再放置參數，flag前面放置兩個 -- 號) ; 
// 5. 打開命令介面

func main(){
	Init()

	// 透過flag.Parse()來對命令參數解析
	flag.Parse()

	fmt.Println("args=%s, num=%d\n", flag.Args(), flag.NArg())

	for i:=0; i!=flag.NArg(); i++ {
		fmt.Println("arg[%d]=%s\n", i, flag.Arg(i))
		fmt.Println(i)
	}

	fmt.Println("name =", *Input_pstrName)
	fmt.Println("age =", *Input_piAge)
	// fmt.Println("flagname =", *Input_falgvar)

}
