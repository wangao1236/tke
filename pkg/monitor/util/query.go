package util

import (
	"tkestack.io/tke/api/monitor"
	"tkestack.io/tke/pkg/monitor/storage/types"
)

// ParseClusterConditions determines whether all clusters need to be queried
func ParseClusterConditions(storageType string, condition *monitor.MetricQueryCondition) bool {
	ifQueryAll := false
	if condition.Expr == "=~" && condition.Value == ".*" {
		ifQueryAll = true
		if storageType == types.EsType {
			condition.Expr = "!="
			condition.Value = ""
		}
	}
	return ifQueryAll
}
