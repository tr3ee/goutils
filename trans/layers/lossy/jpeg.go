package lossy

// 由于 image/jepg 包采用DCT算法(有损变换)，huffman编码，故本包不适用于无损转码

import (
	// "fmt"
	"bytes"
	// "errors"
	"image/jpeg"
	"tr3e/trans/layers"
)

type JPEG struct {
	Height,Width uint
}

func NewJPEG(x,y uint) *JPEG {
	if x<=0 || y<=0 {
		return nil
	}
	j := new(JPEG)
	j.Height, j.Width = x, y
	return j
}

func (p *JPEG) Encode(data []byte) ([]byte,error) {
	data = layers.ZeroPadding(data,3)
	//drawing
	IB := layers.NewImageBrush(p.Height,p.Width)
	img,err := IB.Draw(data)
	if err != nil {
		return nil,err
	}

	//write to buffer
	f := bytes.NewBuffer(make([]byte,0))
	if err := jpeg.Encode(f, img,nil); err != nil {
	    return nil,err
	}
	return f.Bytes(),nil
}

// func (p *JPEG) Decode(b []byte) ([]byte,error) {
// 	//load data
// 	reader := bytes.NewReader(b)
// 	img, err := jpeg.Decode(reader)
// 	if err != nil {
// 		return nil,err
// 	}

// 	IB := layers.NewImageBrush(p.Height,p.Width)
// 	if err:=IB.Extract(img);err != nil {
// 		return nil,err
// 	}

// 	//get data header
// 	dhbytes := IB.ReadFromBytesChan(SizeOfDataHeader)
// 	fmt.Println("Decode: get data header=",dhbytes)
// 	dh,err := GetDataHeader(dhbytes)
// 	if err != nil {
// 		return nil,err
// 	}
// 	if uint(dh.Length) > p.Height*p.Width*3-SizeOfDataHeader {
// 		return nil,errors.New(fmt.Sprintln("target data header corrupted! Wrong data length=",dh.Length))
// 	}
// 	// fmt.Println("Decode: getting date header",dh)
	
// 	//get data body
// 	databytes := IB.ReadFromBytesChan(int(dh.Length))
// 	// fmt.Println("Decode: gettiing date body",databytes[:512])
// 	fmt.Printf("%s",databytes)

// 	//close IB Extract routine
// 	IB.ReadFromBytesChan(0)

// 	//verify signature
// 	datasig := DHash(databytes)
// 	if bytes.Compare(datasig,dh.Signature) != 0 {
// 		fmt.Println("target Signature=",datasig)
// 		fmt.Println("origin Signature=",dh.Signature)
// 		return nil,errors.New("target Signature does not match!")
// 	}
// 	return databytes,nil
// }