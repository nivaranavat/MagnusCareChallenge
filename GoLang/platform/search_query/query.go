package query

type Map map[string]interface{}

type Condition struct {
	Type   string   "json:conditions:type"
	Values []string "json:conditions:values"
}

type Pagination struct {
	From int "json:pagination:from"
	Size int "json:pagination:size"
}

type Query struct {
	Conditions []Condition "json:conditions"
	Pagination Pagination  "json:pagination"
}

func New() *Query {
	//create a new query object
	return &Query{
		Conditions: []Condition{},
		Pagination: Pagination{},
	}
}

func (q *Query) GetConditions() []Condition {
	//return all the conditions given
	return q.Conditions
}
func (q *Query) GetPagination() Pagination {
	//return the pagination given
	return q.Pagination
}
