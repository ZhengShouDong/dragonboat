// Copyright 2017-2019 Lei Ni (nilei81@gmail.com) and other Dragonboat authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package settings

const (
	// EntryNonCmdFieldsSize defines the upper limit of the non-cmd field
	// length in pb.Entry.
	EntryNonCmdFieldsSize = 16 * 8
	// LargeEntitySize defines what is considered as a large entity for per node
	// entities.
	LargeEntitySize uint64 = 64 * 1024 * 1024
)

//
// Tuning configuration parameters here will impact the performance of your
// system. It will not corrupt your data. Only tune these parameters when
// you know what you are doing.
//
// To tune these parameters, place a json file named
// dragonboat-soft-settings.json in the current working directory of your
// dragonboat application, all fields in the json file will be applied to
// overwrite the default setting values. e.g. for a json file with the
// following content -
//
// {
//   "GetConnectedTimeoutSecond": 15,
// }
//
// soft.GetConnectedTimeoutSecond will be 15,
//
// The application need to be restarted to apply such configuration changes.
//

// Soft is the soft settings that can be changed after the deployment of a
// system.
var Soft = getSoftSettings()

type soft struct {
	//
	// Raft
	//

	// MaxEntrySize defines the max total entry size that can be included in
	// the Replicate message.
	MaxEntrySize uint64
	// InMemEntrySliceSize defines the maximum length of the in memory entry
	// slice.
	InMemEntrySliceSize uint64
	// MinEntrySliceFreeSize defines the minimum length of the free in memory
	// entry slice. A new entry slice of length InMemEntrySliceSize will be
	// allocated once the free entry size in the current slice is less than
	// MinEntrySliceFreeSize.
	MinEntrySliceFreeSize uint64
	// InMemGCTimeout defines how often dragonboat collects partial object.
	// It is defined in terms of number of ticks.
	InMemGCTimeout uint64

	//
	// Multiraft
	//

	// PendingProposalShards defines the number of shards for the pending
	// proposal data structure.
	PendingProposalShards uint64
	// SyncTaskInterval defines the interval in millisecond of periodic sync
	// state machine task.
	SyncTaskInterval uint64
	// IncomingReadIndexQueueLength defines the number of pending read index
	// requests allowed for each raft group.
	IncomingReadIndexQueueLength uint64
	// IncomingProposalQueueLength defines the number of pending proposals
	// allowed for each raft group.
	IncomingProposalQueueLength uint64
	// ReceiveQueueLength is the length of the receive queue on each node.
	ReceiveQueueLength uint64
	// SnapshotStatusPushDelayMS is the number of millisecond delays we impose
	// before pushing the snapshot results to raft node.
	SnapshotStatusPushDelayMS uint64
	// TaskQueueTargetLength defined the target length of each node's taskQ.
	// Dragonboat tries to make sure the queue is no longer than this target
	// length.
	TaskQueueTargetLength uint64
	// TaskQueueInitialCap defines the initial capcity of a task queue.
	TaskQueueInitialCap uint64
	// NodeHostSyncPoolSize defines the number of sync pools.
	NodeHostSyncPoolSize uint64
	// LatencySampleRatio defines the ratio how often latency is sampled.
	// It samples roughly every LatencySampleRatio ops.
	LatencySampleRatio uint64
	// LazyFreeCycle defines how often should entry queue and message queue
	// to be freed.
	LazyFreeCycle uint64
	// PanicOnSizeMismatch defines whether dragonboat should panic when snapshot
	// file size doesn't match the size recorded in snapshot metadata.
	PanicOnSizeMismatch bool

	//
	// RSM
	//
	BatchedEntryApply bool

	//
	// step engine
	//

	// TaskBatchSize defines the length of the committed batch slice.
	TaskBatchSize uint64
	// NodeReloadMillisecond defines how often step engine should reload
	// nodes, it is defined in number of millisecond.
	NodeReloadMillisecond uint64
	// StepEngineTaskWorkerCount is the number of workers to use to apply
	// proposals (processing committed proposals) to application state
	// machines.
	StepEngineTaskWorkerCount uint64
	// StepEngineSnapshotWorkerCount is the number of workers to take and
	// apply application state machine snapshots.
	StepEngineSnapshotWorkerCount uint64

	//
	// transport
	//

	// GetConnectedTimeoutSecond is the default timeout value in second when
	// trying to connect to a gRPC based server.
	GetConnectedTimeoutSecond uint64
	// MaxSnapshotConnections defines the max number of concurrent outgoing
	// snapshot connections.
	MaxSnapshotConnections uint64
	// MaxConcurrentStreamingSnapshot defines the max number of concurrent
	// incoming snapshot streams.
	MaxConcurrentStreamingSnapshot uint64
	// SendQueueLength is the length of the send queue used to hold messages
	// exchanged between nodehosts. You may need to increase this value when
	// you want to host large number nodes per nodehost.
	SendQueueLength uint64
	// StreamConnections defines how many connections to use for each remote
	// nodehost whene exchanging raft messages
	StreamConnections uint64
	// PerConnBufSize is the size of the per connection buffer used for
	// receiving incoming messages.
	PerConnectionSendBufSize uint64
	// PerConnectionRecvBufSize is the size of the recv buffer size.
	PerConnectionRecvBufSize uint64
	// SnapshotGCTick defines the number of ticks between two snapshot GC
	// operations.
	SnapshotGCTick uint64
	// SnapshotChunkTimeoutTick defines the max time allowed to receive
	// a snapshot.
	SnapshotChunkTimeoutTick uint64

	//
	// LogDB
	//
	KVTolerateCorruptedTailRecords bool
	// KVUseUniversalCompaction defines whether to use universal compaction to
	// reduce write amplification. This setting is default to false, change it to
	// true for existing system might cause unexpected consequences, please check
	// the documentation of your KV store for more details.
	//
	// KVUseUniversalCompaction support is experimental - it is not fully tested.
	KVUseUniversalCompaction bool
}

func getSoftSettings() soft {
	org := getDefaultSoftSettings()
	overwriteSoftSettings(&org)
	return org
}

func getDefaultSoftSettings() soft {
	return soft{
		MaxConcurrentStreamingSnapshot: 128,
		MaxSnapshotConnections:         64,
		SyncTaskInterval:               180000,
		PanicOnSizeMismatch:            true,
		LazyFreeCycle:                  1,
		LatencySampleRatio:             0,
		BatchedEntryApply:              true,
		GetConnectedTimeoutSecond:      5,
		MaxEntrySize:                   MaxMessageBatchSize,
		InMemGCTimeout:                 100,
		InMemEntrySliceSize:            512,
		MinEntrySliceFreeSize:          96,
		IncomingReadIndexQueueLength:   4096,
		IncomingProposalQueueLength:    2048,
		SnapshotStatusPushDelayMS:      1000,
		PendingProposalShards:          16,
		TaskQueueInitialCap:            64,
		TaskQueueTargetLength:          1024,
		NodeHostSyncPoolSize:           8,
		TaskBatchSize:                  512,
		NodeReloadMillisecond:          200,
		StepEngineTaskWorkerCount:      16,
		StepEngineSnapshotWorkerCount:  64,
		SendQueueLength:                1024 * 2,
		ReceiveQueueLength:             1024,
		StreamConnections:              4,
		PerConnectionSendBufSize:       LargeEntitySize,
		PerConnectionRecvBufSize:       2 * 1024 * 1024,
		SnapshotGCTick:                 30,
		SnapshotChunkTimeoutTick:       900,
		KVTolerateCorruptedTailRecords: true,
		KVUseUniversalCompaction:       false,
	}
}
