package gopdf


import (
	"bytes"
	"fmt"
	"strconv"
)

type PagesObj struct { //impl IObj
	buffer    bytes.Buffer
	PageCount int
	Kids string
	getRoot	 func()(*GoPdf)
}

func (me *PagesObj) Init(funcGetRoot func()(*GoPdf)) {
	me.PageCount = 0
	me.getRoot = funcGetRoot
}

func (me *PagesObj) Build() {

	
	height := fmt.Sprintf("%0.2f",me.getRoot().config.PageSize.H);
	width := fmt.Sprintf("%0.2f",me.getRoot().config.PageSize.W);
	
	me.buffer.WriteString("\t/Type /" + me.GetType() + "\n")
	me.buffer.WriteString("\t/MediaBox [0 0 "+width+" "+height+" ]\n") 
	me.buffer.WriteString("\t/Count "+strconv.Itoa(me.PageCount)+"\n")
	me.buffer.WriteString("\t/Kids [ "+me.Kids+" ]\n") //sample Kids [ 3 0 R ]
}

func (me *PagesObj) GetType() string {
	return "Pages"
}

func (me *PagesObj) GetObjBuff() *bytes.Buffer {
	return &(me.buffer)
}

func (me *PagesObj) Test() {
	fmt.Print(me.GetType() + "\n")
}

