package ocsf

type OccurrenceDetails struct {
	CellName           *string `json:"cell_name,omitempty" parquet:"cell_name,optional"`
	ColumnName         *string `json:"column_name,omitempty" parquet:"column_name,optional"`
	ColumnNumber       *int32  `json:"column_number,omitempty" parquet:"column_number,optional"`
	EndLine            *int32  `json:"end_line,omitempty" parquet:"end_line,optional"`
	JSONPath           *string `json:"json_path,omitempty" parquet:"json_path,optional"`
	PageNumber         *int32  `json:"page_number,omitempty" parquet:"page_number,optional"`
	RecordIndexInArray *int32  `json:"record_index_in_array,omitempty" parquet:"record_index_in_array,optional"`
	RowNumber          *int32  `json:"row_number,omitempty" parquet:"row_number,optional"`
	StartLine          *int32  `json:"start_line,omitempty" parquet:"start_line,optional"`
}
