package database

import (
	"crypto/ecdsa"
	"github.com/ardanlabs/blockchain/foundation/blockchain/signature"
	"math/big"
)

// Tx is the transactional information between two parties.
type Tx struct {
	ChainID uint16 `json:"chain_id"` // Ethereum: The chain id that is listed in the genesis file.
	Nonce   uint64 `json:"nonce"`    // Ethereum: Unique id for the transaction supplied by the user.
	FromID  string `json:"from"`     // Ethereum: Account sending the transaction. Will be checked against signature.
	ToID    string `json:"to"`       // Ethereum: Account receiving the benefit of the transaction.
	Value   uint64 `json:"value"`    // Ethereum: Monetary value received from this transaction.
	Tip     uint64 `json:"tip"`      // Ethereum: Tip offered by the sender as an incentive to mine this transaction.
	Data    []byte `json:"data"`     // Ethereum: Extra data related to the transaction.
}

// NewTx constructs a new transaction.
func NewTx(chainID uint16, nonce uint64, fromID string, toID string, value uint64, tip uint64, data []byte) (Tx, error) {

	return Tx{
		ChainID: chainID,
		Nonce:   nonce,
		FromID:  fromID,
		ToID:    toID,
		Value:   value,
		Tip:     tip,
		Data:    data,
	}, nil
}

// SignedTx is a signed version of the transaction. This is how clients like
// a wallet provide transactions for inclusion into the blockchain.
type SignedTx struct {
	Tx
	V *big.Int `json:"v"` // Ethereum: Recovery identifier, either 29 or 30 with ardanID.
	R *big.Int `json:"r"` // Ethereum: First coordinate of the ECDSA signature.
	S *big.Int `json:"s"` // Ethereum: Second coordinate of the ECDSA signature.
}

// Sign uses the specified private key to sign the transaction.
func (tx Tx) Sign(privateKey *ecdsa.PrivateKey) (SignedTx, error) {

	// Sign the transaction with the private key to produce a signature.
	v, r, s, err := signature.Sign(tx, privateKey)
	if err != nil {
		return SignedTx{}, err
	}

	// Construct the signed transaction by adding the signature
	// in the [R|S|V] format.
	signedTx := SignedTx{
		Tx: tx,
		V:  v,
		R:  r,
		S:  s,
	}

	return signedTx, nil
}
