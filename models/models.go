package models

/*Domain estructura con anotaciones para gorp y json*/
type Domain struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

/*Rule estructura con anotaciones para gorp y json*/
type Rule struct {
	Id          int64  `db:"id" json:"id"`
	Domain      int64  `db:"domain" json:"domain"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

/*Component estructura con anotaciones para gorp y json*/
type Component struct {
	Id         			int64  	`db:"id" json:"id"`
	Rule       			int64  	`db:"rule" json:"rule"`
	Comparator 			int64  	`db:"comparator" json:"comparator"`
	Path       			string 	`db:"path" json:"path"`
	Value      			string 	`db:"value" json:"value"`
	Literal    			int16  	`db:"literal" json:"literal"`
}

/*ComponentExtended estructura con anotaciones para gorp y json*/
type ComponentExtended struct {
	Id         			int64  	`db:"id" json:"id"`
	Rule       			int64  	`db:"rule" json:"rule"`
	RuleName       	string  `db:"ruleName" json:"ruleName"`
	Comparator 			int64  	`db:"comparator" json:"comparator"`
	ComparatorName 	string  `db:"comparatorName" json:"comparatorName"`
	Path       			string 	`db:"path" json:"path"`
	Value      			string 	`db:"value" json:"value"`
	Literal    			int16  	`db:"literal" json:"literal"`
}

/*ComponentByRule estructura con anotaciones para gorp y json*/
type ComponentByRule struct {
	Comparator string `db:"comparator" json:"Comparator"`
	Path       string `db:"path" json:"path"`
	Value      string `db:"value" json:"value"`
	Literal    int16  `db:"literal" json:"literal"`
}

/*ParamsByRule estructura con anotaciones para gorp y json*/
type ParamsByRule struct {
	Comparator string `db:"comparator" json:"Comparator"`
	Path       string `db:"path" json:"path"`
	Value      string `db:"value" json:"value"`
	Literal    int16  `db:"literal" json:"literal"`
}
