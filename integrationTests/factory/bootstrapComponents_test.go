package factory

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/core"
	mainFactory "github.com/ElrondNetwork/elrond-go/factory"
	"github.com/ElrondNetwork/elrond-go/process/headerCheck"
	"github.com/ElrondNetwork/elrond-go/sharding"
	storageFactory "github.com/ElrondNetwork/elrond-go/storage/factory"
	"github.com/ElrondNetwork/elrond-go/storage/storageUnit"
	"github.com/stretchr/testify/require"
)

// ------------ Test CryptoComponents --------------------
func TestBootstrapComponents_Create_ShouldWork(t *testing.T) {
	//t.Skip()

	generalConfig, _ := loadMainConfig(configPath)
	ratingsConfig, _ := loadRatingsConfig(ratingsPath)
	economicsConfig, _ := loadEconomicsConfig(economicsPath)
	prefsConfig, _ := loadPreferencesConfig(prefsPath)
	p2pConfig, _ := core.LoadP2PConfig(p2pPath)
	systemSCConfig, _ := loadSystemSmartContractsConfig(systemSCConfigPath)

	coreComponents, _ := createCoreComponents(*generalConfig, *ratingsConfig, *economicsConfig)
	cryptoComponents, err := createCryptoComponents(*generalConfig, *systemSCConfig, coreComponents)
	networkComponents, err := createNetworkComponents(*generalConfig, *p2pConfig, *ratingsConfig, coreComponents)

	bootstrapComponents, err := createBootstrapComponents(
		*generalConfig,
		prefsConfig.Preferences,
		coreComponents,
		cryptoComponents,
		networkComponents)

	require.Nil(t, err)
	require.NotNil(t, bootstrapComponents)
}

func TestBootstrapComponents_Create_Close_ShouldWork(t *testing.T) {
	//	t.Skip()

	_ = logger.SetLogLevel("*:DEBUG")

	generalConfig, _ := loadMainConfig(configPath)
	ratingsConfig, _ := loadRatingsConfig(ratingsPath)
	economicsConfig, _ := loadEconomicsConfig(economicsPath)
	prefsConfig, _ := loadPreferencesConfig(prefsPath)
	p2pConfig, _ := core.LoadP2PConfig(p2pPath)
	systemSCConfig, _ := loadSystemSmartContractsConfig(systemSCConfigPath)

	coreComponents, _ := createCoreComponents(*generalConfig, *ratingsConfig, *economicsConfig)
	cryptoComponents, _ := createCryptoComponents(*generalConfig, *systemSCConfig, coreComponents)
	networkComponents, _ := createNetworkComponents(*generalConfig, *p2pConfig, *ratingsConfig, coreComponents)
	time.Sleep(2 * time.Second)

	nrBefore := runtime.NumGoroutine()
	bootstrapComponents, _ := createBootstrapComponents(
		*generalConfig,
		prefsConfig.Preferences,
		coreComponents,
		cryptoComponents,
		networkComponents)
	time.Sleep(2 * time.Second)
	err := bootstrapComponents.Close()
	require.Nil(t, err)
	time.Sleep(5 * time.Second)
	nrAfter := runtime.NumGoroutine()

	if nrBefore != nrAfter {
		printStack()
	}

	require.Equal(t, nrBefore, nrAfter)

}

func createBootstrapComponents(
	config config.Config,
	preferencesConfig config.PreferencesConfig,
	managedCoreComponents mainFactory.CoreComponentsHandler,
	managedCryptoComponents mainFactory.CryptoComponentsHandler,
	managedNetworkComponents mainFactory.NetworkComponentsHandler,
) (mainFactory.BootstrapComponentsHandler, error) {

	nodesSetup := managedCoreComponents.GenesisNodesSetup()

	nodesShuffler := sharding.NewHashValidatorsShuffler(
		nodesSetup.MinNumberOfShardNodes(),
		nodesSetup.MinNumberOfMetaNodes(),
		nodesSetup.GetHysteresis(),
		nodesSetup.GetAdaptivity(),
		true,
	)

	destShardIdAsObserver, err := core.ProcessDestinationShardAsObserver(preferencesConfig.DestinationShardAsObserver)
	if err != nil {
		return nil, err
	}

	versionsCache, err := storageUnit.NewCache(storageFactory.GetCacherFromConfig(config.Versions.Cache))
	if err != nil {
		return nil, err
	}

	headerIntegrityVerifier, err := headerCheck.NewHeaderIntegrityVerifier(
		[]byte(managedCoreComponents.ChainID()),
		config.Versions.VersionsByEpochs,
		config.Versions.DefaultVersion,
		versionsCache,
	)
	if err != nil {
		return nil, err
	}

	genesisShardCoordinator, _, err := mainFactory.CreateShardCoordinator(
		managedCoreComponents.GenesisNodesSetup(),
		managedCryptoComponents.PublicKey(),
		preferencesConfig,
		logger.GetOrCreate("bootstrapTest"),
	)

	bootstrapComponentsFactoryArgs := mainFactory.BootstrapComponentsFactoryArgs{
		Config:                  config,
		WorkingDir:              "workingDir",
		DestinationAsObserver:   destShardIdAsObserver,
		GenesisNodesSetup:       nodesSetup,
		NodeShuffler:            nodesShuffler,
		ShardCoordinator:        genesisShardCoordinator,
		CoreComponents:          managedCoreComponents,
		CryptoComponents:        managedCryptoComponents,
		NetworkComponents:       managedNetworkComponents,
		HeaderIntegrityVerifier: headerIntegrityVerifier,
	}

	bootstrapComponentsFactory, err := mainFactory.NewBootstrapComponentsFactory(bootstrapComponentsFactoryArgs)
	if err != nil {
		return nil, fmt.Errorf("NewBootstrapComponentsFactory failed: %w", err)
	}

	managedBootstrapComponents, err := mainFactory.NewManagedBootstrapComponents(bootstrapComponentsFactory)
	if err != nil {
		return nil, err
	}

	err = managedBootstrapComponents.Create()
	if err != nil {
		return nil, err
	}

	return managedBootstrapComponents, nil
}
