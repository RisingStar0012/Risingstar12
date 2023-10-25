package mysqlwip

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnSetDefaultForNotNullAdvisor)(nil)
	_ ast.Visitor     = (*columnSetDefaultForNotNullChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLColumnSetDefaultForNotNull, &ColumnSetDefaultForNotNullAdvisor{})
}

// ColumnSetDefaultForNotNullAdvisor is the advisor checking for set default value for not null column.
type ColumnSetDefaultForNotNullAdvisor struct {
}

// Check checks for set default value for not null column.
func (*ColumnSetDefaultForNotNullAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnSetDefaultForNotNullChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		(stmt).Accept(checker)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type columnSetDefaultForNotNullChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	text       string
}

// Enter implements the ast.Visitor interface.
func (checker *columnSetDefaultForNotNullChecker) Enter(in ast.Node) (ast.Node, bool) {
	var notNullColumnWithNoDefault []columnName
	switch node := in.(type) {
	// CREATE TABLE
	case *ast.CreateTableStmt:
		pkColumn := make(map[string]bool)
		for _, cons := range node.Constraints {
			if cons.Tp == ast.ConstraintPrimaryKey {
				for _, key := range cons.Keys {
					pkColumn[key.Column.Name.O] = true
				}
			}
		}

		for _, column := range node.Cols {
			_, ok := pkColumn[column.Name.Name.O]
			notNull := ok || !canNull(column)
			if notNull && !setDefault(column) && needDefault(column) {
				notNullColumnWithNoDefault = append(notNullColumnWithNoDefault, columnName{
					tableName:  node.Table.Name.O,
					columnName: column.Name.Name.O,
					line:       column.OriginTextPosition(),
				})
			}
		}
	// ALTER TABLE
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			// ADD COLUMNS
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					if !canNull(column) && !setDefault(column) && needDefault(column) {
						notNullColumnWithNoDefault = append(notNullColumnWithNoDefault, columnName{
							tableName:  node.Table.Name.O,
							columnName: column.Name.Name.O,
							line:       node.OriginTextPosition(),
						})
					}
				}
			// CHANGE COLUMN and MODIFY COLUMN
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				if !canNull(spec.NewColumns[0]) && !setDefault(spec.NewColumns[0]) && needDefault(spec.NewColumns[0]) {
					notNullColumnWithNoDefault = append(notNullColumnWithNoDefault, columnName{
						tableName:  node.Table.Name.O,
						columnName: spec.NewColumns[0].Name.Name.O,
						line:       node.OriginTextPosition(),
					})
				}
			}
		}
	}

	for _, column := range notNullColumnWithNoDefault {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.NotNullColumnWithNoDefault,
			Title:   checker.title,
			Content: fmt.Sprintf("Column `%s`.`%s` is NOT NULL but doesn't have DEFAULT", column.tableName, column.columnName),
			Line:    column.line,
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnSetDefaultForNotNullChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func setDefault(column *ast.ColumnDef) bool {
	for _, option := range column.Options {
		if option.Tp == ast.ColumnOptionDefaultValue {
			return true
		}
	}
	return false
}