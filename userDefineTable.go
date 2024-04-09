package Model

import (
	"errors"
	"fmt"
	"github.com/mohae/deepcopy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	"time"
)

type UserDefinedTable struct {
	ID              primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	TableName       string                   `json:"tableName" bson:"tableName"`
	Description     string                   `json:"description" bson:"description"`
	Overview        string                   `json:"overview" bson:"overview"`
	ParentTableName string                   `json:"parentTableName" bson:"parentTableName"`
	ChildTableName  string                   `json:"childTableName" bson:"childTableName"`
	DesktopJSON     string                   `json:"desktopJSON" bson:"desktopJSON"`
	MobileJSON      string                   `json:"mobileJSON" bson:"mobileJSON"`
	Series          TableSeries              `json:"series" bson:"series"`
	LayoutHtml      string                   `json:"layoutHtml" bson:"layoutHtml"`
	Design          string                   `json:"design" bson:"design"`
	DesignMobile    string                   `json:"designMobile" bson:"designMobile"`
	IsSummaryTable  bool                     `json:"isSummaryTable" bson:"isSummaryTable"`
	LinkTable       string                   `json:"linkTable" bson:"linkTable"`
	LinkCondition   string                   `json:"linkCondition" bson:"linkCondition"`
	EnableLog       bool                     `json:"enableLog" bson:"enableLog"`
	RecordLock      RecordLock               `json:"recordLock" bson:"recordLock"`
	CreatedDate     int64                    `json:"createdDate" bson:"createdDate"`
	CreatedBy       string                   `json:"createdBy" bson:"createdBy"`
	UpdatedDate     int64                    `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy       string                   `json:"updatedBy" bson:"updatedBy"`
	Buttons         []UserDefinedTableButton `json:"buttons" bson:"buttons"`
	Filters         []UserDefinedTableFilter `json:"filters" bson:"filters"`
	DataFlows       []DataFlow               `json:"dataFlows" bson:"dataFlows"`
	Fields          []UserDefinedField       `json:"fields" bson:"fields"`
}

type UserDefinedTableRequest struct {
	ID              primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	TableName       string                   `json:"tableName" bson:"tableName"`
	Description     string                   `json:"description" bson:"description"`
	Overview        string                   `json:"overview" bson:"overview"`
	ParentTableName string                   `json:"parentTableName" bson:"parentTableName"`
	ChildTableName  string                   `json:"childTableName" bson:"childTableName"`
	DesktopJSON     string                   `json:"desktopJSON" bson:"desktopJSON"`
	MobileJSON      string                   `json:"mobileJSON" bson:"mobileJSON"`
	Series          TableSeries              `json:"series" bson:"series"`
	LayoutHtml      string                   `json:"layoutHtml" bson:"layoutHtml"`
	Design          string                   `json:"design" bson:"design"`
	DesignMobile    string                   `json:"designMobile" bson:"designMobile"`
	IsSummaryTable  bool                     `json:"isSummaryTable" bson:"isSummaryTable"`
	LinkTable       string                   `json:"linkTable" bson:"linkTable"`
	LinkCondition   string                   `json:"linkCondition" bson:"linkCondition"`
	EnableLog       bool                     `json:"enableLog" bson:"enableLog"`
	RecordLock      RecordLock               `json:"recordLock" bson:"recordLock"`
	CreatedBy       string                   `json:"createdBy" bson:"createdBy"`
	UpdatedBy       string                   `json:"updatedBy" bson:"updatedBy"`
	Buttons         []UserDefinedTableButton `json:"buttons" bson:"buttons"`
	Filters         []UserDefinedTableFilter `json:"filters" bson:"filters"`
	DataFlows       []DataFlow               `json:"dataFlows" bson:"dataFlows"`
	Fields          []UserDefinedField       `json:"fields" bson:"fields"`
}

type UserDefinedTableUpdateRequest struct {
	Description     string                   `json:"description" bson:"description"`
	Overview        string                   `json:"overview" bson:"overview"`
	ParentTableName string                   `json:"parentTableName" bson:"parentTableName"`
	ChildTableName  string                   `json:"childTableName" bson:"childTableName"`
	DesktopJSON     string                   `json:"desktopJSON" bson:"desktopJSON"`
	MobileJSON      string                   `json:"mobileJSON" bson:"mobileJSON"`
	LayoutHtml      string                   `json:"layoutHtml" bson:"layoutHtml"`
	Design          string                   `json:"design" bson:"design"`
	DesignMobile    string                   `json:"designMobile" bson:"designMobile"`
	IsSummaryTable  bool                     `json:"isSummaryTable" bson:"isSummaryTable"`
	LinkTable       string                   `json:"linkTable" bson:"linkTable"`
	LinkCondition   string                   `json:"linkCondition" bson:"linkCondition"`
	EnableLog       bool                     `json:"enableLog" bson:"enableLog"`
	RecordLock      RecordLock               `json:"recordLock" bson:"recordLock"`
	Buttons         []UserDefinedTableButton `json:"buttons" bson:"buttons"`
	Filters         []UserDefinedTableFilter `json:"filters" bson:"filters"`
	DataFlows       []DataFlow               `json:"dataFlows" bson:"dataFlows"`
	Fields          []UserDefinedField       `json:"fields" bson:"fields"`
}

type UserDefinedField struct {
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	Overview        string             `json:"overview" bson:"overview"`
	GridColumnWidth int                `json:"gridColumnWidth" bson:"gridColumnWidth"`
	DataType        DataType           `json:"dataType" bson:"dataType"`
	Decimal         int                `json:"decimal" bson:"decimal"`
	ValidValue      []ValidValue       `json:"validValue" bson:"validValue"`
	DataLength      int                `json:"dataLength" bson:"dataLength"`
	Formula         string             `json:"formula" bson:"formula"`
	TableName       string             `json:"tableName" bson:"tableName"`
	ImagePath       string             `json:"imagePath" bson:"imagePath"`
	Filter          string             `json:"filter" bson:"filter"`
	DefaultValue    interface{}        `json:"defaultValue" bson:"defaultValue"`
	Sequence        int                `json:"sequence" bson:"sequence"`
	MinLength       int                `json:"minLength" bson:"minLength"`
	Design          string             `json:"design" bson:"design"`
	DesignMobile    string             `json:"designMobile" bson:"designMobile"`
	LinkedField     string             `json:"linkedField" bson:"linkedField"`
	RegEx           string             `json:"regEx" bson:"regEx"`
	FieldValueLock  string             `json:"fieldValueLock" bson:"fieldValueLock"`
	ShowInGrid      bool               `json:"showInGrid" bson:"showInGrid"`
	IsPrimary       bool               `json:"isPrimary" bson:"isPrimary"`
	IsRequired      bool               `json:"isRequired" bson:"isRequired"`
	IsArray         bool               `json:"isArray" bson:"isArray"`
	IsSearchable    bool               `json:"isSearchable" bson:"isSearchable"`
	IsFormula       bool               `json:"isFormula" bson:"isFormula"`
	IsMasked        bool               `json:"isMasked" bson:"isMasked"`
	IsEditable      bool               `json:"isEditable" bson:"isEditable"`
	IsSeries        bool               `json:"isSeries" bson:"isSeries"`
	IsReadOnly      bool               `json:"isReadOnly" bson:"isReadOnly"`
	BarcodeScan     bool               `json:"barcode_scan" bson:"barcode_scan"`
	IsUnique        bool               `json:"isUnique" bson:"isUnique"`
	IsSummary       bool               `json:"isSummary" bson:"isSummary"`
	IsMongoQuery    bool               `json:"isMongoQuery" bson:"isMongoQuery"`
	MongoQuery      *MongoQuery        `json:"mongoQuery" bson:"mongoQuery"`
	Fields          []UserDefinedField `json:"fields" bson:"fields"`
}

type MongoQuery struct {
	TableName string `json:"tableName"`
	Query     string `json:"query"`
}

// --------------------- Field ---------------------

type UserDefinedTableFilter struct {
	FieldName       string          `json:"fieldName" bson:"fieldName"`
	Operation       FilterOperation `json:"operation" bson:"operation"`
	ConfigurationID string          `json:"configurationID" bson:"configurationID"`
}

type FilterOperation string

const (
	FilterOperationEqual          FilterOperation = "EQUAL"
	FilterOperationNotEqual       FilterOperation = "NOT_EQUAL"
	FilterOperationLess           FilterOperation = "LESS"
	FilterOperationLessOrEqual    FilterOperation = "LESS_OR_EQUAL"
	FilterOperationGreater        FilterOperation = "GREATER"
	FilterOperationGreaterOrEqual FilterOperation = "GREATER_OR_EQUAL"
	FilterOperationContain        FilterOperation = "CONTAIN"
	FilterOperationNotContain     FilterOperation = "NOT_CONTAIN"
	FilterOperationStartWith      FilterOperation = "START_WITH"
	FilterOperationEndWith        FilterOperation = "END_WITH"
	FilterOperationIn             FilterOperation = "IN"
	FilterOperationNotIn          FilterOperation = "NOT_IN"
	FilterOperationAllIn          FilterOperation = "ALL_IN"
	FilterOperationNotAllIn       FilterOperation = "NOT_ALL_IN"
	FilterOperationAllOut         FilterOperation = "ALL_OUT"
	FilterOperationBlank          FilterOperation = "BLANK"
	FilterOperationNotBlank       FilterOperation = "NOT_BLANK"
)

func (udfFO FilterOperation) Mapping() string {
	return map[FilterOperation]string{
		FilterOperationEqual:          "$eq",
		FilterOperationNotEqual:       "$ne",
		FilterOperationLess:           "$lt",
		FilterOperationLessOrEqual:    "$lte",
		FilterOperationGreater:        "$gt",
		FilterOperationGreaterOrEqual: "$gte",
		FilterOperationContain:        "$regex",
		FilterOperationIn:             "$in",
		FilterOperationNotIn:          "$nin",
		FilterOperationAllIn:          "$all",
		FilterOperationNotAllIn:       "$all",
	}[udfFO]
}

func (udtF UserDefinedTableFilter) GenerateFilterBson(udf UserDefinedField, value string) (op bson.M, err error) {
	operation := udtF.Operation.Mapping()
	op = bson.M{operation: value}

	switch udf.DataType {
	//case DataTypeText:
	//	if udtF.Operation != FilterOperationEqual {
	//		return op, errors.New("Field's type is Text, can not make operation different than 'EQUAL' ")
	//	}
	//
	//	op["$eq"] = value
	case DataTypeNumber:
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return op, errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", value, err.Error()))
		}

		op[operation] = RoundNumberDecimalN(number, 2)
	}

	return
}

func (fc FieldCompare) MappingBsonM(udf UserDefinedField) (firstMatch bson.M, nextStages []bson.M, err error) {
	var value interface{}

	arrayOperationList := []FilterOperation{
		FilterOperationIn,
		FilterOperationNotIn,
		FilterOperationAllIn,
		FilterOperationNotAllIn,
		FilterOperationAllOut,
	}
	if IsItemExistsInArray(fc.Operation, arrayOperationList) {
		valueStrList := strings.Split(fc.Value, ";")

		switch udf.DataType {
		case DataTypeText:
			value = valueStrList
		case DataTypeNumber:
			numberList := make([]float64, 0)
			for _, v := range valueStrList {
				number, errF := strconv.ParseFloat(v, 64)
				if errF != nil {
					err = errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, errF.Error()))
					return
				}

				numberList = append(numberList, number)
			}

			value = numberList
		case DataTypeDate:
			numberList := make([]int64, 0)
			for _, v := range valueStrList {
				number, errF := strconv.ParseInt(v, 10, 64)
				if errF != nil {
					err = errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, errF.Error()))
					return
				}

				numberList = append(numberList, number)
			}

			value = numberList
		case DataTypeBoolean:
			booleanList := make([]bool, 0)
			for _, v := range valueStrList {
				b, errF := strconv.ParseBool(v)
				if errF != nil {
					err = errors.New(fmt.Sprintf("Parse [%s] to boolean ERROR - %s", fc.Value, errF.Error()))
					return
				}

				booleanList = append(booleanList, b)
			}

			value = booleanList
		}

	} else {
		switch udf.DataType {
		case DataTypeText:
			value = fc.Value
			//if fc.Operation == FilterOperationContain {
			//	op[operation] = primitive.Regex{Pattern: EscapeToRegex(fc.Value), Options: "i"}
			//}
		case DataTypeNumber:
			number, errF := strconv.ParseFloat(fc.Value, 64)
			if errF != nil {
				err = errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, errF.Error()))
				return
			}

			value = RoundNumberDecimalN(number, 2)
		case DataTypeDate:
			number, errF := strconv.ParseInt(fc.Value, 10, 64)
			if errF != nil {
				err = errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, errF.Error()))
				return
			}

			value = number
		case DataTypeBoolean:
			b, errF := strconv.ParseBool(fc.Value)
			if errF != nil {
				err = errors.New(fmt.Sprintf("Parse [%s] to boolean ERROR - %s", fc.Value, errF.Error()))
				return
			}

			if fc.Operation == FilterOperationEqual || fc.Operation == FilterOperationNotEqual {
				value = b
			}
		}
	}

	switch fc.Operation {
	case FilterOperationEqual:
		firstMatch = bson.M{fc.FieldName: bson.M{"$eq": value}}
	case FilterOperationNotEqual:
		firstMatch = bson.M{fc.FieldName: bson.M{"$ne": value}}
	case FilterOperationLess:
		firstMatch = bson.M{fc.FieldName: bson.M{"$lt": value}}
	case FilterOperationLessOrEqual:
		firstMatch = bson.M{fc.FieldName: bson.M{"$lte": value}}
	case FilterOperationGreater:
		firstMatch = bson.M{fc.FieldName: bson.M{"$gt": value}}
	case FilterOperationGreaterOrEqual:
		firstMatch = bson.M{fc.FieldName: bson.M{"$gte": value}}
	case FilterOperationContain:
		firstMatch = bson.M{fc.FieldName: bson.M{"$regex": primitive.Regex{Pattern: EscapeToRegex(fc.Value), Options: "i"}}}
	case FilterOperationNotContain:
		firstMatch = bson.M{fc.FieldName: bson.M{"$not": bson.M{"$regex": primitive.Regex{Pattern: EscapeToRegex(fc.Value), Options: "i"}}}}
	case FilterOperationStartWith:
		firstMatch = bson.M{fc.FieldName: bson.M{"$regex": primitive.Regex{Pattern: "^" + EscapeToRegex(fc.Value) + ".*", Options: "i"}}}
	case FilterOperationEndWith:
		firstMatch = bson.M{fc.FieldName: bson.M{"$regex": primitive.Regex{Pattern: EscapeToRegex(fc.Value) + ".*$", Options: "i"}}}
	case FilterOperationIn:
		firstMatch = bson.M{fc.FieldName: bson.M{"$in": value}}
	case FilterOperationNotIn:
		firstMatch = bson.M{fc.FieldName: bson.M{"$nin": value}}
	case FilterOperationAllIn:
		firstMatch = bson.M{fc.FieldName: bson.M{"$all": value}}
	case FilterOperationNotAllIn:
		firstMatch = bson.M{fc.FieldName: bson.M{"$not": bson.M{"$all": value}}}
	case FilterOperationAllOut:
		firstMatch = bson.M{
			"$or": bson.A{
				bson.M{fc.FieldName: bson.M{"$in": value}},
				bson.M{fc.FieldName: bson.M{"$exists": false}},
			},
		}

		nextStages = []bson.M{
			{
				"$addFields": bson.M{
					"unmatched": bson.M{
						"$setDifference": bson.A{
							fmt.Sprintf("$%s", fc.FieldName),
							value,
						},
					},
				},
			},
			{
				"$match": bson.M{
					"unmatched": bson.M{"$size": 0},
				},
			},
		}
	case FilterOperationBlank:
		firstMatch = bson.M{
			"$or": bson.A{
				bson.M{fc.FieldName: bson.M{"$exists": false}},
				bson.M{fc.FieldName: bson.M{"$in": []interface{}{nil, ""}}},
			},
		}
	case FilterOperationNotBlank:
		firstMatch = bson.M{
			"$nor": bson.A{
				bson.M{fc.FieldName: bson.M{"$exists": false}},
				bson.M{fc.FieldName: bson.M{"$in": []interface{}{nil, ""}}},
			},
		}
	}

	return
}

func (fc FieldCompare) GenerateFilterBson(udf UserDefinedField) (op bson.M, err error) {
	operation := fc.Operation.Mapping()
	op = bson.M{operation: fc.Value}

	arrayOperationList := []FilterOperation{FilterOperationIn, FilterOperationNotIn, FilterOperationAllIn, FilterOperationNotAllIn}
	if IsItemExistsInArray(fc.Operation, arrayOperationList) {
		valueStrList := strings.Split(fc.Value, ";")

		switch udf.DataType {
		case DataTypeText:
			op[operation] = valueStrList
		case DataTypeNumber:
			numberList := make([]float64, 0)
			for _, v := range valueStrList {
				number, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return op, errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, err.Error()))
				}

				numberList = append(numberList, number)
			}

			op[operation] = numberList
		case DataTypeDate:
			numberList := make([]int64, 0)
			for _, v := range valueStrList {
				number, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return op, errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, err.Error()))
				}

				numberList = append(numberList, number)
			}

			op[operation] = numberList
		case DataTypeBoolean:
			booleanList := make([]bool, 0)
			for _, v := range valueStrList {
				b, err := strconv.ParseBool(v)
				if err != nil {
					return op, errors.New(fmt.Sprintf("Parse [%s] to boolean ERROR - %s", fc.Value, err.Error()))
				}

				booleanList = append(booleanList, b)
			}

			op[operation] = booleanList
		}

		if fc.Operation == FilterOperationNotAllIn {
			tmp := deepcopy.Copy(op).(bson.M)
			op = make(bson.M)
			op["$not"] = tmp
		}

		return
	}

	switch udf.DataType {
	case DataTypeText:
		if fc.Operation == FilterOperationContain {
			op[operation] = primitive.Regex{Pattern: EscapeToRegex(fc.Value), Options: "i"}
		}
	case DataTypeNumber:
		number, err := strconv.ParseFloat(fc.Value, 64)
		if err != nil {
			return op, errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, err.Error()))
		}

		op[operation] = RoundNumberDecimalN(number, 2)
	case DataTypeDate:
		number, err := strconv.ParseInt(fc.Value, 10, 64)
		if err != nil {
			return op, errors.New(fmt.Sprintf("Parse [%s] to number ERROR - %s", fc.Value, err.Error()))
		}

		op[operation] = number
	case DataTypeBoolean:
		b, err := strconv.ParseBool(fc.Value)
		if err != nil {
			return op, errors.New(fmt.Sprintf("Parse [%s] to boolean ERROR - %s", fc.Value, err.Error()))
		}

		if fc.Operation == FilterOperationEqual || fc.Operation == FilterOperationNotEqual {
			op[operation] = b
		}
	}

	return
}

type FieldCompare struct {
	FieldName string          `json:"fieldName"`
	Value     string          `json:"value"`
	Operation FilterOperation `json:"operation"`
	Logical   string          `json:"logical"`
}

type TableSeries struct {
	Prefix          string `json:"prefix" bson:"prefix"`
	StartNumber     int    `json:"startNumber" bson:"startNumber"`
	Length          int    `json:"length" bson:"length"`
	CurrentIdentity int64  `json:"currentIdentity,omitempty" bson:"currentIdentity,omitempty"`
}

type DataFlow struct {
	TargetTableName string                  `json:"targetTableName" bson:"targetTableName"`
	TargetField     string                  `json:"targetField" bson:"targetField"`
	Method          DataFlowMethod          `json:"method" bson:"method"`
	MappingFields   []DataFlowMappingField  `json:"mappingFields" bson:"mappingFields"` // Field to find the parent document
	SourceFormula   string                  `json:"sourceFormula" bson:"sourceFormula"`
	UpdateOperation DataFlowUpdateOperation `json:"updateOperation" bson:"updateOperation"`
	//MappingArrays   []DataFlowMappingArray  `json:"mappingArrays" bson:"mappingArrays"` // Mapping Array to update nested array's element
}

type DataFlowMappingField struct {
	SourceField string `json:"sourceField" bson:"sourceField"`
	TargetField string `json:"targetField" bson:"targetField"`
}

type DataFlowMappingArray struct {
	ArrayName     string                       `json:"arrayName" bson:"arrayName"`
	MappingFields []DataFlowArrayMappingFields `json:"mappingFields" bson:"mappingFields"`
	UpdateFields  DataFlowMappingField         `json:"updateFields" bson:"updateFields"`
}

type DataFlowArrayMappingFields struct {
	SourceArray DataFlowMappingArray
	TargetField string
}
type DataFlowMethod string

const (
	DataFlowMethodUpdate DataFlowMethod = "UPDATE"
	DataFlowMethodAdd    DataFlowMethod = "ADD"
)

type DataFlowUpdateOperation string

const (
	DataFlowUpdateOperationAdd      DataFlowUpdateOperation = "ADD"
	DataFlowUpdateOperationSubtract DataFlowUpdateOperation = "SUBTRACT"
	DataFlowUpdateOperationMultiply DataFlowUpdateOperation = "MULTIPLY"
	DataFlowUpdateOperationDivide   DataFlowUpdateOperation = "DIVIDE"
)

type ValidValue struct {
	ID          interface{} `json:"id" bson:"id" `
	Description string      `json:"description" bson:"description" `
}

type DataType string

type RecordLock struct {
	FieldName  string `json:"fieldName" bson:"fieldName"`
	FieldValue string `json:"fieldValue" bson:"fieldValue"`
}

type UserDefinedTableButton struct {
	ButtonName string `json:"buttonName" bson:"buttonName"`
	ApiURL     string `json:"apiURL" bson:"apiURL"`
}

const (
	DataTypeText    DataType = "text"
	DataTypeImage   DataType = "image"
	DataTypeNumber  DataType = "number"
	DataTypeBoolean DataType = "boolean"
	DataTypeDate    DataType = "date"
	DataTypeTime    DataType = "time"
	DataTypeObject  DataType = "object"
	DataTypeFile    DataType = "file"
	DataTypeHTML    DataType = "html"
)

func (ft *DataType) isValid() bool {
	*ft = DataType(strings.ToLower(strings.ReplaceAll(strings.Trim(string(*ft), " "), " ", "_")))
	switch string(*ft) {
	case "image":
		*ft = DataTypeImage
	case "string", "text":
		*ft = DataTypeText
	case "int", "int32", "int64", "number", "float32", "float64", "double":
		*ft = DataTypeNumber
	case "date":
		*ft = DataTypeDate
	case "time":
		*ft = DataTypeTime
	case "object":
		*ft = DataTypeObject
	case "boolean":
		*ft = DataTypeBoolean
	case "html":
		*ft = DataTypeHTML
	case "file":
		*ft = DataTypeFile
	default:
		return false
	}

	return true
}

func (ft *DataType) toCorrectCase() {
	*ft = DataType(strings.ToLower(strings.ReplaceAll(strings.Trim(string(*ft), " "), " ", "_")))
}

func (tableReq *UserDefinedTableRequest) Validate() (err error) {
	if len(tableReq.Fields) == 0 {
		err = errors.New("Invalid table - Empty fields! ")
	}

	numberOfPrimary := 0
	countIsSeries := 0

	for _, field := range tableReq.Fields {
		if field.IsPrimary {
			numberOfPrimary++
		}

		if field.IsSeries {
			countIsSeries++
		}

		//Check isSeries and Data type text
		if field.IsSeries && field.DataType != DataTypeText {
			return errors.New(fmt.Sprintf("Field [%s] must have data type [text] when having series property ! ", field.Name))
		}
	}

	if countIsSeries > 1 {
		return errors.New("Table only allow to have 1 field has document series! ")
	}

	if numberOfPrimary == 0 {
		err = errors.New("Table missing primary field! ")
	} else if numberOfPrimary > 1 {
		err = errors.New("Table only allow to have 1 primary field! ")
	}

	for i := 0; i < len(tableReq.Fields); i++ {
		if err := tableReq.Fields[i].Validate(); err != nil {
			return errors.New(fmt.Sprintf("Validate field [%s] ERROR - %s", tableReq.Fields[i].Name, err.Error()))
		}
	}

	return
}

func (table UserDefinedTable) GetPrimaryField() (f UserDefinedField) {
	for _, field := range table.Fields {
		if field.IsPrimary {
			return field
		}
	}

	return
}

func (table UserDefinedTable) GetSeriesField() (f UserDefinedField) {
	for _, field := range table.Fields {
		if field.IsSeries {
			return field
		}
	}

	return
}

func (table UserDefinedTable) CheckDefaultHeaderField(fieldName string) bool {
	for _, field := range table.Fields {
		if StringToCompareString(field.Name) == StringToCompareString(fieldName) {
			return true
		}
	}

	return false
}

func (table *UserDefinedTable) AddDefaultHeaderFields() {
	defaultFields := []UserDefinedField{
		{
			Name:            "createdBy",
			Description:     "Created By",
			Overview:        "",
			ShowInGrid:      false,
			GridColumnWidth: 0,
			DataType:        DataTypeText,
			DataLength:      0,
			IsPrimary:       false,
			IsRequired:      false,
			IsArray:         false,
			IsSearchable:    true,
			IsFormula:       false,
			Formula:         "",
			Fields:          nil,
		},
		{
			Name:            "createdDate",
			Description:     "Created Date ( Unix )",
			Overview:        "",
			ShowInGrid:      false,
			GridColumnWidth: 0,
			DataType:        DataTypeDate,
			DataLength:      0,
			IsPrimary:       false,
			IsRequired:      false,
			IsArray:         false,
			IsSearchable:    true,
			IsFormula:       false,
			Formula:         "",
			Fields:          nil,
		},
		{
			Name:            "updatedDate",
			Description:     "Updated Date ( Unix )",
			Overview:        "",
			ShowInGrid:      false,
			GridColumnWidth: 0,
			DataType:        DataTypeDate,
			DataLength:      0,
			IsPrimary:       false,
			IsRequired:      false,
			IsArray:         false,
			IsSearchable:    true,
			IsFormula:       false,
			Formula:         "",
			Fields:          nil,
		},
		{
			Name:            "updatedBy",
			Description:     "Updated By",
			Overview:        "",
			ShowInGrid:      false,
			GridColumnWidth: 0,
			DataType:        DataTypeText,
			DataLength:      0,
			IsPrimary:       false,
			IsRequired:      false,
			IsArray:         false,
			IsSearchable:    true,
			IsFormula:       false,
			Formula:         "",
			Fields:          nil,
		},
		{
			Name:         "summaryFlag",
			Description:  "Summary Flag",
			DataType:     DataTypeBoolean,
			IsSearchable: true,
		},
	}

	for _, df := range defaultFields {
		if !checkFieldExists(table.Fields, df.Name) {
			table.Fields = append(table.Fields, df)
		}
	}

	return
}

func (field *UserDefinedField) Validate() (err error) {
	if field.DataType == DataTypeNumber && field.IsPrimary {
		field.Decimal = 0
	}

	// Validate Field data type
	if !field.DataType.isValid() {
		return errors.New(fmt.Sprintf("Field [%s] Invalid Type [%s]", field.Name, field.DataType))
	}

	if len(field.ValidValue) > 0 {
		if field.DataType == DataTypeBoolean {
			field.ValidValue = nil
		}
	}

	//field.DefaultValue, err = field.ValidTypeValue(field.DefaultValue)
	//if err != nil {
	//	return errors.New("Valid Default value ERROR - " + err.Error())
	//}

	for i, value := range field.ValidValue {
		retVal, err := field.ValidTypeValue(value.ID)
		if err != nil {
			return errors.New("Valid type value ERROR - " + err.Error())
		}

		field.ValidValue[i].ID = retVal
	}

	for i := 0; i < len(field.Fields); i++ {
		if err := field.Fields[i].Validate(); err != nil {
			return errors.New(fmt.Sprintf("Validate field [%s] ERROR - %s", field.Fields[i].Name, err.Error()))
		}
	}

	return nil
}

func (field UserDefinedField) ValidTypeValue(value interface{}) (returnVal interface{}, err error) {
	returnVal = value
	ft := field.DataType

	// Skip valid object field
	if ft == DataTypeObject {
		return
	}

	// This convert unknown float64 to int64 if type = number
	switch i := value.(type) {
	case float64:
		if ft == DataTypeDate {
			if i-float64(int64(i)) > 0 {
				return nil, errors.New(fmt.Sprintf("Invalid value [%f] for [%s] type", i, ft))
			}
			returnVal = int64(i)

		} else if ft == DataTypeNumber {
			returnVal = RoundNumberDecimalN(i, field.Decimal)
		}
	}

	switch v := returnVal.(type) {
	case string:
		if ft != DataTypeText && ft != DataTypeTime && ft != DataTypeImage && ft != DataTypeFile && ft != DataTypeHTML {
			return nil, errors.New(fmt.Sprintf("Invalid value [%s] for [%s] type", v, ft))
		} else if ft == DataTypeTime {
			v = strings.ReplaceAll(v, " ", "")
			if _, err := time.Parse("15:04", v); err != nil {
				return nil, errors.New(fmt.Sprintf("Invalid value [%s] for [%s] type", v, ft))
			}

			returnVal = v
		}
	case int64:
		if ft != DataTypeDate {
			return nil, errors.New(fmt.Sprintf("Invalid value [%d] for [%s] type", v, ft))
		}
	case float64:
		if ft != DataTypeNumber {
			return nil, errors.New(fmt.Sprintf("Invalid value [%f] for [%s] type", v, ft))
		} else if field.IsPrimary {
			returnVal = int64(v)
		}
	case map[string]interface{}:
		if ft != DataTypeObject {
			return nil, errors.New(fmt.Sprintf("Invalid value [%v] for [%s] type", v, ft))
		}
	case bool:
		if ft != DataTypeBoolean {
			return nil, errors.New(fmt.Sprintf("Invalid value [%v] for [%s] type", v, ft))
		}
		//case nil:
		//	if ft == DataTypeBoolean {
		//		if field.DefaultValue != nil {
		//			returnVal = field.DefaultValue
		//		} else {
		//			returnVal = false
		//		}
		//	}
	}

	if len(field.ValidValue) > 0 && field.TableName == "" {
		isValidValue := false
		for _, validValue := range field.ValidValue {
			if validValue.ID == returnVal {
				isValidValue = true
				break
			}
		}

		if !isValidValue {
			return nil, errors.New("Value not in valid value list ")
		}
	}

	return
}

func checkFieldExists(fields []UserDefinedField, fieldName string) bool {
	for _, field := range fields {
		if StringToCompareString(field.Name) == StringToCompareString(fieldName) {
			return true
		}
	}

	return false
}

type FieldLogical struct {
	Logical    string         `json:"logical"`
	Operations []FieldCompare `json:"operations"`
	Logicals   []FieldLogical `json:"logicals"`
}
