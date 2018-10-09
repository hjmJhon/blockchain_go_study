package types

type UTXO struct {
	TxHash string
	Index  int
	Output *TxOutput
}
