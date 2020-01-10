# facade 表象模式
- 使用場景 : 方便使用者直接呼叫某些函數，可以執行多件函數後面的工作

```
某天你心血來潮，想在家裡準備家庭劇院組。你為此做了一番研究，找來了覺得適合的播放器、投影機、螢幕、音響等等的設備，當你要準備開始看電影時，你要做哪些事呢？
將燈光調暗
開啟螢幕
打開投影機
設定投影機輸入模式
打開音響
設定音響音量
打開播放器
開始播放

這樣看起來做的事好像沒很多，但是當看完電影後，要把所有設備關掉要怎麼做？全部反向做一次嗎？假如要聽音樂而已，也要這麼麻煩嗎？未來升級新設備時，還要重新學習操作流程嗎？

這時候最直覺的想法一定是：把這些事包成一個 function 就好啦。這就是表象模式的精神所在，將一個或數個類別複雜的一切都隱藏起來，只露出美好的表面(就是簡化介面啦)。

使用表象模式，可以將一個複雜的次系統，變得容易使用。表象類別提供更合理的介面，來簡化原先的複雜介面。假如不想用表象介面時，還是可以直接操作次系統。

在這邊特別說明一下，表象模式跟轉接器模式雖然都是用來封裝類別，但是他們的目的卻是不同的。轉接器模式的目的是改變介面以符合客戶的期望。而表象模式是提供次系統一個簡化的介面。
```

# refer
- http://corrupt003-design-pattern.blogspot.com/2016/07/facade-pattern.html

# keypoint
在golang使用facade，只要注意有定義到所有需要的方法有確實被加入facade的interface就好
```go
// 1. 製作出一個外觀，供使用者呼叫及使用
type Modules struct {
    AModule
    BModule
}

// 2. 針對創建出來的Module定義出方法Test()
func (m Module) Test() string {
    return m.AModule.TestA() + m.BModule.TestB()
}

// 3. --定義生成函數--
type ModulesAPI interface {
    Test() string
}

func NewModuleInterface() ModulesAPI {
    return Modules{}
}
```