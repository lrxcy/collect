package main

import "fmt"

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

type ImageFlyweight struct {
	data string
}

var imageFactory *ImageFlyweightFactory

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

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}

// GetImageFlyweightFactory 返回 共享的工廠
func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}

	return image
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

func main() {
	viewer1 := NewImageViewer("image1.png")
	viewer1.Display()

	viewer2 := NewImageViewer("image1.png")
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		fmt.Println("Failed")
	}
}
