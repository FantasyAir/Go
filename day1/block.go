package main

import (
	"crypto/sha256"
)

//1.定义结构
type Block struct {
	//1.前区块哈希
	PrevHash []byte
	//2.当前区块哈希
	Hash []byte
	//3.数据
	Data []byte

}

//2.创建区块
func NewBlock(data string, preBlockHash []byte) *Block  {
	//肯定是需要返回的Block的
	//创建一个区块需要前区块哈希、数据;这俩个数据是没有的，肯定基于外部数据(data string, preBlockHash []byte)

	//定义一个block,接收传参
	block := Block{
		PrevHash : preBlockHash,
		Hash : []byte{},//先填空，后面再计算 //TODO
		Data : []byte(data),
	}

	//在此处调用生成hash
	block.SetHash()

	return &block //返回给&block指针
}
//3.生成哈希
func (block *Block) SetHash()  {
	//写个函数，无需返回值，原因是Block内部哈希

	//1.拼装数据  由于下面sum256需要而创建的；
	//制造blockInfo数据的由来：表示把每个Data的数组，打碎逐一追加到上个区块中
	blockInfo := append(block.PrevHash,block.Data...)

	//2.sha256
	//func Sum256(data []byte) [Size]byte  要的切片，返回数组
	//所以blockInfo是创建的要的数据，
	hash := sha256.Sum256(blockInfo) //返回一个hash值
	block.Hash = hash[:]  //返回的hash值赋予给Block结构体中的Hash

}