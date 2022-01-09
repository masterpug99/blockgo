package blockchain

import "github.com/dgraph-io/badger"

//badger DB doesn't have key to seperate data you need to use Prefix
var (
	utxoPrefix   = []byte("utxo-")
	prefixLength = len(utxoPrefix)
)

type UTXOSet struct {
	Blockchain *BlockChain
}

func (u *UTXOSet) DeleteByPrefix(preifx []byte) {
	deletKeys := func(keysForDelete [][]byte) error {
		if err := u.Blockchain.Database.Update(func(txn *badger.Txn) error {
			for _, key := range keysForDelete {
				if err := txn.Delete(key); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}
}
