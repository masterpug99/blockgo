package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash         []byte         // 블록의 정보들을 해쉬한 뒤 바이트 슬라이스에 담는다
	Transactions []*Transaction // 데이터 부분 -> 비트코인에서는 거래 내역들
	PrevHash     []byte         // 전 블록의 정보들의 해쉬값 -> 블럭들은 연결되어 있다(링크드 리스트)
	Nonce        int            // pow에 부합하는 조건을 만들기 위해 추가하는 난수
}

// 블록의 거래내역들을 합친 뒤 해시에 넣은 후 바이트 슬라이스 반환
func (b *Block) HashTransaction() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

// 블럭 하나 만들기
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 맨 처음 블럭 만들기
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// 블럭을 인코딩 해줌 -> 직렬화 과정 (파일에 저장할때 사용)
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// 바이트 슬라이스를 블럭 구조체로 디코딩 해줌 -> 역직렬화 과정
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
