package schemas

// ConditionStruct struct
type ConditionStruct struct {
	Condition string `json:"condition"`
}

// ColumnsStruct struct
type ColumnsStruct struct {
	Columns string `json:"columns" binding:"required"`
}

// TableNameStruct struct
type TableNameStruct struct {
	TableName string `json:"table_name" binding:"required"`
}

// DbNameStruct struct
type DbNameStruct struct {
	DbName string `json:"db_name" binding:"required"`
}

// ValuesStruct struct
type ValuesStruct struct {
	Values []interface{} `json:"values"`
}

// JSONData struct
type JSONData struct {
	DbNameStruct
	TableNameStruct
	ConditionStruct
	ColumnsStruct
	ValuesStruct
}
