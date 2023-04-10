package dbrepository

type QueryByField struct {
	Field  string
	Values []interface{} // will run in batch mode

	Joins    []string // will run joins
	Preloads []string
	Limit    int
}
