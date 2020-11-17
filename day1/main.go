package main

import "fmt"

//7.重构代码

func main() {
	//调用NewBlockChain，返回一个bc（缩写）
	bc := NewBlockChain()

	bc.AddBlock("A向B转了50枚比特币！")
	bc.AddBlock("A又向B转了50枚比特币！")

	//返回的值是数组，需要逐个遍历
	for i, block := range bc.blocks  {
		fmt.Printf("====== 当前区块高度:  %d=====\n",i)
		fmt.Printf("前区块Hash值: %x\n",block.PrevHash)
		fmt.Printf("当前区块Hash值: %x\n",block.Hash)
		fmt.Printf("区块数据: %s\n",block.Data)

	}



}