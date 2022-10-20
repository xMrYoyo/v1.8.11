package heartbeat

import (
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-crypto"
	"github.com/ElrondNetwork/elrond-go/common"
	"github.com/ElrondNetwork/elrond-go/sharding/nodesCoordinator"
	"github.com/ElrondNetwork/elrond-go/state"
)

// P2PMessenger defines a subset of the p2p.Messenger interface
type P2PMessenger interface {
	Broadcast(topic string, buff []byte)
	BroadcastUsingPrivateKey(topic string, buff []byte, pid core.PeerID, skBytes []byte)
	ID() core.PeerID
	Sign(payload []byte) ([]byte, error)
	ConnectedPeersOnTopic(topic string) []core.PeerID
	SignUsingPrivateKey(skBytes []byte, payload []byte) ([]byte, error)
	IsInterfaceNil() bool
}

// PeerTypeProviderHandler defines what a component which computes the type of a peer should do
type PeerTypeProviderHandler interface {
	ComputeForPubKey(pubKey []byte) (common.PeerType, uint32, error)
	GetAllPeerTypeInfos() []*state.PeerTypeInfo
	IsInterfaceNil() bool
}

// HardforkTrigger defines the behavior of a hardfork trigger
type HardforkTrigger interface {
	TriggerReceived(payload []byte, data []byte, pkBytes []byte) (bool, error)
	RecordedTriggerMessage() ([]byte, bool)
	NotifyTriggerReceivedV2() <-chan struct{}
	CreateData() []byte
	IsInterfaceNil() bool
}

// CurrentBlockProvider can provide the current block that the node was able to commit
type CurrentBlockProvider interface {
	GetCurrentBlockHeader() data.HeaderHandler
	IsInterfaceNil() bool
}

// NodeRedundancyHandler defines the interface responsible for the redundancy management of the node
type NodeRedundancyHandler interface {
	IsRedundancyNode() bool
	IsMainMachineActive() bool
	ObserverPrivateKey() crypto.PrivateKey
	IsInterfaceNil() bool
}

// NodesCoordinator defines the behavior of a struct able to do validator selection
type NodesCoordinator interface {
	GetAllEligibleValidatorsPublicKeys(epoch uint32) (map[uint32][][]byte, error)
	GetAllWaitingValidatorsPublicKeys(epoch uint32) (map[uint32][][]byte, error)
	GetValidatorWithPublicKey(publicKey []byte) (validator nodesCoordinator.Validator, shardId uint32, err error)
	IsInterfaceNil() bool
}

// PeerShardMapper saves the shard for a peer ID
type PeerShardMapper interface {
	PutPeerIdShardId(pid core.PeerID, shardID uint32)
	IsInterfaceNil() bool
}

// ManagedPeersHolder defines the operations of an entity that holds managed identities for a node
type ManagedPeersHolder interface {
	AddManagedPeer(privateKeyBytes []byte) error
	GetPrivateKey(pkBytes []byte) (crypto.PrivateKey, error)
	GetP2PIdentity(pkBytes []byte) ([]byte, core.PeerID, error)
	GetMachineID(pkBytes []byte) (string, error)
	GetNameAndIdentity(pkBytes []byte) (string, string, error)
	IncrementRoundsWithoutReceivedMessages(pkBytes []byte)
	ResetRoundsWithoutReceivedMessages(pkBytes []byte)
	GetManagedKeysByCurrentNode() map[string]crypto.PrivateKey
	IsKeyManagedByCurrentNode(pkBytes []byte) bool
	IsKeyRegistered(pkBytes []byte) bool
	IsPidManagedByCurrentNode(pid core.PeerID) bool
	IsKeyValidator(pkBytes []byte) bool
	SetValidatorState(pkBytes []byte, state bool)
	GetNextPeerAuthenticationTime(pkBytes []byte) (time.Time, error)
	SetNextPeerAuthenticationTime(pkBytes []byte, nextTime time.Time)
	IsInterfaceNil() bool
}

// ShardCoordinator defines the operations of a shard coordinator
type ShardCoordinator interface {
	SelfId() uint32
	ComputeId(address []byte) uint32
	IsInterfaceNil() bool
}
