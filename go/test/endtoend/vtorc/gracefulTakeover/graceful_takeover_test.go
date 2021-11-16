/*
Copyright 2021 The Vitess Authors.

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

package gracefulTakeover

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"vitess.io/vitess/go/test/endtoend/cluster"
	"vitess.io/vitess/go/test/endtoend/vtorc/utils"
)

// make an api call to graceful primary takeover and let vtorc fix it
// covers the test case graceful-master-takeover from orchestrator
func TestGracefulPrimaryTakeover(t *testing.T) {
	defer cluster.PanicHandler(t)
	utils.SetupVttabletsAndVtorc(t, clusterInfo, 2, 0, nil, "test_config.json")
	keyspace := &clusterInfo.ClusterInstance.Keyspaces[0]
	shard0 := &keyspace.Shards[0]

	// find primary from topo
	curPrimary := utils.ShardPrimaryTablet(t, clusterInfo, keyspace, shard0)
	assert.NotNil(t, curPrimary, "should have elected a primary")

	// find the replica tablet
	var replica *cluster.Vttablet
	for _, tablet := range shard0.Vttablets {
		// we know we have only two tablets, so the one not the primary must be the replica
		if tablet.Alias != curPrimary.Alias {
			replica = tablet
		}
	}
	assert.NotNil(t, replica, "could not find replica tablet")

	// check that the replication is setup correctly before we failover
	utils.CheckReplication(t, clusterInfo, curPrimary, []*cluster.Vttablet{replica}, 10*time.Second)

	status, _ := utils.MakeAPICall(t, fmt.Sprintf("http://localhost:3000/api/graceful-primary-takeover/localhost/%d/localhost/%d", curPrimary.MySQLPort, replica.MySQLPort))
	assert.Equal(t, 200, status)

	// check that the replica gets promoted
	utils.CheckPrimaryTablet(t, clusterInfo, replica, true)
	utils.VerifyWritesSucceed(t, clusterInfo, replica, []*cluster.Vttablet{curPrimary}, 10*time.Second)
}

// make an api call to graceful primary takeover without specifying the primary tablet to promote
// covers the test case graceful-master-takeover-fail-no-target from orchestrator
func TestGracefulPrimaryTakeoverFailNoTarget(t *testing.T) {
	defer cluster.PanicHandler(t)
	utils.SetupVttabletsAndVtorc(t, clusterInfo, 3, 0, nil, "test_config.json")
	keyspace := &clusterInfo.ClusterInstance.Keyspaces[0]
	shard0 := &keyspace.Shards[0]

	// find primary from topo
	curPrimary := utils.ShardPrimaryTablet(t, clusterInfo, keyspace, shard0)
	assert.NotNil(t, curPrimary, "should have elected a primary")

	// find the replica tablet
	var replicas []*cluster.Vttablet
	for _, tablet := range shard0.Vttablets {
		// we know we have only two tablets, so the one not the primary must be the replica
		if tablet.Alias != curPrimary.Alias {
			replicas = append(replicas, tablet)
		}
	}
	assert.Equal(t, 2, len(replicas), "could not find replica tablets")

	// check that the replication is setup correctly before we failover
	utils.CheckReplication(t, clusterInfo, curPrimary, replicas, 10*time.Second)

	status, response := utils.MakeAPICall(t, fmt.Sprintf("http://localhost:3000/api/graceful-primary-takeover/localhost/%d/", curPrimary.MySQLPort))
	assert.Equal(t, 500, status)
	assert.Contains(t, response, "GracefulPrimaryTakeover: target instance not indicated")

	// check that the replica doesn't get promoted and the previous primary is still the primary
	utils.CheckPrimaryTablet(t, clusterInfo, curPrimary, true)
	utils.VerifyWritesSucceed(t, clusterInfo, curPrimary, replicas, 10*time.Second)
}

// make an api call to graceful primary takeover auto and let vtorc fix it
// covers the test case graceful-master-takeover-auto from orchestrator
func TestGracefulPrimaryTakeoverAuto(t *testing.T) {
	defer cluster.PanicHandler(t)
	utils.SetupVttabletsAndVtorc(t, clusterInfo, 2, 1, nil, "test_config.json")
	keyspace := &clusterInfo.ClusterInstance.Keyspaces[0]
	shard0 := &keyspace.Shards[0]

	// find primary from topo
	primary := utils.ShardPrimaryTablet(t, clusterInfo, keyspace, shard0)
	assert.NotNil(t, primary, "should have elected a primary")

	// find the replica tablet and the rdonly tablet
	var replica, rdonly *cluster.Vttablet
	for _, tablet := range shard0.Vttablets {
		// we know we have only two replcia tablets, so the one not the primary must be the other replica
		if tablet.Alias != primary.Alias && tablet.Type == "replica" {
			replica = tablet
		}
		if tablet.Type == "rdonly" {
			rdonly = tablet
		}
	}
	assert.NotNil(t, replica, "could not find replica tablet")
	assert.NotNil(t, rdonly, "could not find rdonly tablet")

	// check that the replication is setup correctly before we failover
	utils.CheckReplication(t, clusterInfo, primary, []*cluster.Vttablet{replica, rdonly}, 10*time.Second)

	status, _ := utils.MakeAPICall(t, fmt.Sprintf("http://localhost:3000/api/graceful-primary-takeover-auto/localhost/%d/localhost/%d", primary.MySQLPort, replica.MySQLPort))
	assert.Equal(t, 200, status)

	// check that the replica gets promoted
	utils.CheckPrimaryTablet(t, clusterInfo, replica, true)
	utils.VerifyWritesSucceed(t, clusterInfo, replica, []*cluster.Vttablet{primary, rdonly}, 10*time.Second)

	status, _ = utils.MakeAPICall(t, fmt.Sprintf("http://localhost:3000/api/graceful-primary-takeover-auto/localhost/%d/", replica.MySQLPort))
	assert.Equal(t, 200, status)

	// check that the primary gets promoted back
	utils.CheckPrimaryTablet(t, clusterInfo, primary, true)
	utils.VerifyWritesSucceed(t, clusterInfo, primary, []*cluster.Vttablet{replica, rdonly}, 10*time.Second)
}
