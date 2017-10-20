package layers

import (
	"fmt"
	"bytes"
	"errors"
	"image/png"
	"math"
)

type PNG struct {
	Height,Width uint
}

func NewSepPNG(x,y uint) *PNG {
	if x<=0 || y<=0 {
		return nil
	}
	return &PNG{x,y}
}

func NewPNG(size uint) *PNG {
	size+=SizeOfDataHeader
	size=(size-1)/3+1
	r:=uint(math.Sqrt(float64(size))+1)
	return &PNG{r,r}
}

func (p *PNG) Encode(data []byte,fix bool) ([]byte,error) {
	//generate bytes of Data Header
	HeaderBytes,err := NewDataHeader(data)
	if err != nil {
		return nil,err
	}
	//fmt.Println("Encode: getting data header length=",len(HeaderBytes)," bytes=",HeaderBytes)

	//prepend data header, and pad it with zero
	data = append(HeaderBytes,data...)
	data = ZeroPadding(data,3)
	
	//fix image size
	if fix {
		size:=(len(data)-1)/3+1
		p.Width=uint(math.Sqrt(float64(size))+1)
		p.Height=p.Width
	}

	//drawing
	IB := NewImageBrush(p.Height,p.Width)
	img,err := IB.Draw(data)
	if err != nil {
		return nil,err
	}

	//write to buffer
	f := bytes.NewBuffer(make([]byte,0))
	if err := png.Encode(f, img); err != nil {
	    return nil,err
	}
	return f.Bytes(),nil
}

func (p *PNG) Decode(b []byte,fix bool) ([]byte,error) {
	//load data
	reader := bytes.NewReader(b)
	img, err := png.Decode(reader)
	if err != nil {
		return nil,err
	}
	if fix{
		p.Height,p.Width = uint(img.Bounds().Max.X), uint(img.Bounds().Max.Y)
	}
	//prepare image brush
	IB := NewImageBrush(p.Height,p.Width)
	if err:=IB.Extract(img);err != nil {
		return nil,err
	}
	
	//get data header
	dhbytes := IB.ReadFromBytesChan(SizeOfDataHeader)
	dh,err := GetDataHeader(dhbytes)
	if err != nil {
		return nil,err
	}
	if uint(dh.Length) > p.Height*p.Width*3-SizeOfDataHeader {
		return nil,errors.New("target data header corrupted! Wrong data length")
	}
	// fmt.Println("Decode: getting date header",dh)
	
	//get data body
	databytes := IB.ReadFromBytesChan(int(dh.Length))
	//fmt.Println("Decode: gettiing date body",databytes[:512])

	//close IB Extract routine
	IB.ReadFromBytesChan(0)
	
	//verify signature
	datasig := DHash(databytes)
	if bytes.Compare(datasig,dh.Signature) != 0 {
		fmt.Println("target Signature=",datasig)
		fmt.Println("origin Signature=",dh.Signature)
		return nil,errors.New("target Signature does not match!")
	}
	return databytes,nil
}