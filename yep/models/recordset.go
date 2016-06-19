// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
recordStruct implements RecordSet
*/
type RecordSet struct {
	query     Query
	mi        *modelInfo
	env       *Environment
	ids       []int64
	callStack []*methodLayer
}

func (rs RecordSet) String() string {
	idsStr := make([]string, len(rs.ids))
	for i, id := range rs.ids {
		idsStr[i] = strconv.Itoa(int(id))
		i++
	}
	rsIds := strings.Join(idsStr, ",")
	return fmt.Sprintf("%s(%s)", rs.mi.name, rsIds)
}

/*
Env returns the RecordSet's Environment
*/
func (rs RecordSet) Env() *Environment {
	return rs.env
}

/*
ModelName returns the model name of the RecordSet
*/
func (rs RecordSet) ModelName() string {
	return rs.mi.name
}

/*
Ids return the ids of the RecordSet
*/
func (rs RecordSet) Ids() []int64 {
	return rs.ids
}

/*
Search query the database with the current filter and fills the RecordSet with the queries ids.
Does nothing in case RecordSet already has Ids. It panics in case of error.
It returns a pointer to the same RecordSet.
*/
func (rs *RecordSet) Search() *RecordSet {
	if len(rs.Ids()) == 0 {
		return rs.ForceSearch()
	}
	return rs
}

/*
Search query the database with the current filter and fills the RecordSet with the queries ids.
Overwrite RecordSet Ids if any. It panics in case of error.
It returns a pointer to the same RecordSet.
*/
func (rs *RecordSet) ForceSearch() *RecordSet {
	//var idParams []interface{}
	//num := rs.ValuesFlat(&idParams, "ID")
	//ids := make([]int64, num)
	//for i := 0; i < int(num); i++ {
	//	ids[i] = idParams[i].(int64)
	//}
	//return copyRecordStruct(rs).withIds(ids)
	return rs
}

/*
Write updates the database with the given data and returns the number of updated rows.
It panics in case of error.
*/
func (rs RecordSet) Write(data FieldMap) bool {
	//_, err := rs.qs.Update(data)
	//if err != nil {
	//	panic(fmt.Errorf("recordSet `%s` Write error: %s", rs, err))
	//}
	//rs.updateStoredFields(data)
	return true
}

/*
Unlink deletes the database record of this RecordSet and returns the number of deleted rows.
*/
func (rs RecordSet) Unlink() int64 {
	sql, args := rs.query.deleteQuery()
	res := DBExecute(rs.env.cr, sql, args)
	num, _ := res.RowsAffected()
	return num
}

/*
Filter returns a new RecordSet with the given additional filter condition.
*/
func (rs RecordSet) Filter(cond, op string, data ...interface{}) *RecordSet {
	rs.query.cond = rs.query.cond.And(cond, op, data...)
	return &rs
}

/*
Exclude returns a new RecordSet with the given additional NOT filter condition.
*/
func (rs RecordSet) Exclude(cond, op string, data ...interface{}) *RecordSet {
	rs.query.cond = rs.query.cond.AndNot(cond, op, data...)
	return &rs
}

/*
SetCond returns a new RecordSet with the given additional condition
*/
func (rs RecordSet) SetCond(cond *Condition) *RecordSet {
	rs.query.cond = rs.query.cond.AndCond(cond)
	return &rs
}

/*
Limit returns a new RecordSet with the given limit as additional condition
*/
func (rs RecordSet) Limit(limit int, args ...int) *RecordSet {
	rs.query.limit = limit
	if len(args) > 0 {
		rs.query.offset = args[0]
	}
	return &rs
}

/*
Offset returns a new RecordSet with the given offset as additional condition
*/
func (rs RecordSet) Offset(offset int) *RecordSet {
	rs.query.offset = offset
	return &rs
}

/*
OrderBy returns a new RecordSet with the given ORDER BY clause in its Query
*/
func (rs RecordSet) OrderBy(exprs ...string) *RecordSet {
	rs.query.orders = append(rs.query.orders, exprs...)
	return &rs
}

/*
GroupBy returns a new RecordSet with the given GROUP BY clause in its Query
*/
func (rs RecordSet) GroupBy(exprs ...string) *RecordSet {
	rs.query.groups = append(rs.query.groups, exprs...)
	return &rs
}

// Distinct returns a new RecordSet with its Query filtering duplicates
func (rs RecordSet) Distinct() *RecordSet {
	rs.query.distinct = true
	return &rs
}

/*
SearchCount fetch from the database the number of records that match the RecordSet conditions
It panics in case of error
*/
func (rs RecordSet) SearchCount() int {
	sql, args := rs.query.countQuery()
	var res int
	DBGet(rs.env.cr, res, sql, args)
	return res
}

/*
All query all data pointed by the RecordSet and map to containers.
It panics in case of error
*/
func (rs RecordSet) ReadAll(container interface{}, cols ...string) int64 {
	//if err := checkStructPtr(container, true); err != nil {
	//	panic(fmt.Errorf("recordSet `%s` ReadAll() error: %s", rs, err))
	//}
	//num, err := rs.OrderBy("ID").All(container, cols...)
	//if err != nil {
	//	panic(fmt.Errorf("recordSet `%s` ReadAll() error: %s", rs, err))
	//}
	//val := reflect.ValueOf(container)
	//ind := reflect.Indirect(val)
	//if ind.Kind() == reflect.Slice {
	//	contSlice := make([]interface{}, ind.Len())
	//	for i := 0; i < ind.Len(); i++ {
	//		csIndex := reflect.ValueOf(contSlice).Index(i)
	//		csIndex.Set(ind.Index(i))
	//	}
	//	rs = rs.Search()
	//	for i, item := range rs.Records() {
	//		item.computeFields(contSlice[i])
	//	}
	//	return num
	//}
	//rs.computeFields(container)
	return 1
}

/*
One query the RecordSet row and map to containers.
it panics if the RecordSet does not contain exactly one row.
*/
func (rs RecordSet) ReadOne(container interface{}, cols ...string) {
	//if err := checkStructPtr(container); err != nil {
	//	panic(fmt.Errorf("recordSet `%s` ReadOne() error: %s", rs, err))
	//}
	//if err := rs.query.One(container, cols...); err != nil {
	//	panic(fmt.Errorf("recordSet `%s` ReadOne() error: %s", rs, err))
	//}
	//rs.computeFields(container)
}

///*
//Values query all data of the RecordSet and map to []map[string]interface.
//exprs means condition expression.
//it converts data to []map[column]value.
//*/
//func (rs RecordSet) Values(results *[]FieldMap, exprs ...string) int64 {
//	dbFields := filteredOnDBFields(rs.ModelName(), exprs)
//	num, err := rs.query.Values(results, dbFields...)
//	if err != nil {
//		panic(fmt.Errorf("recordSet `%s` Values() error: %s", rs, err))
//	}
//	for i, rec := range rs.Records() {
//		rec.computeFieldValues(&(*results)[i], exprs...)
//	}
//	return num
//}

///*
//ValuesFlat query all data and map to []interface.
//it's designed for one column record set, auto change to []value, not [][column]value.
//*/
//func (rs RecordSet) ValuesFlat(result *[]interface{}, expr string) int64 {
//	if getFieldColumn(rs.ModelName(), expr) != "" {
//		// expr is a stored field
//		num, err := rs.query.ValuesFlat(result, expr)
//		if err != nil {
//			panic(fmt.Errorf("recordSet `%s` ValuesFlat() error: %s", rs, err))
//		}
//		return num
//	} else {
//		// expr is a computed field
//		*result = make([]interface{}, int(rs.SearchCount()))
//		for i, rec := range rs.Records() {
//			params := make(FieldMap)
//			rec.computeFieldValues(&params, expr)
//			(*result)[i] = params[GetFieldJSON(rs.ModelName(), expr)]
//		}
//		return 0
//	}
//}

/*
Call calls the given method name methName with the given arguments and return the
result as interface{}.
*/
func (rs RecordSet) Call(methName string, args ...interface{}) interface{} {
	methInfo, ok := methodsCache.get(method{modelName: rs.ModelName(), name: methName})
	if !ok {
		panic(fmt.Errorf("Unknown method `%s` in model `%s`", methName, rs.ModelName()))
	}
	methLayer := methInfo.topLayer

	rsCopy := copyRecordStruct(rs)
	rsCopy.callStack = append([]*methodLayer{methLayer}, rsCopy.callStack...)
	return rsCopy.call(methLayer, args...)
}

/*
call is a wrapper around reflect.Value.Call() to use with interface{} type.
*/
func (rs RecordSet) call(methLayer *methodLayer, args ...interface{}) interface{} {
	fnVal := methLayer.funcValue
	fnTyp := fnVal.Type()

	rsVal := reflect.ValueOf(rs)
	inVals := []reflect.Value{rsVal}
	methName := fmt.Sprintf("%s.%s()", methLayer.methInfo.ref.modelName, methLayer.methInfo.ref.name)
	for i := 1; i < fnTyp.NumIn(); i++ {
		if i > len(args) {
			panic(fmt.Errorf("Not enough argument when Calling `%s`", methName))
		}
		inVals = append(inVals, reflect.ValueOf(args[i-1]))
	}
	retVal := fnVal.Call(inVals)
	if len(retVal) == 0 {
		return nil
	}
	return retVal[0].Interface()
}

/*
Super calls the next method Layer after the given funcPtr.
This method is meant to be used inside a method layer function to call its parent,
passing itself as funcPtr.
*/
func (rs RecordSet) Super(args ...interface{}) interface{} {
	if len(rs.callStack) == 0 {
		panic(fmt.Errorf("Internal error: empty call stack !"))
	}
	methLayer := rs.callStack[0]
	methInfo := methLayer.methInfo
	methLayer = methInfo.getNextLayer(methLayer)
	if methLayer == nil {
		// No parent
		return nil
	}

	rsCopy := copyRecordStruct(rs)
	rsCopy.callStack[0] = methLayer
	return rsCopy.call(methLayer, args...)
}

/*
MethodType returns the type of the method given by methName
*/
func (rs RecordSet) MethodType(methName string) reflect.Type {
	methInfo, ok := methodsCache.get(method{modelName: rs.ModelName(), name: methName})
	if !ok {
		panic(fmt.Errorf("Unknown method `%s` in model `%s`", methName, rs.ModelName()))
	}
	return methInfo.methodType
}

/*
Records returns the slice of RecordSet singletons that constitute this RecordSet
*/
func (rs RecordSet) Records() []*RecordSet {
	rs.Search()
	res := make([]*RecordSet, len(rs.Ids()))
	for i, id := range rs.Ids() {
		res[i] = rs.withIds([]int64{id})
	}
	return res
}

/*
EnsureOne panics if rs is not a singleton
*/
func (rs RecordSet) EnsureOne() {
	rs.Search()
	if len(rs.Ids()) != 1 {
		panic(fmt.Errorf("Expected singleton, got : %s", rs))
	}
}

/*
withIdMap returns a copy of rs filtered on the given ids slice (overwriting current queryset).
*/
func (rs RecordSet) withIds(ids []int64) *RecordSet {
	newRs := copyRecordStruct(rs)
	newRs.ids = ids
	//newRs.query = rs.env.Cr().QueryTable(rs.ModelName())
	//if len(ids) > 0 {
	//	domStr := fmt.Sprintf("id%sin", ExprSep)
	//	newRs.query = newRs.query.Filter(domStr, ids)
	//}
	return newRs
}

///*
//computeFields sets the value of the computed (non stored) fields of structPtr.
//*/
//func (rs RecordSet) computeFields(structPtr interface{}) {
//	val := reflect.ValueOf(structPtr)
//	ind := reflect.Indirect(val)
//
//	fInfos, _ := fieldsCache.getComputedFields(rs.ModelName())
//	params := make(FieldMap)
//	for _, fInfo := range fInfos {
//		sf := ind.FieldByName(fInfo.name)
//		if !sf.IsValid() {
//			// Computed field is not present in structPtr
//			continue
//		}
//		if _, exists := params[fInfo.name]; exists {
//			// We already have the value we need in params
//			continue
//		}
//		newParams := rs.Call(fInfo.compute).(orm.Params)
//		for k, v := range newParams {
//			params[k] = v
//		}
//		structField := ind.FieldByName(fInfo.name)
//		structField.Set(reflect.ValueOf(params[fInfo.name]))
//	}
//}

///*
//computeFieldValues updates the given params with the given computed (non stored) fields
//or all the computed fields of the model if not given.
//*/
//func (rs RecordSet) computeFieldValues(params *FieldMap, fields ...string) {
//	fInfos, _ := fieldsCache.getComputedFields(rs.ModelName(), fields...)
//	for _, fInfo := range fInfos {
//		if _, exists := (*params)[fInfo.name]; exists {
//			// We already have the value we need in params
//			// probably because it was computed with another field
//			continue
//		}
//		newParams := rs.Call(fInfo.compute).(FieldMap)
//		for k, v := range newParams {
//			key := GetFieldJSON(rs.ModelName(), k)
//			(*params)[key] = v
//		}
//	}
//}

///*
//updateStoredFields updates all dependent fields of rs that are included in structPtrOrParams.
//*/
//func (rs RecordSet) updateStoredFields(structPtrOrParams interface{}) {
//	// First get list of fields that have been passed through structPtrOrParams
//	var fieldNames []string
//	if params, ok := structPtrOrParams.(FieldMap); ok {
//		cpsf, _ := fieldsCache.getComputedStoredFields(rs.ModelName())
//		fieldNames = make([]string, len(params)+len(cpsf))
//		i := 0
//		for k, _ := range params {
//			fieldNames[i] = k
//			i++
//		}
//		for _, v := range cpsf {
//			fieldNames[i] = v.name
//			i++
//		}
//	} else {
//		val := reflect.ValueOf(structPtrOrParams)
//		typ := reflect.Indirect(val).Type()
//		fieldNames = make([]string, typ.NumField())
//		for i := 0; i < typ.NumField(); i++ {
//			fieldNames[i] = typ.Field(i).Name
//		}
//	}
//	// Then get all fields to update
//	var toUpdate []computeData
//	for _, fieldName := range fieldNames {
//		refField := fieldRef{modelName: rs.ModelName(), name: fieldName}
//		targetFields, ok := fieldsCache.getDependentFields(refField)
//		if !ok {
//			continue
//		}
//		toUpdate = append(toUpdate, targetFields...)
//	}
//	// Compute all that must be computed and store the values
//	computed := make(map[string]bool)
//	rs = rs.Search()
//	for _, cData := range toUpdate {
//		methUID := fmt.Sprintf("%s.%s", cData.modelName, cData.compute)
//		if _, ok := computed[methUID]; ok {
//			continue
//		}
//		recs := NewRecordSet(rs.env, cData.modelName)
//		if cData.path != "" {
//			domainString := fmt.Sprintf("%s%s%s", cData.path, orm.ExprSep, "in")
//			recs.Filter(domainString, rs.Ids())
//		} else {
//			recs = rs
//		}
//		for _, rec := range recs.Records() {
//			vals := rec.Call(cData.compute)
//			if len(vals.(FieldMap)) > 0 {
//				rec.Write(vals.(FieldMap))
//			}
//		}
//	}
//}

/*
newRecordStruct returns a new empty recordStruct.
*/
func newRecordStruct(env *Environment, ptrStructOrTableName interface{}) *RecordSet {
	//modelName := getModelName(ptrStructOrTableName)
	//qs := env.Cr().QueryTable(modelName)
	rs := RecordSet{
		query: Query{},
		env:   NewEnvironment(env.Cr(), env.Uid(), env.Context()),
		ids:   make([]int64, 0),
	}
	rs.query.recordSet = &rs
	return &rs
}

///*
//newRecordStructFromData returns a recordStruct pointing to data.
//*/
//func newRecordStructFromData(env Environment, data interface{}) *RecordSet {
//	rs := newRecordStruct(env, data)
//	if err := checkStructPtr(data); err != nil {
//		panic(fmt.Errorf("newRecordStructFromData: %s", err))
//	}
//	val := reflect.ValueOf(data)
//	ind := reflect.Indirect(val)
//	id := ind.FieldByName("ID").Int()
//	return rs.withIds([]int64{id})
//}

func copyRecordStruct(rs RecordSet) *RecordSet {
	newRs := newRecordStruct(rs.env, rs.ModelName())
	newRs.query = rs.query
	newRs.ids = make([]int64, len(rs.ids))
	copy(newRs.ids, rs.ids)
	newRs.callStack = make([]*methodLayer, len(rs.callStack))
	copy(newRs.callStack, rs.callStack)
	return newRs
}

/*
NewRecordSet returns a new empty Recordset on the model given by ptrStructOrTableName and the
given Environment.
*/
func NewRecordSet(env *Environment, ptrStructOrTableName interface{}) *RecordSet {
	//return newRecordStruct(env, ptrStructOrTableName)
	return new(RecordSet)
}
