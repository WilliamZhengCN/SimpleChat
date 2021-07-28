package utils

import (
	message "SimpleChat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [6000]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	pkgLen := uint32(len(data))
	//var bytes [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	count, err := this.Conn.Write(this.Buf[:4])
	if count != 4 || err != nil {
		fmt.Println("Fail to Write. Err= ", err)
	}

	count, err = this.Conn.Write(data)
	if count != int(pkgLen) || err != nil {
		fmt.Println("Fail to Write. Err= ", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 5000)
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("fail to get length of message. error is :", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("fail to read message from server. error is :", err)
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("fail to Unmarshal. error is ", err)
	}
	return
}
