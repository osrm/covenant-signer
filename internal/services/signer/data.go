package signer

// SignUnbondingTxPayload carries all data necessary to sign unbonding transaction
type SignUnbondingTxRequest struct {
	StakingOutputPkScriptHex string `json:"staking_output_pk_script_hex"`
	UnbondingTxHex           string `json:"unbonding_tx_hex"`
	// 33 bytes compressed public key
	CovenantPublicKey string `json:"covenant_public_key"`
}

// SignUnbondingTxResponse covenant member schnorr signature
type SignUnbondingTxResponse struct {
	SignatureHex string `json:"signature_hex"`
}
