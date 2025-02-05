/*
Copyright 2019 The Vitess Authors.

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

// Command to generate a mock for this interface with mockgen.
//go:generate mockgen -source resharding_wrangler.go -destination mock_resharding_wrangler_test.go -package resharding -mock_names "Wrangler=MockReshardingWrangler"

package resharding

import (
	"time"

	"context"

	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
)

// Wrangler is the interface to be used in creating mock interface for wrangler, which is used for unit test. It includes a subset of the methods in go/vt/Wrangler.
type Wrangler interface {
	CopySchemaShardFromShard(ctx context.Context, tables, excludeTables []string, includeViews bool, sourceKeyspace, sourceShard, destKeyspace, destShard string, waitReplicasTimeout time.Duration, skipVerify bool) error

	WaitForFilteredReplication(ctx context.Context, keyspace, shard string, maxDelay time.Duration) error

	MigrateServedTypes(ctx context.Context, keyspace, shard string, cells []string, servedType topodatapb.TabletType, reverse, skipReFreshState bool, filteredReplicationWaitTime time.Duration, reverseReplication bool) error
}
