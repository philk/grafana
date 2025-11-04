package ualert

import "github.com/grafana/grafana/pkg/services/sqlstore/migrator"

// AddAlertRuleErrorEvalThreshold adds error_eval_threshold column to alert_rule and alert_rule_version tables.
func AddAlertRuleErrorEvalThreshold(mg *migrator.Migrator) {
	column := &migrator.Column{Name: "error_eval_threshold", Type: migrator.DB_SmallInt, Nullable: true}

	mg.AddMigration(
		"add error_eval_threshold column to alert_rule",
		migrator.NewAddColumnMigration(migrator.Table{Name: "alert_rule"}, column),
	)
	mg.AddMigration(
		"add error_eval_threshold column to alert_rule_version",
		migrator.NewAddColumnMigration(migrator.Table{Name: "alert_rule_version"}, column),
	)
}
