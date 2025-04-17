package filterx

import (
	"encoding/json"
)

// constraints 约束支持比较的类型
type constraints interface {
	any
}

// Range 表示范围谓词
type Range[T constraints] struct {
	Max T `json:"max"`
	Min T `json:"min"`
}

func (r Range[T]) Filter() Filter[T] {
	return Filter[T]{rangePtr: &r}
}

// MarshalJSON 实现 json.Marshaler 接口
func (r Range[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"range": struct {
		Max T `json:"max"`
		Min T `json:"min"`
	}{r.Max, r.Min}})
}

// Like 表示匹配谓词
type Like[T constraints] struct {
	Pattern T `json:"pattern"`
}

func (e Like[T]) Filter() Filter[T] {
	return Filter[T]{likePtr: &e}
}

// MarshalJSON 实现 json.Marshaler 接口
func (e Like[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"like": e.Pattern})
}

// EQ 表示相等谓词
type EQ[T constraints] struct {
	Value T `json:"value"`
}

func (e EQ[T]) Filter() Filter[T] {
	return Filter[T]{eqPtr: &e}
}

// MarshalJSON 实现 json.Marshaler 接口
func (e EQ[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"eq": e.Value})
}

// NEQ 表示不相等谓词
type NEQ[T constraints] struct {
	Value T `json:"value"`
}

func (e NEQ[T]) Filter() Filter[T] {
	return Filter[T]{neqPtr: &e}
}

// MarshalJSON 实现 json.Marshaler 接口
func (e NEQ[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"neq": e.Value})
}

// LT 表示小于谓词
type LT[T constraints] struct {
	Value T `json:"value"`
}

func (l LT[T]) Filter() Filter[T] {
	return Filter[T]{ltPtr: &l}
}

// MarshalJSON 实现 json.Marshaler 接口
func (l LT[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"lt": l.Value})
}

// GT 表示大于谓词
type GT[T constraints] struct {
	Value T `json:"value"`
}

func (g GT[T]) Filter() Filter[T] {
	return Filter[T]{gtPtr: &g}
}

// MarshalJSON 实现 json.Marshaler 接口
func (g GT[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"gt": g.Value})
}

// In 表示包含在集合中的谓词
type In[T constraints] struct {
	Values []T `json:"values"`
}

func (i In[T]) Filter() Filter[T] {
	return Filter[T]{inPtr: &i}
}

// MarshalJSON 实现 json.Marshaler 接口
func (i In[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]T{"in": i.Values})
}

// NIn 表示不包含在集合中的谓词
type NIn[T constraints] struct {
	Values []T `json:"values"`
}

func (i NIn[T]) Filter() Filter[T] {
	return Filter[T]{ninPtr: &i}
}

// MarshalJSON 实现 json.Marshaler 接口
func (i NIn[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]T{"nin": i.Values})
}

// GE 表示大于等于谓词
type GE[T constraints] struct {
	Value T
}

func (g GE[T]) Filter() Filter[T] {
	return Filter[T]{gePtr: &g}
}

// MarshalJSON 实现 json.Marshaler 接口
func (g GE[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"ge": g.Value})
}

// LE 表示小于等于谓词
type LE[T constraints] struct {
	Value T
}

func (l LE[T]) Filter() Filter[T] {
	return Filter[T]{lePtr: &l}
}

// MarshalJSON 实现 json.Marshaler 接口
func (l LE[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]T{"le": l.Value})
}

// Filter 包装不同的谓词类型
type Filter[T constraints] struct {
	rangePtr *Range[T]
	likePtr  *Like[T]
	eqPtr    *EQ[T]
	ltPtr    *LT[T]
	gtPtr    *GT[T]
	inPtr    *In[T]
	ninPtr   *NIn[T]
	gePtr    *GE[T]
	lePtr    *LE[T]
	neqPtr   *NEQ[T]
}

// GetNIn 获取 NIn 谓词
func (f Filter[T]) GetNIn() (NIn[T], bool) {
	if f.rangePtr != nil {
		return *f.ninPtr, true
	}
	return NIn[T]{}, false
}

// GetRange 获取 Range 谓词
func (f Filter[T]) GetRange() (Range[T], bool) {
	if f.rangePtr != nil {
		return *f.rangePtr, true
	}
	return Range[T]{}, false
}

// GetLike 获取 like 谓词
func (f Filter[T]) GetLike() (Like[T], bool) {
	if f.rangePtr != nil {
		return *f.likePtr, true
	}
	return Like[T]{}, false
}

// GetEQ 获取 EQ 谓词
func (f Filter[T]) GetEQ() (EQ[T], bool) {
	if f.eqPtr != nil {
		return *f.eqPtr, true
	}
	return EQ[T]{}, false
}

// GetNEQ 获取 NEQ 谓词
func (f Filter[T]) GetNEQ() (NEQ[T], bool) {
	if f.neqPtr != nil {
		return *f.neqPtr, true
	}
	return NEQ[T]{}, false
}

// GetLT 获取 LT 谓词
func (f Filter[T]) GetLT() (LT[T], bool) {
	if f.ltPtr != nil {
		return *f.ltPtr, true
	}
	return LT[T]{}, false
}

// GetGT 获取 GT 谓词
func (f Filter[T]) GetGT() (GT[T], bool) {
	if f.gtPtr != nil {
		return *f.gtPtr, true
	}
	return GT[T]{}, false
}

// GetIn 获取 In 谓词
func (f Filter[T]) GetIn() (In[T], bool) {
	if f.inPtr != nil {
		return *f.inPtr, true
	}
	return In[T]{}, false
}

// GetGE 获取 GE 谓词
func (f Filter[T]) GetGE() (GE[T], bool) {
	if f.gePtr != nil {
		return *f.gePtr, true
	}
	return GE[T]{}, false
}

// GetLE 获取 LE 谓词
func (f Filter[T]) GetLE() (LE[T], bool) {
	if f.lePtr != nil {
		return *f.lePtr, true
	}
	return LE[T]{}, false
}

// Predicate 表示可以生成 Filter 的谓词接口
type Predicate[T constraints] interface {
	Filter() Filter[T]
	MarshalJSON() ([]byte, error)
}
