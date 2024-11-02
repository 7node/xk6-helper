package helper

import (
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"

	goHex "encoding/hex"
)

// Module is the type for our custom API.
type Helper struct {
	vu modules.VU
}

// Decode returns the decoded string.
func (h *Helper) HexDecode(hex string) []byte {
	decoded, err := goHex.DecodeString(hex)
	if err != nil {
		return nil
	}
	return decoded
}

// HexEncode returns the encoded string.
func (h *Helper) HexEncode(data []byte) string {
	return goHex.EncodeToString(data)
}

// EncodeMessage returns the encoded message.
func (h *Helper) EncodeMessage(payload, encKey, signKey string) []byte {
	encKeyBytes := h.HexDecode(encKey)
	signKeyBytes := h.HexDecode(signKey)

	encoded, err := Parser.Encoder.Encode([]byte(payload), encKeyBytes, signKeyBytes)
	if err != nil {
		common.Throw(h.vu.Runtime(), err)
		return nil
	}

	return encoded
}

// DecodeMessage returns the decoded message.
func (h *Helper) DecodeMessage(payload, encKey, signKey string) []byte {
	encKeyBytes := h.HexDecode(encKey)
	signKeyBytes := h.HexDecode(signKey)

	decoded, err := Parser.Decoder.Decode([]byte(payload), encKeyBytes, signKeyBytes)
	if err != nil {
		common.Throw(h.vu.Runtime(), err)
		return nil
	}

	return decoded
}
