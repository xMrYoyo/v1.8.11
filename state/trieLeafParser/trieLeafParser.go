package trieLeafParser

import (
	"bytes"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/keyValStorage"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/ElrondNetwork/elrond-go/common"
	"github.com/ElrondNetwork/elrond-go/errors"
	"github.com/ElrondNetwork/elrond-go/state/dataTrieValue"
)

type trieLeafParser struct {
	address             []byte
	marshaller          marshal.Marshalizer
	enableEpochsHandler common.EnableEpochsHandler
}

// NewTrieLeafParser returns a new instance of trieLeafParser
func NewTrieLeafParser(address []byte, marshaller marshal.Marshalizer, enableEpochsHandler common.EnableEpochsHandler) (*trieLeafParser, error) {
	if check.IfNil(marshaller) {
		return nil, errors.ErrNilMarshalizer
	}
	if check.IfNil(enableEpochsHandler) {
		return nil, errors.ErrNilEnableEpochsHandler
	}

	return &trieLeafParser{
		address:             address,
		marshaller:          marshaller,
		enableEpochsHandler: enableEpochsHandler,
	}, nil
}

// ParseLeaf returns a new KeyValStorage with the actual key and value
func (tlp *trieLeafParser) ParseLeaf(trieKey []byte, trieVal []byte) (core.KeyValueHolder, error) {
	if tlp.enableEpochsHandler.IsAutoBalanceDataTriesEnabled() {
		data := &dataTrieValue.TrieLeafData{}
		err := tlp.marshaller.Unmarshal(data, trieVal)
		if err == nil {
			return keyValStorage.NewKeyValStorage(data.Key, data.Value), nil
		}
	}

	suffix := append(trieKey, tlp.address...)
	lenSuffix := len(suffix)
	if lenSuffix == 0 {
		return keyValStorage.NewKeyValStorage(trieKey, trieVal), nil
	}

	lenValue := len(trieVal)
	position := bytes.Index(trieVal, suffix)
	if position != lenValue-lenSuffix || position < 0 {
		return nil, core.ErrSuffixNotPresentOrInIncorrectPosition
	}

	return keyValStorage.NewKeyValStorage(trieKey, trieVal[:position]), nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (tlp *trieLeafParser) IsInterfaceNil() bool {
	return tlp == nil
}
