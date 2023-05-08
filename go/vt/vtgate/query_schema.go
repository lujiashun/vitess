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

package vtgate

import (
	"context"
	"fmt"

	querypb "vitess.io/vitess/go/vt/proto/query"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	"vitess.io/vitess/go/vt/sqlparser"
)

func SchemaCheck(vc *vcursorImpl, ctx context.Context, ksName string, vschemaDDL *sqlparser.AlterVschema) error {
	if vschemaDDL.Action == sqlparser.AddAutoIncDDLAction {
		var tableName = vschemaDDL.AutoIncSpec.Sequence.Name.String()
		var qualifier = vschemaDDL.AutoIncSpec.Sequence.Qualifier.String()
		if qualifier == "" {
			qualifier = ksName
		}
		// Get Shard
		var ts = vc.topoServer
		shards, err := ts.GetShardNames(ctx, qualifier)
		if len(shards) == 0 {
			return fmt.Errorf("no shards in keyspace %v", qualifier)
		}
		// Get gateway
		var executor = vc.executor.(*Executor)
		var gw = executor.scatterConn.gateway
		// Get target
		target := &querypb.Target{
			Keyspace:   qualifier,
			Shard:      shards[0],
			TabletType: topodatapb.TabletType_PRIMARY,
			Cell:       executor.cell,
		}

		// Query
		var sql = "SHOW TABLES like '" + tableName + "'"
		ftRes, err := gw.Execute(ctx, target, sql, nil, 0, 0, nil)
		if err != nil {
			return err
		}

		if len(ftRes.Rows) == 0 {
			return fmt.Errorf("table %v does not exist in keyspace %v", tableName, qualifier)
		}
		return nil
	}
	return nil
}
