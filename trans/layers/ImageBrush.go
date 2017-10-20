package layers

import (
	"fmt"
	"errors"
	"image"
	"image/color"
)

type ImageBrush struct {
	Height,Width uint
	PixelChan chan color.NRGBA
	ByteChan chan byte
	IntChan chan int
}

func NewImageBrush(x,y uint) *ImageBrush {
	ib := new(ImageBrush)
	ib.Height,ib.Width=x,y
	return ib
}

func (IB *ImageBrush) Draw(data []byte) (*image.NRGBA,error) {
	//check data length
	if len(data)==0 || len(data)%3 != 0 {
		return nil,errors.New("The length of the data must be a multiple of 3")
	}

	height,width := int(IB.Height),int(IB.Width)
	//check height and width
	PixelOfRect := height*width
	SizeOfRect,SizeOfData := PixelOfRect*3, len(data)
	if SizeOfRect<=0 || SizeOfData>SizeOfRect{
		return nil,errors.New(fmt.Sprintf("The given Rect(%dB) is not large enough to store all the data(%dB)",SizeOfRect,SizeOfData))
	}
	
	//drawing
	img := image.NewNRGBA(image.Rect(0, 0, height, width))
	IB.PixelChan = make(chan color.NRGBA,22)
	go IB.SerializeTo(data)
	okmark := true
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			pixel,okmark := <-IB.PixelChan
			if !okmark {
				break;
			}
			img.Set(i,j,pixel)
		}
		if !okmark {
			break;
		}
	}
	return img,nil
}

func (IB *ImageBrush) SerializeTo(data []byte) error{
	if len(data)%3 != 0 {
		return errors.New("The length of the data must be a multiple of 3")
	}
	for i := 0; i < len(data); i+=3 {
		IB.PixelChan <- color.NRGBA{
            R:  data[i],
            G:  data[i+1],
            B:  data[i+2],
            A:  255,
		}
	}
	close(IB.PixelChan)
	return nil
}

func (IB *ImageBrush) Extract(img image.Image) (error) {
	//check image bounds
	height,width := img.Bounds().Max.X, img.Bounds().Max.Y
	// if width != int(IB.Width) || height != int(IB.Height) {
	// 	return errors.New(fmt.Sprintf("target height/width does not match (height=%v,width=%v)",height,width))
	// }
	if width <=0 || height <=0 {
		return errors.New("cannot Extract data from nil")
	}

	//Unserialize Data
	IB.ByteChan,IB.IntChan = make(chan byte),make(chan int)
	go IB.extract(height,width,img)
	return nil
}

func (IB *ImageBrush) extract(height,width int, img image.Image) {
	num,ok := 0,true;
	defer func() {
		fmt.Println("extract: closing IB.ByteChan")
		close(IB.ByteChan)
	}()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if num <= 0 {
				//fmt.Println("UnserializeTo: getting q.size")
				num,ok = <- IB.IntChan
				if !ok {
					return
				}
				if num <= 0 {
					return
				}
				num += num%3
			}
			//fmt.Println("UnserializeTo: writing to c,num=",num)

			r,g,b,_ := img.At(i,j).RGBA()
			//fmt.Println("extract: read pixel,","At",i,j," RGB",r,g,b)
			IB.ByteChan <- byte(r)
			IB.ByteChan <- byte(g)
			IB.ByteChan <- byte(b)
			num-=3
		}
	}
	return
}

func (IB *ImageBrush)ReadFromBytesChan(size int) []byte {
	buf := []byte{}
	// fmt.Println("ReadFromBytesChan: writing to q.size, size=",size)
	if size==0{
		close(IB.IntChan)
	}else{
		IB.IntChan <- size
	}
	

	for i := 0; i < size; i++ {
		v,ok := <- IB.ByteChan
		if !ok {
			return buf
		}
		buf=append(buf,v)
	}
	//fmt.Println("ReadFromBytesChan: return buf",buf[:])
	return buf
}