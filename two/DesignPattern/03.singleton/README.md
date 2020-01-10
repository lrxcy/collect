# singleton 單例模式
需求: 某一個類只應該存在一個實例的時候; 單例模式: 確保一個類只有一個實例，並提供一個全局訪問點
> 使用懶惰模式的單例模式，使用雙重檢查加鎖保證線程安全

# refer:
- https://www.jianshu.com/p/10908f3b5399

# keynotes:
定義出一個全局變數`singleton`並使用`sync.Once`來同步物件
```go
// 1. 定義出要實作Singleton的struct
type Singleton struct {
    a int
}

// 2. 針對該struct宣告出一個全局的變量
var singleton *Singleton


// 3. 宣告once當作sync.Once物件
var once sync.Once

// 4. 定義出拿取Singleton物件的方法
func GetInstance() *Singleton {
    once.Do(func(){
        singleton = &Singleton{}
    })
    return singleton
}

/* 實作 */
ins1 := GetInstance()
ins2 := GetInstance()
if ins1 != ins2 {
    log.Println("Failed")
}
ins2.a = 1
log.Println(ins1.a) // 輸出應為 1

```