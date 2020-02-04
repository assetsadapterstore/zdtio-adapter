package addrdec

import (
	"fmt"
	"github.com/blocktree/openwallet/openwallet"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	ZDTPublicKeyPrefix       = "PUB_"
	ZDTPublicKeyK1Prefix     = "PUB_K1_"
	ZDTPublicKeyR1Prefix     = "PUB_R1_"
	ZDTPublicKeyPrefixCompat = "ZDT"

	//ZDT stuff
	ZDT_mainnetPublic = addressEncoder.AddressType{"zdt", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(ZDTPublicKeyPrefixCompat), nil}
	// ZDT_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// ZDT_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	openwallet.AddressDecoderV2Base
	IsTestNet bool
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string, opts ...interface{}) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, ZDTPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(ZDTPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, ZDTPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(ZDTPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, ZDTPublicKeyPrefixCompat) { // "ZDT"
		pubKeyMaterial = pubKey[len(ZDTPublicKeyPrefixCompat):] // strip "ZDT"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", ZDTPublicKeyK1Prefix, ZDTPublicKeyR1Prefix, ZDTPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(ZDT_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, ZDT_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, ZDT_mainnetPublic.ChecksumType))
	return string(ZDT_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", ZDT_mainnetPublic.Alphabet), nil
}
