package operations

import (
	"fmt"
	"io"

	"github.com/sqlbunny/sqlschema/schema"
)

type CreateIndex struct {
	SchemaName string
	TableName  string
	IndexName  string
	Columns    []string
}

func (o CreateIndex) GetSQL() string {
	return fmt.Sprintf("CREATE INDEX CONCURRENTLY \"%s\" ON %s (%s)", o.IndexName, sqlName(o.SchemaName, o.TableName), columnList(o.Columns))
}

func (o CreateIndex) Dump(w io.Writer) {
	fmt.Fprint(w, "operations.CreateIndex {\n")
	fmt.Fprint(w, "SchemaName: "+esc(o.SchemaName)+",\n")
	fmt.Fprint(w, "TableName: "+esc(o.TableName)+",\n")
	fmt.Fprint(w, "IndexName: "+esc(o.IndexName)+",\n")
	fmt.Fprint(w, "Columns: []string{"+columnList(o.Columns)+"},\n")
	fmt.Fprint(w, "}")
}

func (o CreateIndex) Apply(d *schema.Database) error {
	s, ok := d.Schemas[o.SchemaName]
	if !ok {
		return fmt.Errorf("no such schema: %s", o.SchemaName)
	}
	t, ok := s.Tables[o.TableName]
	if !ok {
		return fmt.Errorf("no such table: %s", o.TableName)
	}
	if _, ok := t.Indexes[o.IndexName]; ok {
		return fmt.Errorf("index already exists: %s ", o.IndexName)
	}
	t.Indexes[o.IndexName] = &schema.Index{
		Columns: o.Columns,
	}
	return nil
}
