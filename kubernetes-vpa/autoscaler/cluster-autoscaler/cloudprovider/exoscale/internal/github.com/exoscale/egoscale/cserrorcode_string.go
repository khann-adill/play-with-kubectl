/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package egoscale

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CloudRuntimeException-4250]
	_ = x[ExecutionException-4260]
	_ = x[HypervisorVersionChangedException-4265]
	_ = x[CloudException-4275]
	_ = x[AccountLimitException-4280]
	_ = x[AgentUnavailableException-4285]
	_ = x[CloudAuthenticationException-4290]
	_ = x[ConcurrentOperationException-4300]
	_ = x[ConflictingNetworkSettingsException-4305]
	_ = x[DiscoveredWithErrorException-4310]
	_ = x[HAStateException-4315]
	_ = x[InsufficientAddressCapacityException-4320]
	_ = x[InsufficientCapacityException-4325]
	_ = x[InsufficientNetworkCapacityException-4330]
	_ = x[InsufficientServerCapacityException-4335]
	_ = x[InsufficientStorageCapacityException-4340]
	_ = x[InternalErrorException-4345]
	_ = x[InvalidParameterValueException-4350]
	_ = x[ManagementServerException-4355]
	_ = x[NetworkRuleConflictException-4360]
	_ = x[PermissionDeniedException-4365]
	_ = x[ResourceAllocationException-4370]
	_ = x[ResourceInUseException-4375]
	_ = x[ResourceUnavailableException-4380]
	_ = x[StorageUnavailableException-4385]
	_ = x[UnsupportedServiceException-4390]
	_ = x[VirtualMachineMigrationException-4395]
	_ = x[AsyncCommandQueued-4540]
	_ = x[RequestLimitException-4545]
	_ = x[ServerAPIException-9999]
}

const _CSErrorCode_name = "CloudRuntimeExceptionExecutionExceptionHypervisorVersionChangedExceptionCloudExceptionAccountLimitExceptionAgentUnavailableExceptionCloudAuthenticationExceptionConcurrentOperationExceptionConflictingNetworkSettingsExceptionDiscoveredWithErrorExceptionHAStateExceptionInsufficientAddressCapacityExceptionInsufficientCapacityExceptionInsufficientNetworkCapacityExceptionInsufficientServerCapacityExceptionInsufficientStorageCapacityExceptionInternalErrorExceptionInvalidParameterValueExceptionManagementServerExceptionNetworkRuleConflictExceptionPermissionDeniedExceptionResourceAllocationExceptionResourceInUseExceptionResourceUnavailableExceptionStorageUnavailableExceptionUnsupportedServiceExceptionVirtualMachineMigrationExceptionAsyncCommandQueuedRequestLimitExceptionServerAPIException"

var _CSErrorCode_map = map[CSErrorCode]string{
	4250: _CSErrorCode_name[0:21],
	4260: _CSErrorCode_name[21:39],
	4265: _CSErrorCode_name[39:72],
	4275: _CSErrorCode_name[72:86],
	4280: _CSErrorCode_name[86:107],
	4285: _CSErrorCode_name[107:132],
	4290: _CSErrorCode_name[132:160],
	4300: _CSErrorCode_name[160:188],
	4305: _CSErrorCode_name[188:223],
	4310: _CSErrorCode_name[223:251],
	4315: _CSErrorCode_name[251:267],
	4320: _CSErrorCode_name[267:303],
	4325: _CSErrorCode_name[303:332],
	4330: _CSErrorCode_name[332:368],
	4335: _CSErrorCode_name[368:403],
	4340: _CSErrorCode_name[403:439],
	4345: _CSErrorCode_name[439:461],
	4350: _CSErrorCode_name[461:491],
	4355: _CSErrorCode_name[491:516],
	4360: _CSErrorCode_name[516:544],
	4365: _CSErrorCode_name[544:569],
	4370: _CSErrorCode_name[569:596],
	4375: _CSErrorCode_name[596:618],
	4380: _CSErrorCode_name[618:646],
	4385: _CSErrorCode_name[646:673],
	4390: _CSErrorCode_name[673:700],
	4395: _CSErrorCode_name[700:732],
	4540: _CSErrorCode_name[732:750],
	4545: _CSErrorCode_name[750:771],
	9999: _CSErrorCode_name[771:789],
}

func (i CSErrorCode) String() string {
	if str, ok := _CSErrorCode_map[i]; ok {
		return str
	}
	return "CSErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
}
