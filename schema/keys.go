package schema

// PrimaryKey represents a primary key in a database
type PrimaryKey struct {
	Columns []string
}

type IndexType int

func (t IndexType) String() string {
	switch t {
	case DefaultIndex:
		return "default"
	case GISTIndex:
		return "GIST"
	default:
		return "default"
	}
}

const DefaultIndex = 0
const GISTIndex = 1

// Index represents an index in a database
type Index struct {
	Columns []string
	Type    IndexType
}

// Unique represents a unique constraint in a database
type Unique struct {
	Columns []string
}

// ForeignKey represents a foreign key constraint in a database
type ForeignKey struct {
	LocalColumns   []string
	ForeignSchema  string
	ForeignTable   string
	ForeignColumns []string
}
