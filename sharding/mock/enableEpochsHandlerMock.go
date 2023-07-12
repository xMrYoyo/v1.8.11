package mock

// EnableEpochsHandlerMock -
type EnableEpochsHandlerMock struct {
	WaitingListFixEnableEpochField            uint32
	RefactorPeersMiniBlocksEnableEpochField   uint32
	IsRefactorPeersMiniBlocksFlagEnabledField bool
	IsSCProcessorV2FlagEnabledField           bool
}

// BlockGasAndFeesReCheckEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) BlockGasAndFeesReCheckEnableEpoch() uint32 {
	return 0
}

// StakingV2EnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) StakingV2EnableEpoch() uint32 {
	return 0
}

// ScheduledMiniBlocksEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) ScheduledMiniBlocksEnableEpoch() uint32 {
	return 0
}

// SwitchJailWaitingEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) SwitchJailWaitingEnableEpoch() uint32 {
	return 0
}

// BalanceWaitingListsEnableEpoch returns WaitingListFixEnableEpochField
func (mock *EnableEpochsHandlerMock) BalanceWaitingListsEnableEpoch() uint32 {
	return 0
}

// WaitingListFixEnableEpoch returns WaitingListFixEnableEpochField
func (mock *EnableEpochsHandlerMock) WaitingListFixEnableEpoch() uint32 {
	return mock.WaitingListFixEnableEpochField
}

// MultiESDTTransferAsyncCallBackEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) MultiESDTTransferAsyncCallBackEnableEpoch() uint32 {
	return 0
}

// FixOOGReturnCodeEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) FixOOGReturnCodeEnableEpoch() uint32 {
	return 0
}

// RemoveNonUpdatedStorageEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) RemoveNonUpdatedStorageEnableEpoch() uint32 {
	return 0
}

// CreateNFTThroughExecByCallerEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) CreateNFTThroughExecByCallerEnableEpoch() uint32 {
	return 0
}

// FixFailExecutionOnErrorEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) FixFailExecutionOnErrorEnableEpoch() uint32 {
	return 0
}

// ManagedCryptoAPIEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) ManagedCryptoAPIEnableEpoch() uint32 {
	return 0
}

// DisableExecByCallerEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) DisableExecByCallerEnableEpoch() uint32 {
	return 0
}

// RefactorContextEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) RefactorContextEnableEpoch() uint32 {
	return 0
}

// CheckExecuteReadOnlyEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) CheckExecuteReadOnlyEnableEpoch() uint32 {
	return 0
}

// StorageAPICostOptimizationEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) StorageAPICostOptimizationEnableEpoch() uint32 {
	return 0
}

// MiniBlockPartialExecutionEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) MiniBlockPartialExecutionEnableEpoch() uint32 {
	return 0
}

// RefactorPeersMiniBlocksEnableEpoch returns 0
func (mock *EnableEpochsHandlerMock) RefactorPeersMiniBlocksEnableEpoch() uint32 {
	return mock.RefactorPeersMiniBlocksEnableEpochField
}

// GetCurrentEpoch -
func (mock *EnableEpochsHandlerMock) GetCurrentEpoch() uint32 {
	return 0
}

// IsSCDeployFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSCDeployFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsBuiltInFunctionsFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsBuiltInFunctionsFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsFixOOGReturnCodeFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsFixOOGReturnCodeFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRelayedTransactionsFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRelayedTransactionsFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsPenalizedTooMuchGasFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsPenalizedTooMuchGasFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSwitchJailWaitingFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSwitchJailWaitingFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsBelowSignedThresholdFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsBelowSignedThresholdFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsTransactionSignedWithTxHashFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsTransactionSignedWithTxHashFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsMetaProtectionFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsMetaProtectionFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsAheadOfTimeGasUsageFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsAheadOfTimeGasUsageFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRepairCallbackFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRepairCallbackFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsReturnDataToLastTransferFlagEnabledAfterEpoch returns false
func (mock *EnableEpochsHandlerMock) IsReturnDataToLastTransferFlagEnabledAfterEpoch(_ uint32) bool {
	return false
}

// IsSenderInOutTransferFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSenderInOutTransferFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsStakeFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsStakeFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsStakingV2FlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsStakingV2FlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsStakingV2FlagEnabledAfterEpoch -
func (mock *EnableEpochsHandlerMock) IsStakingV2FlagEnabledAfterEpoch(_ uint32) bool {
	return false
}

// IsDoubleKeyProtectionFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsDoubleKeyProtectionFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsGovernanceFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsGovernanceFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsDelegationManagerFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsDelegationManagerFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsDelegationSmartContractFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsDelegationSmartContractFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCorrectLastUnJailedFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCorrectLastUnJailedFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRelayedTransactionsV2FlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRelayedTransactionsV2FlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsUnBondTokensV2FlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsUnBondTokensV2FlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSaveJailedAlwaysFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSaveJailedAlwaysFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsReDelegateBelowMinCheckFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsReDelegateBelowMinCheckFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsValidatorToDelegationFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsValidatorToDelegationFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsWaitingListFixFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsWaitingListFixFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsIncrementSCRNonceInMultiTransferFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsIncrementSCRNonceInMultiTransferFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTMultiTransferFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTMultiTransferFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsGlobalMintBurnFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsGlobalMintBurnFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTTransferRoleFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTTransferRoleFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsBuiltInFunctionOnMetaFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsBuiltInFunctionOnMetaFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsComputeRewardCheckpointFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsComputeRewardCheckpointFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSCRSizeInvariantCheckFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSCRSizeInvariantCheckFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsBackwardCompSaveKeyValueFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsBackwardCompSaveKeyValueFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTNFTCreateOnMultiShardFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTNFTCreateOnMultiShardFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsMetaESDTSetFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsMetaESDTSetFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsAddTokensToDelegationFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsAddTokensToDelegationFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsMultiESDTTransferFixOnCallBackFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsMultiESDTTransferFixOnCallBackFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsOptimizeGasUsedInCrossMiniBlocksFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsOptimizeGasUsedInCrossMiniBlocksFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCorrectFirstQueuedFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCorrectFirstQueuedFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsDeleteDelegatorAfterClaimRewardsFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsDeleteDelegatorAfterClaimRewardsFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRemoveNonUpdatedStorageFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRemoveNonUpdatedStorageFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsOptimizeNFTStoreFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsOptimizeNFTStoreFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCreateNFTThroughExecByCallerFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCreateNFTThroughExecByCallerFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsStopDecreasingValidatorRatingWhenStuckFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsStopDecreasingValidatorRatingWhenStuckFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsFrontRunningProtectionFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsFrontRunningProtectionFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsPayableBySCFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsPayableBySCFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCleanUpInformativeSCRsFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCleanUpInformativeSCRsFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsStorageAPICostOptimizationFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsStorageAPICostOptimizationFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTRegisterAndSetAllRolesFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTRegisterAndSetAllRolesFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsScheduledMiniBlocksFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsScheduledMiniBlocksFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCorrectJailedNotUnStakedEmptyQueueFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCorrectJailedNotUnStakedEmptyQueueFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsAddFailedRelayedTxToInvalidMBsFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsAddFailedRelayedTxToInvalidMBsFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSCRSizeInvariantOnBuiltInResultFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSCRSizeInvariantOnBuiltInResultFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCheckCorrectTokenIDForTransferRoleFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCheckCorrectTokenIDForTransferRoleFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsFailExecutionOnEveryAPIErrorFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsFailExecutionOnEveryAPIErrorFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsMiniBlockPartialExecutionFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsMiniBlockPartialExecutionFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsManagedCryptoAPIsFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsManagedCryptoAPIsFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTMetadataContinuousCleanupFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTMetadataContinuousCleanupFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsDisableExecByCallerFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsDisableExecByCallerFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRefactorContextFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRefactorContextFlagEnabledInEpoch(_ uint32) bool {
	return mock.IsRefactorPeersMiniBlocksFlagEnabledField
}

// IsCheckFunctionArgumentFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCheckFunctionArgumentFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCheckExecuteOnReadOnlyFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCheckExecuteOnReadOnlyFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSetSenderInEeiOutputTransferFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSetSenderInEeiOutputTransferFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsFixAsyncCallbackCheckFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsFixAsyncCallbackCheckFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSaveToSystemAccountFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSaveToSystemAccountFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCheckFrozenCollectionFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCheckFrozenCollectionFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSendAlwaysFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSendAlwaysFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsValueLengthCheckFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsValueLengthCheckFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsCheckTransferFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsCheckTransferFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsTransferToMetaFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsTransferToMetaFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsESDTNFTImprovementV1FlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsESDTNFTImprovementV1FlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsChangeDelegationOwnerFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsChangeDelegationOwnerFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRefactorPeersMiniBlocksFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRefactorPeersMiniBlocksFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSCProcessorV2FlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSCProcessorV2FlagEnabledInEpoch(_ uint32) bool {
	return mock.IsSCProcessorV2FlagEnabledField
}

// IsFixAsyncCallBackArgsListFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsFixAsyncCallBackArgsListFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsFixOldTokenLiquidityEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsFixOldTokenLiquidityEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRuntimeMemStoreLimitEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRuntimeMemStoreLimitEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRuntimeCodeSizeFixEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRuntimeCodeSizeFixEnabledInEpoch(_ uint32) bool {
	return false
}

// IsMaxBlockchainHookCountersFlagEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsMaxBlockchainHookCountersFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsWipeSingleNFTLiquidityDecreaseEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsWipeSingleNFTLiquidityDecreaseEnabledInEpoch(_ uint32) bool {
	return false
}

// IsAlwaysSaveTokenMetaDataEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsAlwaysSaveTokenMetaDataEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSetGuardianEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsSetGuardianEnabledInEpoch(_ uint32) bool {
	return false
}

// IsRelayedNonceFixEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsRelayedNonceFixEnabledInEpoch(_ uint32) bool {
	return false
}

// IsConsistentTokensValuesLengthCheckEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsConsistentTokensValuesLengthCheckEnabledInEpoch(_ uint32) bool {
	return false
}

// IsKeepExecOrderOnCreatedSCRsEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsKeepExecOrderOnCreatedSCRsEnabledInEpoch(_ uint32) bool {
	return false
}

// IsMultiClaimOnDelegationEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsMultiClaimOnDelegationEnabledInEpoch(_ uint32) bool {
	return false
}

// IsChangeUsernameEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsChangeUsernameEnabledInEpoch(_ uint32) bool {
	return false
}

// IsAutoBalanceDataTriesEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) IsAutoBalanceDataTriesEnabledInEpoch(_ uint32) bool {
	return false
}

// FixDelegationChangeOwnerOnAccountEnabledInEpoch -
func (mock *EnableEpochsHandlerMock) FixDelegationChangeOwnerOnAccountEnabledInEpoch(_ uint32) bool {
	return false
}

// IsSwitchHysteresisForMinNodesFlagEnabledInSpecificEpochOnly returns false
func (mock *EnableEpochsHandlerMock) IsSwitchHysteresisForMinNodesFlagEnabledInSpecificEpochOnly(_ uint32) bool {
	return false
}

// IsGasPriceModifierFlagEnabledInEpoch returns false
func (mock *EnableEpochsHandlerMock) IsGasPriceModifierFlagEnabledInEpoch(_ uint32) bool {
	return false
}

// IsStakingV2OwnerFlagEnabledInSpecificEpochOnly returns false
func (mock *EnableEpochsHandlerMock) IsStakingV2OwnerFlagEnabledInSpecificEpochOnly(_ uint32) bool {
	return false
}

// IsESDTFlagEnabledInSpecificEpochOnly returns false
func (mock *EnableEpochsHandlerMock) IsESDTFlagEnabledInSpecificEpochOnly(_ uint32) bool {
	return false
}

// IsGovernanceFlagEnabledInSpecificEpochOnly returns false
func (mock *EnableEpochsHandlerMock) IsGovernanceFlagEnabledInSpecificEpochOnly(_ uint32) bool {
	return false
}

// IsDelegationSmartContractFlagEnabledInSpecificEpochOnly returns false
func (mock *EnableEpochsHandlerMock) IsDelegationSmartContractFlagEnabledInSpecificEpochOnly(_ uint32) bool {
	return false
}

// IsCorrectLastUnJailedFlagEnabledInSpecificEpochOnly returns false
func (mock *EnableEpochsHandlerMock) IsCorrectLastUnJailedFlagEnabledInSpecificEpochOnly(_ uint32) bool {
	return false
}

// TODO[Sorin]: Remove the methods below

// IsDoubleKeyProtectionFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsDoubleKeyProtectionFlagEnabled() bool {
	return false
}

// IsESDTFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTFlagEnabled() bool {
	return false
}

// IsESDTFlagEnabledForCurrentEpoch returns false
func (mock *EnableEpochsHandlerMock) IsESDTFlagEnabledForCurrentEpoch() bool {
	return false
}

// IsGovernanceFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsGovernanceFlagEnabled() bool {
	return false
}

// IsGovernanceFlagEnabledForCurrentEpoch returns false
func (mock *EnableEpochsHandlerMock) IsGovernanceFlagEnabledForCurrentEpoch() bool {
	return false
}

// IsDelegationManagerFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsDelegationManagerFlagEnabled() bool {
	return false
}

// IsDelegationSmartContractFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsDelegationSmartContractFlagEnabled() bool {
	return false
}

// IsDelegationSmartContractFlagEnabledForCurrentEpoch returns false
func (mock *EnableEpochsHandlerMock) IsDelegationSmartContractFlagEnabledForCurrentEpoch() bool {
	return false
}

// IsCorrectLastUnJailedFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCorrectLastUnJailedFlagEnabled() bool {
	return false
}

// IsCorrectLastUnJailedFlagEnabledForCurrentEpoch returns false
func (mock *EnableEpochsHandlerMock) IsCorrectLastUnJailedFlagEnabledForCurrentEpoch() bool {
	return false
}

// IsRelayedTransactionsV2FlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsRelayedTransactionsV2FlagEnabled() bool {
	return false
}

// IsUnBondTokensV2FlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsUnBondTokensV2FlagEnabled() bool {
	return false
}

// IsSaveJailedAlwaysFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsSaveJailedAlwaysFlagEnabled() bool {
	return false
}

// IsReDelegateBelowMinCheckFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsReDelegateBelowMinCheckFlagEnabled() bool {
	return false
}

// IsValidatorToDelegationFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsValidatorToDelegationFlagEnabled() bool {
	return false
}

// IsWaitingListFixFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsWaitingListFixFlagEnabled() bool {
	return false
}

// IsIncrementSCRNonceInMultiTransferFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsIncrementSCRNonceInMultiTransferFlagEnabled() bool {
	return false
}

// IsESDTMultiTransferFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTMultiTransferFlagEnabled() bool {
	return false
}

// IsGlobalMintBurnFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsGlobalMintBurnFlagEnabled() bool {
	return false
}

// IsESDTTransferRoleFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTTransferRoleFlagEnabled() bool {
	return false
}

// IsBuiltInFunctionOnMetaFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsBuiltInFunctionOnMetaFlagEnabled() bool {
	return false
}

// IsComputeRewardCheckpointFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsComputeRewardCheckpointFlagEnabled() bool {
	return false
}

// IsSCRSizeInvariantCheckFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsSCRSizeInvariantCheckFlagEnabled() bool {
	return false
}

// IsBackwardCompSaveKeyValueFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsBackwardCompSaveKeyValueFlagEnabled() bool {
	return false
}

// IsESDTNFTCreateOnMultiShardFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTNFTCreateOnMultiShardFlagEnabled() bool {
	return false
}

// IsMetaESDTSetFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsMetaESDTSetFlagEnabled() bool {
	return false
}

// IsAddTokensToDelegationFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsAddTokensToDelegationFlagEnabled() bool {
	return false
}

// IsMultiESDTTransferFixOnCallBackFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsMultiESDTTransferFixOnCallBackFlagEnabled() bool {
	return false
}

// IsOptimizeGasUsedInCrossMiniBlocksFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsOptimizeGasUsedInCrossMiniBlocksFlagEnabled() bool {
	return false
}

// IsCorrectFirstQueuedFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCorrectFirstQueuedFlagEnabled() bool {
	return false
}

// IsDeleteDelegatorAfterClaimRewardsFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsDeleteDelegatorAfterClaimRewardsFlagEnabled() bool {
	return false
}

// IsFixOOGReturnCodeFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsFixOOGReturnCodeFlagEnabled() bool {
	return false
}

// IsRemoveNonUpdatedStorageFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsRemoveNonUpdatedStorageFlagEnabled() bool {
	return false
}

// IsOptimizeNFTStoreFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsOptimizeNFTStoreFlagEnabled() bool {
	return false
}

// IsCreateNFTThroughExecByCallerFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCreateNFTThroughExecByCallerFlagEnabled() bool {
	return false
}

// IsStopDecreasingValidatorRatingWhenStuckFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsStopDecreasingValidatorRatingWhenStuckFlagEnabled() bool {
	return false
}

// IsFrontRunningProtectionFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsFrontRunningProtectionFlagEnabled() bool {
	return false
}

// IsPayableBySCFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsPayableBySCFlagEnabled() bool {
	return false
}

// IsCleanUpInformativeSCRsFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCleanUpInformativeSCRsFlagEnabled() bool {
	return false
}

// IsStorageAPICostOptimizationFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsStorageAPICostOptimizationFlagEnabled() bool {
	return false
}

// IsESDTRegisterAndSetAllRolesFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTRegisterAndSetAllRolesFlagEnabled() bool {
	return false
}

// IsScheduledMiniBlocksFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsScheduledMiniBlocksFlagEnabled() bool {
	return false
}

// IsCorrectJailedNotUnStakedEmptyQueueFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCorrectJailedNotUnStakedEmptyQueueFlagEnabled() bool {
	return false
}

// IsDoNotReturnOldBlockInBlockchainHookFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsDoNotReturnOldBlockInBlockchainHookFlagEnabled() bool {
	return false
}

// IsAddFailedRelayedTxToInvalidMBsFlag returns false
func (mock *EnableEpochsHandlerMock) IsAddFailedRelayedTxToInvalidMBsFlag() bool {
	return false
}

// IsSCRSizeInvariantOnBuiltInResultFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsSCRSizeInvariantOnBuiltInResultFlagEnabled() bool {
	return false
}

// IsCheckCorrectTokenIDForTransferRoleFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCheckCorrectTokenIDForTransferRoleFlagEnabled() bool {
	return false
}

// IsFailExecutionOnEveryAPIErrorFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsFailExecutionOnEveryAPIErrorFlagEnabled() bool {
	return false
}

// IsMiniBlockPartialExecutionFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsMiniBlockPartialExecutionFlagEnabled() bool {
	return false
}

// IsManagedCryptoAPIsFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsManagedCryptoAPIsFlagEnabled() bool {
	return false
}

// IsESDTMetadataContinuousCleanupFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTMetadataContinuousCleanupFlagEnabled() bool {
	return false
}

// IsDisableExecByCallerFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsDisableExecByCallerFlagEnabled() bool {
	return false
}

// IsRefactorContextFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsRefactorContextFlagEnabled() bool {
	return false
}

// IsCheckFunctionArgumentFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCheckFunctionArgumentFlagEnabled() bool {
	return false
}

// IsCheckExecuteOnReadOnlyFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCheckExecuteOnReadOnlyFlagEnabled() bool {
	return false
}

// IsFixAsyncCallbackCheckFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsFixAsyncCallbackCheckFlagEnabled() bool {
	return false
}

// IsSaveToSystemAccountFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsSaveToSystemAccountFlagEnabled() bool {
	return false
}

// IsCheckFrozenCollectionFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCheckFrozenCollectionFlagEnabled() bool {
	return false
}

// IsSendAlwaysFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsSendAlwaysFlagEnabled() bool {
	return false
}

// IsValueLengthCheckFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsValueLengthCheckFlagEnabled() bool {
	return false
}

// IsCheckTransferFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsCheckTransferFlagEnabled() bool {
	return false
}

// IsTransferToMetaFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsTransferToMetaFlagEnabled() bool {
	return false
}

// IsESDTNFTImprovementV1FlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsESDTNFTImprovementV1FlagEnabled() bool {
	return false
}

// IsSetSenderInEeiOutputTransferFlagEnabled -
func (mock *EnableEpochsHandlerMock) IsSetSenderInEeiOutputTransferFlagEnabled() bool {
	return false
}

// IsChangeDelegationOwnerFlagEnabled -
func (mock *EnableEpochsHandlerMock) IsChangeDelegationOwnerFlagEnabled() bool {
	return false
}

// IsRefactorPeersMiniBlocksFlagEnabled returns false
func (mock *EnableEpochsHandlerMock) IsRefactorPeersMiniBlocksFlagEnabled() bool {
	return mock.IsRefactorPeersMiniBlocksFlagEnabledField
}

// IsSCProcessorV2FlagEnabled -
func (mock *EnableEpochsHandlerMock) IsSCProcessorV2FlagEnabled() bool {
	return mock.IsSCProcessorV2FlagEnabledField
}

// IsFixAsyncCallBackArgsListFlagEnabled -
func (mock *EnableEpochsHandlerMock) IsFixAsyncCallBackArgsListFlagEnabled() bool {
	return false
}

// IsFixOldTokenLiquidityEnabled -
func (mock *EnableEpochsHandlerMock) IsFixOldTokenLiquidityEnabled() bool {
	return false
}

// IsRuntimeMemStoreLimitEnabled -
func (mock *EnableEpochsHandlerMock) IsRuntimeMemStoreLimitEnabled() bool {
	return false
}

// IsRuntimeCodeSizeFixEnabled -
func (mock *EnableEpochsHandlerMock) IsRuntimeCodeSizeFixEnabled() bool {
	return false
}

// IsMaxBlockchainHookCountersFlagEnabled -
func (mock *EnableEpochsHandlerMock) IsMaxBlockchainHookCountersFlagEnabled() bool {
	return false
}

// IsWipeSingleNFTLiquidityDecreaseEnabled -
func (mock *EnableEpochsHandlerMock) IsWipeSingleNFTLiquidityDecreaseEnabled() bool {
	return false
}

// IsAlwaysSaveTokenMetaDataEnabled -
func (mock *EnableEpochsHandlerMock) IsAlwaysSaveTokenMetaDataEnabled() bool {
	return false
}

// IsSetGuardianEnabled returns false
func (mock *EnableEpochsHandlerMock) IsSetGuardianEnabled() bool {
	return false
}

// IsRelayedNonceFixEnabled -
func (mock *EnableEpochsHandlerMock) IsRelayedNonceFixEnabled() bool {
	return false
}

// IsKeepExecOrderOnCreatedSCRsEnabled -
func (mock *EnableEpochsHandlerMock) IsKeepExecOrderOnCreatedSCRsEnabled() bool {
	return false
}

// IsMultiClaimOnDelegationEnabled -
func (mock *EnableEpochsHandlerMock) IsMultiClaimOnDelegationEnabled() bool {
	return false
}

// IsChangeUsernameEnabled -
func (mock *EnableEpochsHandlerMock) IsChangeUsernameEnabled() bool {
	return false
}

// IsConsistentTokensValuesLengthCheckEnabled -
func (mock *EnableEpochsHandlerMock) IsConsistentTokensValuesLengthCheckEnabled() bool {
	return false
}

// IsAutoBalanceDataTriesEnabled -
func (mock *EnableEpochsHandlerMock) IsAutoBalanceDataTriesEnabled() bool {
	return false
}

// FixDelegationChangeOwnerOnAccountEnabled -
func (mock *EnableEpochsHandlerMock) FixDelegationChangeOwnerOnAccountEnabled() bool {
	return false
}

// IsInterfaceNil returns true if there is no value under the interface
func (mock *EnableEpochsHandlerMock) IsInterfaceNil() bool {
	return mock == nil
}
