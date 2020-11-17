package main


//4.引入区块链
//定义一个区块链数组,存储所有的区块；同样的数据类型用数组来存储
type BlockChain struct {
	blocks []*Block
}

//5.定义一个区块链,返回一个BlockChain
func NewBlockChain () *BlockChain {
	//创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()

	//返回指针
	return &BlockChain{
		//[]*Block{}填写内容genesisBlock
		blocks: []*Block{genesisBlock},
	}
}

//定义一个创世块
func GenesisBlock() *Block {
	//NewBlock()传参，谁调用return结果返给谁
	return NewBlock("Go创世块",[]byte{})
}


//6.添加区块；  方法一定属于BlockChain，向区块链上追加
func (bc *BlockChain)AddBlock(data string)  {
	//如何获取钱区块的哈希值呢？

	//获取最后一个区块;(通过区块链最后一个数-1，就表示最后一个区块）
	lasteBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lasteBlock.Hash

	//a.创建新的区块，（需要数据外传；需要前Hash，前Hash通过找区块链最后一个区块）
	block := NewBlock(data, prevHash)
	//b.添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}