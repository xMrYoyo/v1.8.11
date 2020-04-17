package metachain

import (
	"bytes"
	"crypto/rand"
	"errors"
	"testing"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/data/block"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/process/mock"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/ElrondNetwork/elrond-go/storage"
	"github.com/ElrondNetwork/elrond-go/storage/memorydb"
	"github.com/ElrondNetwork/elrond-go/storage/storageUnit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createGenesisBlocks(shardCoordinator sharding.Coordinator) map[uint32]data.HeaderHandler {
	genesisBlocks := make(map[uint32]data.HeaderHandler)
	for shardId := uint32(0); shardId < shardCoordinator.NumberOfShards(); shardId++ {
		genesisBlocks[shardId] = createGenesisBlock(shardId)
	}

	genesisBlocks[core.MetachainShardId] = createGenesisMetaBlock()

	return genesisBlocks
}

func createGenesisBlock(shardId uint32) *block.Header {
	rootHash := []byte("roothash")
	return &block.Header{
		Nonce:         0,
		Round:         0,
		Signature:     rootHash,
		RandSeed:      rootHash,
		PrevRandSeed:  rootHash,
		ShardID:       shardId,
		PubKeysBitmap: rootHash,
		RootHash:      rootHash,
		PrevHash:      rootHash,
	}
}

func createGenesisMetaBlock() *block.MetaBlock {
	rootHash := []byte("roothash")
	return &block.MetaBlock{
		Nonce:         0,
		Round:         0,
		Signature:     rootHash,
		RandSeed:      rootHash,
		PrevRandSeed:  rootHash,
		PubKeysBitmap: rootHash,
		RootHash:      rootHash,
		PrevHash:      rootHash,
	}
}

func createMockEpochStartCreatorArguments() ArgsNewEpochStartData {
	shardCoordinator := mock.NewOneShardCoordinatorMock()
	startHeaders := createGenesisBlocks(shardCoordinator)
	argsNewEpochStartData := ArgsNewEpochStartData{
		Marshalizer:       &mock.MarshalizerMock{},
		Hasher:            &mock.HasherStub{},
		Store:             createMetaStore(),
		DataPool:          &mock.PoolsHolderStub{},
		BlockTracker:      mock.NewBlockTrackerMock(shardCoordinator, startHeaders),
		ShardCoordinator:  shardCoordinator,
		EpochStartTrigger: &mock.EpochStartTriggerStub{},
		RequestHandler:    &mock.RequestHandlerStub{},
	}
	return argsNewEpochStartData
}

func createMemUnit() storage.Storer {
	cache, _ := storageUnit.NewCache(storageUnit.LRUCache, 10, 1)
	persist, _ := memorydb.NewlruDB(100000)
	unit, _ := storageUnit.NewStorageUnit(cache, persist)

	return unit
}

func createMetaStore() dataRetriever.StorageService {
	store := dataRetriever.NewChainStorer()
	store.AddStorer(dataRetriever.MetaBlockUnit, createMemUnit())
	store.AddStorer(dataRetriever.BlockHeaderUnit, createMemUnit())

	return store
}

func TestEpochStartData_NilMarshalizer(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.Marshalizer = nil

	esd, err := NewEpochStartData(arguments)
	require.Nil(t, esd)
	require.Equal(t, process.ErrNilMarshalizer, err)
}

func TestEpochStartData_NilHasher(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.Hasher = nil

	esd, err := NewEpochStartData(arguments)
	require.Nil(t, esd)
	require.Equal(t, process.ErrNilHasher, err)
}

func TestEpochStartData_NilStore(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.Store = nil

	esd, err := NewEpochStartData(arguments)
	require.Nil(t, esd)
	require.Equal(t, process.ErrNilStorage, err)
}

func TestEpochStartData_NilDataPool(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.DataPool = nil

	esd, err := NewEpochStartData(arguments)
	require.Nil(t, esd)
	require.Equal(t, process.ErrNilDataPoolHolder, err)
}

func TestEpochStartData_NilBlockTracker(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.BlockTracker = nil

	esd, err := NewEpochStartData(arguments)
	require.Nil(t, esd)
	require.Equal(t, process.ErrNilBlockTracker, err)
}

func TestEpochStartData_NilShardCoordinator(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.ShardCoordinator = nil

	esd, err := NewEpochStartData(arguments)
	require.Nil(t, esd)
	require.Equal(t, process.ErrNilShardCoordinator, err)
}

func TestVerifyEpochStartDataForMetablock_DataDoesNotMatch(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.Hasher = &mock.HasherStub{
		ComputeCalled: func(s string) []byte {
			token := make([]byte, 4)
			_, _ = rand.Read(token)
			return token
		},
	}

	esd, _ := NewEpochStartData(arguments)

	metaBlock := &block.MetaBlock{
		EpochStart: block.EpochStart{
			LastFinalizedHeaders: []block.EpochStartShardData{
				{HeaderHash: []byte("hash")},
			},
		},
	}

	err := esd.VerifyEpochStartDataForMetablock(metaBlock)
	require.Equal(t, process.ErrEpochStartDataDoesNotMatch, err)
}

func TestEpochStartCreator_getLastFinalizedMetaHashForShardMetaHashNotReturnsGenesis(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	epoch, err := NewEpochStartData(arguments)
	require.Nil(t, err)
	require.False(t, check.IfNil(epoch))

	round := uint64(10)

	shardHdr := &block.Header{Round: round}
	last, lastFinal, shardHdrs, err := epoch.lastFinalizedFirstPendingListHeadersForShard(shardHdr)
	assert.Nil(t, last)
	assert.True(t, bytes.Equal(lastFinal, []byte(core.EpochStartIdentifier(0))))
	assert.Equal(t, shardHdr, shardHdrs[0])
	assert.Nil(t, err)
}

func TestEpochStartCreator_getLastFinalizedMetaHashForShardShouldWork(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.EpochStartTrigger = &mock.EpochStartTriggerStub{
		IsEpochStartCalled: func() bool {
			return false
		},
	}

	dPool := &mock.PoolsHolderStub{}
	dPool.TransactionsCalled = func() dataRetriever.ShardedDataCacherNotifier {
		return &mock.ShardedDataStub{}
	}
	metaHash1 := []byte("hash1")
	metaHash2 := []byte("hash2")
	mbHash1 := []byte("mb_hash1")
	dPool.HeadersCalled = func() dataRetriever.HeadersPool {
		cs := &mock.HeadersCacherStub{}
		cs.GetHeaderByHashCalled = func(hash []byte) (handler data.HeaderHandler, e error) {
			return &block.Header{
				PrevHash:         []byte("hash1"),
				Nonce:            2,
				Round:            2,
				PrevRandSeed:     []byte("roothash"),
				MiniBlockHeaders: []block.MiniBlockHeader{{Hash: mbHash1, SenderShardID: 1}},
				MetaBlockHashes:  [][]byte{metaHash1, metaHash2},
			}, nil
		}
		return cs
	}

	arguments.DataPool = dPool

	epoch, _ := NewEpochStartData(arguments)
	round := uint64(10)
	nonce := uint64(1)

	shardHdr := &block.Header{
		Round:           round,
		Nonce:           nonce,
		MetaBlockHashes: [][]byte{mbHash1},
	}
	last, lastFinal, shardHdrs, err := epoch.lastFinalizedFirstPendingListHeadersForShard(shardHdr)
	assert.NotNil(t, last)
	assert.NotNil(t, lastFinal)
	assert.NotNil(t, shardHdrs)
	assert.Nil(t, err)
}

func TestEpochStartCreator_CreateEpochStartFromMetaBlockEpochIsNotStarted(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.EpochStartTrigger = &mock.EpochStartTriggerStub{
		IsEpochStartCalled: func() bool {
			return false
		},
	}

	epoch, _ := NewEpochStartData(arguments)

	epStart, err := epoch.CreateEpochStartData()
	assert.Nil(t, err)

	emptyEpochStart := block.EpochStart{}
	assert.Equal(t, emptyEpochStart, *epStart)
}

func TestEpochStartCreator_CreateEpochStartFromMetaBlockHashComputeIssueShouldErr(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("err computing hash")

	arguments := createMockEpochStartCreatorArguments()
	arguments.Marshalizer = &mock.MarshalizerStub{
		// trigger an error on the Marshal method called from core's ComputeHash
		MarshalCalled: func(obj interface{}) (i []byte, e error) {
			return nil, expectedErr
		},
	}
	arguments.EpochStartTrigger = &mock.EpochStartTriggerStub{
		IsEpochStartCalled: func() bool {
			return true
		},
	}

	epoch, _ := NewEpochStartData(arguments)

	epStart, err := epoch.CreateEpochStartData()
	assert.Nil(t, epStart)
	assert.Equal(t, expectedErr, err)
}

func TestMetaProcessor_CreateEpochStartFromMetaBlockShouldWork(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochStartCreatorArguments()
	arguments.EpochStartTrigger = &mock.EpochStartTriggerStub{
		IsEpochStartCalled: func() bool {
			return true
		},
	}

	hash1 := []byte("hash1")
	hash2 := []byte("hash2")

	startHeaders := createGenesisBlocks(arguments.ShardCoordinator)
	arguments.BlockTracker = mock.NewBlockTrackerMock(arguments.ShardCoordinator, startHeaders)

	hdr := startHeaders[0].(*block.Header)
	hdr.MetaBlockHashes = [][]byte{hash1, hash2}
	hdr.Nonce = 1
	startHeaders[0] = hdr

	dPool := &mock.PoolsHolderStub{}
	dPool.TransactionsCalled = func() dataRetriever.ShardedDataCacherNotifier {
		return &mock.ShardedDataStub{}
	}
	metaHash1 := []byte("hash1")
	metaHash2 := []byte("hash2")
	mbHash1 := []byte("mb_hash1")
	dPool.HeadersCalled = func() dataRetriever.HeadersPool {
		cs := &mock.HeadersCacherStub{}
		cs.GetHeaderByHashCalled = func(hash []byte) (handler data.HeaderHandler, e error) {
			return &block.Header{
				PrevHash:         []byte("hash1"),
				Nonce:            1,
				Round:            1,
				PrevRandSeed:     []byte("roothash"),
				MiniBlockHeaders: []block.MiniBlockHeader{{Hash: mbHash1, SenderShardID: 1}},
				MetaBlockHashes:  [][]byte{metaHash1, metaHash2},
			}, nil
		}

		return cs
	}
	arguments.DataPool = dPool
	metaHdrStorage := arguments.Store.GetStorer(dataRetriever.MetaBlockUnit)
	meta1 := &block.MetaBlock{Nonce: 100}

	var hdrs []block.MiniBlockHeader
	hdrs = append(hdrs, block.MiniBlockHeader{
		Hash:            hash1,
		ReceiverShardID: 0,
		SenderShardID:   1,
		TxCount:         2,
	})
	hdrs = append(hdrs, block.MiniBlockHeader{
		Hash:            mbHash1,
		ReceiverShardID: 0,
		SenderShardID:   1,
		TxCount:         2,
	})
	shardData := block.ShardData{ShardID: 1, ShardMiniBlockHeaders: hdrs}
	meta2 := &block.MetaBlock{Nonce: 101, PrevHash: metaHash1, ShardInfo: []block.ShardData{shardData}}

	marshaledData, _ := arguments.Marshalizer.Marshal(meta1)
	_ = metaHdrStorage.Put(metaHash1, marshaledData)

	marshaledData, _ = arguments.Marshalizer.Marshal(meta2)
	_ = metaHdrStorage.Put(metaHash2, marshaledData)

	epoch, _ := NewEpochStartData(arguments)

	epStart, err := epoch.CreateEpochStartData()
	assert.Nil(t, err)
	assert.NotNil(t, epStart)
	assert.Equal(t, hash1, epStart.LastFinalizedHeaders[0].LastFinishedMetaBlock)
	assert.Equal(t, hash2, epStart.LastFinalizedHeaders[0].FirstPendingMetaBlock)
	assert.Equal(t, 1, len(epStart.LastFinalizedHeaders[0].PendingMiniBlockHeaders))

	err = epoch.VerifyEpochStartDataForMetablock(&block.MetaBlock{EpochStart: *epStart})
	assert.Nil(t, err)
}
