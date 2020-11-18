package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

//1.定义结构
type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevHash []byte
	//3.Merkel根（梅克尔根，这就是一个哈希值，先不管V4介绍
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数，也就是挖矿要找的数据
	Nonce uint64

	//a.当前区块哈希,正常比特币区块中没有当前区块的哈希，我们为了是方便做了简化
	Hash []byte
	//b.数据
	Data []byte

}

//实现一个辅助函数，功能是将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	//网络传输过程;一个字节8位，十六进制就是2个字符，
	err := binary.Write(&buffer,binary.BigEndian,num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

//2.创建区块
func NewBlock(data string, preBlockHash []byte) *Block  {
	//肯定是需要返回的Block的
	//创建一个区块需要前区块哈希、数据;这俩个数据是没有的，肯定基于外部数据(data string, preBlockHash []byte)

	//定义一个block,接收传参
	block := Block{
		Version : 00,
		PrevHash : preBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0,//随便填写无效值
		Nonce:0,//同上
		Hash : []byte{},//先填空，后面再计算 //TODO
		Data : []byte(data),
	}

	//在此处调用生成hash
	block.SetHash()

	return &block //返回给&block指针
}
//3.生成哈希
func (block *Block) SetHash()  {
	//var blockInfo []byte

	//写个函数，无需返回值，原因是Block内部哈希

	//1.拼装数据  由于下面sum256需要而创建的；
	//制造blockInfo数据的由来：表示把每个Data的数组，打碎逐一追加到上个区块中
	/*blockInfo = append(blockInfo,Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo,block.PrevHash...)
	blockInfo = append(blockInfo,block.MerkelRoot...)
	blockInfo = append(blockInfo,Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo,Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo,Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo,block.Data...)
	*/

	//优化上面代码
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	//将二维的切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp,[]byte{})

	//2.sha256
	//func Sum256(data []byte) [Size]byte  要的切片，返回数组
	//所以blockInfo是创建的要的数据，
	hash := sha256.Sum256(blockInfo) //返回一个hash值
	block.Hash = hash[:]  //返回的hash值赋予给Block结构体中的Hash

}