# flyweight
物件之間，若有共同的部分可以共享，則將可共用的部分獨立為共享物件，
不能共享的部份外部化，使用時再將外部化的部分傳給共享物件。

# 應用場景
當情境涵括
- 一個應用程序使用大量的對象
- 完全由於使用大量對象，造成很大的存儲開銷
- 對象的大多數狀態都可變為外部狀態
- 如果刪除對象的外部狀態，那麼可以用相對較少的共享對象取代很多組對象
- 應用程序不依賴於對象標誌


# 舉例:
黑白棋，在棋盤上放 3 顆黑棋， 3 顆白棋。一般情況，可能會實例化 6 顆獨立的棋子物件，每一個物件都有：黑棋或白棋、X作標、Y作標，三個資料。享元模式則是將共用的資料獨立為共用物件，假設"黑棋"兩個字是黑棋共用的，"白棋"兩個字是白棋共用的，故規劃成共享物件；X、Y座標每個棋子都不一樣，所以獨立為外部資料，使用時再傳給共享物件。如此便可減少記憶體中的重覆資料。

> 實現重點在於，區分共享資料和不可以共享的資料。共享資料做成共享物件，不可共享資料則使用時再傳給共享物件。實例化共享物件時，則由享元工廠把關，判斷目前始否已經實例化過該共享物件，若有則直接回傳現存的共享物件。

# Unified Modeling Language
```
FlyweightFactory <---flyweights---> Flyweight
[GetFlyweight(key)]                 [Operation(extrinsicState)]
    |            |                                |
    |            |                                |
    |       if (flyweight[key] exists) {          |
    |           return existing flyweight;        |
    |       } else {                              |
    |           create new flyweight;             |
    |           add it to pool of flyweights;     |
    |           return the new flyweight;         |
    |       }                                     |
    |                                             |
    |                                             |
    |                                       |------------------------------|
    |                                   ConcreteFlyweight          UnsharedConcreteFlyweight
    |                                   Operation(extrinsicState)  Operaion(extrinsicState)
    |                                   intrinsicState             allState
  client --------------------------------------|---------------------------|
```
1. Flyweight: 描述一個接口，通過這個接口Flyweight可以接受並作用於外部狀態
2. ConcreteFlyweight: 實現Flyweight接口，並為內部狀態增加存儲空間。該對象必須是可共享的。它所存儲的狀態必須是內部的，即必須獨立於對象的場景
3. UnsharedConcreteFlyweight: 並非所有的Flyweight子類都需要被共享。Flyweight接口使共享成為可能，但他並不強制共享
4. FlyweightFactory: 創建並管理Flyweight對象 & 確保合理的共享Flyweight
5. Client: 維持一個對Flyweight的引用 & 計算或存儲Flyweight的外部狀態



# refer:
- https://xyz.cinc.biz/2013/07/flyweight-pattern.html
- https://www.jianshu.com/p/f88b903a166a
- https://www.cnblogs.com/gaochundong/p/design_pattern_flyweight.html


# keynotes:
1. 透過創建一個共享map，預儲存一些可以使用的物件
2. 定義一個方法來拿取共享元件裡面的物件
```go
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

type ImageFlyweight struct {
	data string
}

var imageFactory *ImageFlyweightFactory
```
- ImageFlyweightFactor: 定義出一個存儲享元的資料結構
- ImageFlyweight: 做字串映射()
- imageFactory: 實例上面的變數做享元模式

```go
func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

// ImageViewer: 紀錄ImageFlyweight的地址
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

// GetImageFlyweightFactory 返回 共享的工廠存儲變數
func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

// Get 返回 指定享元 ImageFlyweight 的位址
func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}

	return image
}
```
> flow: NewImageViewer("imageName") -> GetImageFlyweightFactory().Get("imageName") -> return NewImageFlyweight("imageName") -> &ImageFlyweight{data: "imageName"}

- NewImageViewer: 製造新的Image景觀
- GetImageFlyweightFactor(): 拿到`imageFactory`
- GetImageFlyweightFactory().Get(filename): 獲取`imageFactory`這個全域變數內的指定享元，通常只有初始化會呼叫一次，生成實例
- NewImageFlyweight(filename): 如果拿不到特定的`享元`，就做一個新的，並且將該`享元`存回全域變數`imageFlyweight`
- &ImageFlyweight{data: data}: 回傳一個新的`享元`
