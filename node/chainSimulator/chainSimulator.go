package chainSimulator

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/endProcess"
	"github.com/multiversx/mx-chain-go/config"
	"github.com/multiversx/mx-chain-go/node/chainSimulator/components"
	"github.com/multiversx/mx-chain-go/node/chainSimulator/configs"
	"github.com/multiversx/mx-chain-go/node/chainSimulator/testdata"
)

type simulator struct {
	chanStopNodeProcess    chan endProcess.ArgEndProcess
	syncedBroadcastNetwork components.SyncedBroadcastNetworkHandler
	nodes                  []ChainHandler
	numOfShards            uint32
}

// NewChainSimulator will create a new instance of simulator
func NewChainSimulator(numOfShards uint32, pathToInitialConfig string) (*simulator, error) {
	syncedBroadcastNetwork := components.NewSyncedBroadcastNetwork()

	instance := &simulator{
		syncedBroadcastNetwork: syncedBroadcastNetwork,
		nodes:                  make([]ChainHandler, 0),
		numOfShards:            numOfShards,
		chanStopNodeProcess:    make(chan endProcess.ArgEndProcess),
	}

	err := instance.createChainHandlers(numOfShards, pathToInitialConfig)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (s *simulator) createChainHandlers(numOfShards uint32, originalConfigPath string) error {
	outputConfigs, err := configs.CreateChainSimulatorConfigs(configs.ArgsChainSimulatorConfigs{
		NumOfShards:               numOfShards,
		OriginalConfigsPath:       originalConfigPath,
		GenesisAddressWithStake:   testdata.GenesisAddressWithStake,
		GenesisAddressWithBalance: testdata.GenesisAddressWithBalance,
	})
	if err != nil {
		return err
	}

	metaChainHandler, err := s.createChainHandler(core.MetachainShardId, outputConfigs.Configs, 0, outputConfigs.GasScheduleFilename)
	if err != nil {
		return err
	}

	s.nodes = append(s.nodes, metaChainHandler)

	for idx := uint32(0); idx < numOfShards; idx++ {
		shardChainHandler, errS := s.createChainHandler(idx, outputConfigs.Configs, int(idx)+1, outputConfigs.GasScheduleFilename)
		if errS != nil {
			return errS
		}

		s.nodes = append(s.nodes, shardChainHandler)
	}

	return nil
}

func (s *simulator) createChainHandler(shardID uint32, configs *config.Configs, skIndex int, gasScheduleFilename string) (ChainHandler, error) {
	args := components.ArgsTestOnlyProcessingNode{
		Config:                   *configs.GeneralConfig,
		EpochConfig:              *configs.EpochConfig,
		EconomicsConfig:          *configs.EconomicsConfig,
		RoundsConfig:             *configs.RoundConfig,
		PreferencesConfig:        *configs.PreferencesConfig,
		ImportDBConfig:           *configs.ImportDbConfig,
		ContextFlagsConfig:       *configs.FlagsConfig,
		SystemSCConfig:           *configs.SystemSCConfig,
		ConfigurationPathsHolder: *configs.ConfigurationPathsHolder,
		ChanStopNodeProcess:      s.chanStopNodeProcess,
		SyncedBroadcastNetwork:   s.syncedBroadcastNetwork,
		NumShards:                s.numOfShards,
		GasScheduleFilename:      gasScheduleFilename,
		ShardID:                  shardID,
		SkIndex:                  skIndex,
	}

	return components.NewTestOnlyProcessingNode(args)
}

// GenerateBlocks will generate the provided number of blocks
func (s *simulator) GenerateBlocks(_ int) error {
	return nil
}

// Stop will stop the simulator
func (s *simulator) Stop() {
}

// IsInterfaceNil returns true if there is no value under the interface
func (s *simulator) IsInterfaceNil() bool {
	return s == nil
}
