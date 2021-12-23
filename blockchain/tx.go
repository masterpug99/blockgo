package blockchain

// references of the previous output
type TxInput struct {
	ID  []byte
	Out int // index of the transaction
	Sig string
}

type TxOutput struct {
	Value  int
	PubKey string // needed to unlock the value field
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
