/**
 * Created by GoLand
 * Author: zhangjian
 * Date: 2019-06-21 14:30
 */

package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"git.medlinker.com/service/common/ecode/errors"
)

// CHANGE
// 1: 去除字段兜底赋值(指定目标字段就是不想要源字段)
// 2: 新增结构体多级字段指定(copy:"med.pid"),会自动忽略多余的点

/*
type (
	dest struct {
		Id        *int64 `copy:"pid"`      // source中Pid
		Name      string `copy:"noname"`   // source不存在noname, 使用Name
		Age       int    `copy:"real_age"` // source中的RealAge
		Ignore    bool   `copy:"-"`        // 忽略该字段
		Status    bool   // source中的Status
		AliasName string `copy:"alias_name, origin"` // origin代表不对copy中的值做转换
	}
	source struct {
		Pid     int
		Name    string
		RealAge int
		Ignore  bool
	}
)

var anotherSource = map[string]interface{}{"pid": 1, "alias_name": "med"}
*/

// 字段赋值
// 结构体赋值时对应字段可指定源的字段(`copy:"source's field"`)
type Copy struct {
	// 赋值体
	// 请使用SetSource方法
	Source interface{}

	// 强制转换(慎用)
	// 多级指针只有层级一致或者dv和sv最后都不再是指针才不会有问题
	// Recursion时, 多级指针层级一致或者sv最后不再是指针才不会有问题
	Convert bool

	// 出错是否继续下一个
	Next bool

	// 是否递归(结构体、数组)
	// 不过这样也会让赋值都变成值赋值而不是遇到指针是引用，也就是部分指针会重新申请内存，类似于深度拷贝
	// 并且只有在Convert为true时生效
	// 递归有性能损耗，并且如果赋值中有循环依赖可能导致死循环，与json解析一样，慎用
	Recursion bool

	// 当copy赋值字段为空时是否需要读取JSON TAG中的字段
	JSONTag bool
}

const OriginCopyField = "origin"

// 初始化默认
func NewCopy() *Copy {
	return &Copy{Convert: true, Next: true, Recursion: true, JSONTag: true}
}

// 克隆
func (c *Copy) Clone() *Copy {
	return &Copy{
		Source:    c.Source,
		Convert:   c.Convert,
		Next:      c.Next,
		Recursion: c.Recursion,
		JSONTag:   c.JSONTag,
	}
}

// 是否强转
func (c *Copy) SetConvert(convert bool) *Copy {
	cp := c.Clone()
	cp.Convert = convert
	return cp
}

// 出错是否继续赋值下一个字段
func (c *Copy) SetNext(next bool) *Copy {
	cp := c.Clone()
	cp.Next = next
	return cp
}

// 是否递归（依赖强转）
func (c *Copy) SetRecursion(recursion bool) *Copy {
	cp := c.Clone()
	cp.Recursion = recursion
	return cp
}

// 是否读取JSON TAG
func (c *Copy) SetJSONTag(jsonTag bool) *Copy {
	cp := c.Clone()
	cp.JSONTag = jsonTag
	return cp
}

// 设置源
func (c *Copy) SetSource(source interface{}) *Copy {
	cp := c.Clone()
	cp.Source = source
	return cp
}

// 重新定义赋值(带上dv的类型)
func (c *Copy) copyStf(dv, sv reflect.Value, dt reflect.Type) (ok bool, err error) {
	var (
		fst    reflect.StructField
		sf     string
		origin bool
		se     error
		aok    bool
		fv     reflect.Value
	)
	// 赋值所有可能的字段
	num := dt.NumField()
	for i := 0; i < num; i++ {
		fv = sv // 中间赋值体，为了多级字段方便运用
		fst = dt.Field(i)
		sf, origin = c.parseTag(fst)
		// 解析多级字段并且查看是否符合预期
		sf = strings.Trim(sf, ".")
		if sf != "" && strings.Index(sf, ".") > -1 {
			tmpSfs := strings.Split(sf, ".")
			sfs := make([]string, 0, len(tmpSfs))
			sfi := 0
			for i := range tmpSfs {
				sff := strings.TrimSpace(tmpSfs[i])
				if sff != "" {
					sfs = append(sfs, sff)
					sfi++
				}
			}
			if sfi == 0 {
				sf = ""
			} else if sfi == 1 {
				sf = sfs[0]
			} else {
				sf = sfs[sfi-1]
				fv, aok = c.parseMultiField(sfs[:sfi-1], sv, origin)
				if !aok && !c.Next {
					err = errors.Errorf("赋值字段[%s]赋值失败: 源字段不存在或为nil", fst.Name)
					ok = false
					return
				}
			}
		}
		// 如果是内嵌字段并且没有指定字段
		if sf == "" && fst.Anonymous {
			dat := fst.Type
			dav := dv.FieldByName(fst.Name)
			isNil := false
			if fst.Type.Kind() == reflect.Ptr {
				dat = dat.Elem()
				if dav.IsNil() {
					if !dav.CanSet() {
						continue
					}
					dav = reflect.New(dat)
					isNil = true
				}
				dav = dav.Elem()

			}
			// 如果是结构体
			if dat.Kind() == reflect.Struct {
				aok, se = c.copyStf(dav, fv, dat)
				if se == nil && aok && isNil {
					dv.FieldByName(fst.Name).Set(dav.Addr())
				}
			}
		} else {
			aok, se = c.SetSf(dv, fv, fst.Name, sf, origin)
		}
		if se == nil || c.Next {
			if aok {
				ok = true
			}
			continue
		}
		err = errors.Wrapf(se, "赋值字段[%s]赋值失败", fst.Name)
		ok = false
		return
	}
	return
}

func (c *Copy) copySf(dv, sv reflect.Value) (bool, error) {
	return c.copyStf(dv, sv, dv.Type())
}

// 解析多级
// 支持数组、map、结构体
func (c *Copy) parseMultiField(sfs []string, sv reflect.Value, osf bool) (v reflect.Value, aok bool) {
	defer func() {
		if pe := recover(); pe != nil {
			aok = false
			return
		}
	}()
	for _, sf := range sfs {
		kind := sv.Kind()
		for kind == reflect.Ptr {
			if sv.IsNil() {
				return sv, false
			}
			sv = sv.Elem()
			kind = sv.Kind()
		}
		if kind == reflect.Array || kind == reflect.Slice {
			// 解析sf为数字
			d, err := strconv.Atoi(sf)
			if err != nil {
				return sv, false
			}
			sv = sv.Index(d)
			if !sv.IsValid() {
				return sv, false
			}
			continue
		}
		if kind == reflect.Map {
			if !osf {
				sf = ToSnake(sf)
			}
			sv = sv.MapIndex(reflect.ValueOf(sf))
			if !sv.IsValid() {
				return sv, false
			}
			continue
		}
		if kind == reflect.Struct {
			if !osf {
				sf = ToCame(sf)
			}
			sv = sv.FieldByName(sf)
			if !sv.IsValid() {
				return sv, false
			}
			continue
		}
		return sv, false
	}
	// 兜底去除指针
	for sv.Kind() == reflect.Ptr {
		if sv.IsNil() {
			return sv, false
		}
		sv = sv.Elem()
	}
	return sv, true
}

// 解析赋值字段
// NOTE
// 优先解析copy，然后解析json，因为按照规则和习惯，大部分的字段最终名与json后的字段名一致
func (c *Copy) parseTag(fst reflect.StructField) (string, bool) {
	// json -
	// copy origin
	sfi := strings.TrimSpace(fst.Tag.Get("copy"))
	sf := ""
	// split一定不为空
	sfs := strings.SplitN(sfi, ",", 2)
	sf = strings.TrimSpace(sfs[0])
	osf := len(sfs) > 1 && strings.TrimSpace(sfs[1]) == OriginCopyField
	if sf == "" && c.JSONTag {
		sf = strings.TrimSpace(strings.SplitN(strings.TrimSpace(fst.Tag.Get("json")), ",", 2)[0])
	}
	return sf, osf
}

// 抽离函数，过滤不合法的字段
func (c *Copy) SetSf(dv reflect.Value, sv reflect.Value, df, sf string, osf bool) (ok bool, err error) {
	// 强制转换可能会出现异常
	defer func() {
		if pe := recover(); pe != nil {
			err = fmt.Errorf("赋值失败: [%#v]", pe)
			ok = false
		}
	}()

	// 检车字段是否有效
	df = strings.TrimSpace(df)
	sf = strings.TrimSpace(sf)
	if sf == "" {
		sf = df
	}
	if df == "" && sf == "" || sf == "-" {
		return
	}

	dFieldValue := dv.FieldByName(df)
	// 判断是否结构体字段存在
	if !dFieldValue.IsValid() || !dFieldValue.CanSet() {
		return
	}
	// 源字段检查
	var sFieldValue reflect.Value
	stKind := sv.Kind()
	if stKind == reflect.Array || stKind == reflect.Slice {
		d := 0
		d, err = strconv.Atoi(sf)
		if err != nil {
			err = errors.Wrapf(err, "被赋值字段[%s]对应源字段必须是数字[%s]", df, sf)
		}
		sFieldValue = sv.Index(d)
	} else if stKind == reflect.Struct {
		// 结构体
		if !osf {
			sf = ToCame(sf)
		}
		sFieldValue = sv.FieldByName(sf)
	} else if stKind == reflect.Map {
		// map
		if !osf {
			sf = ToSnake(sf)
		}
		sFieldValue = sv.MapIndex(reflect.ValueOf(sf))
	}
	if !sFieldValue.IsValid() {
		err = errors.New("赋值类型必须是struct、map或者array(slice); 源字段不存在")
		return
	}
	ok = c.Value(dFieldValue, sFieldValue)
	return
}

// NOTE 复杂逻辑不符合预期: 递归结构体、多级指针
// 抽离函数，真正的赋值操作
func (c *Copy) Value(dFieldValue reflect.Value, sFieldValue reflect.Value) bool {
	dt := dFieldValue.Type()
	st := sFieldValue.Type()
	// nil返回
	if (st.Kind() == reflect.Ptr || st.Kind() == reflect.Interface) && sFieldValue.IsNil() {
		return false
	}

	// 判断类型是否一样或者被赋值是interface
	if dt == st || dt.Kind() == reflect.Interface {
		dFieldValue.Set(sFieldValue)
		return true
	}

	// 如果赋值是interface
	if st.Kind() == reflect.Interface {
		return c.Value(dFieldValue, reflect.ValueOf(sFieldValue.Interface()))
	}

	// 是否强制转换
	if !c.Convert {
		return false
	}

	// 找到未初始化的指针
	for dt.Kind() == reflect.Ptr && !dFieldValue.IsNil() {
		if st.Kind() == reflect.Ptr {
			if sFieldValue.IsNil() {
				return false
			}
			sFieldValue = sFieldValue.Elem()
			st = sFieldValue.Type()
		}
		dFieldValue = dFieldValue.Elem()
		dt = dFieldValue.Type()
	}

	if !c.Recursion {
		return c.notRecursion(dt, st, dFieldValue, sFieldValue)
	}
	// 指针赋值(尽可能重新申请内存拷贝)
	isMalloc := false
	tdFieldValue := dFieldValue
	tmpFieldValue := tdFieldValue
	ok, next := c.recursionPointer(&dt, &st, &tdFieldValue, &tmpFieldValue, &sFieldValue, dFieldValue, &isMalloc)
	if ok {
		return true
	}
	if !next {
		return false
	}
	// 指针过后类型还是不同的话，这里做个判断是为了防止死循环
	if st != dt {
		// 如果都为数组
		if (dt.Kind() == reflect.Array || dt.Kind() == reflect.Slice) &&
			(st.Kind() == reflect.Array || st.Kind() == reflect.Slice) {
			return c.recursionSlice(dt, dFieldValue, tdFieldValue, tmpFieldValue, sFieldValue, isMalloc)
		}
		// 结构体
		if dt.Kind() == reflect.Struct && (st.Kind() == reflect.Map || st.Kind() == reflect.Struct) {
			// note 如果该字段与source是同一个，则会死循环
			return c.recursionStruct(dFieldValue, tdFieldValue, tmpFieldValue, sFieldValue, isMalloc)
		}
		c.convert(dt, st, tmpFieldValue, sFieldValue)
		if isMalloc {
			dFieldValue.Set(tdFieldValue)
		}
	}
	return true
}

// 强制转换: 调用该方法的时候已经不再是指针
func (c *Copy) convert(dt, st reflect.Type, dFieldValue, sFieldValue reflect.Value) bool {
	// 对时间类型做特殊处理
	// 后续有需要可以改为类型注册函数处理
	if st.PkgPath() == "time" && st.String() == "time.Time" {
		switch dt.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			fallthrough
		case reflect.Int64:
			sFieldValue = reflect.ValueOf(sFieldValue.Interface().(time.Time).Unix())
		case reflect.String:
			sFieldValue = reflect.ValueOf(sFieldValue.Interface().(time.Time).Format("2006-01-02 15:04:05"))
		}
	} else if dt.PkgPath() == "time" && dt.String() == "time.Time" {
		switch st.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			sFieldValue = reflect.ValueOf(time.Unix(sFieldValue.Int(), 0))
		case reflect.Uint:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			sFieldValue = reflect.ValueOf(time.Unix(int64(sFieldValue.Uint()), 0)) // 如果uint64超过int64，这里会抛异常
		case reflect.String:
			now, err := time.ParseInLocation("2006-01-02 15:04:05", sFieldValue.String(), time.Local)
			if err == nil {
				sFieldValue = reflect.ValueOf(now)
			}
		}
	}
	dFieldValue.Set(sFieldValue.Convert(dt))
	return true
}

// 不递归进行指针赋值
func (c *Copy) notRecursion(dt, st reflect.Type, dFieldValue, sFieldValue reflect.Value) bool {
	// 赋值指针(不申请内存拷贝)
	if dt.Kind() == reflect.Ptr {
		// 转换赋值为指针
		sv := sFieldValue
		if st.Kind() != reflect.Ptr {
			sv = reflect.New(st)
			sv.Elem().Set(sFieldValue)
		} else if sFieldValue.IsNil() {
			return false
		}
		// 赋值
		dFieldValue.Set(reflect.NewAt(dt.Elem(), unsafe.Pointer(sv.Pointer())))
		return true
	}
	// 多级指针
	for st.Kind() == reflect.Ptr {
		if sFieldValue.IsNil() {
			return false
		}
		sFieldValue = sFieldValue.Elem()
		st = sFieldValue.Type()
	}
	if st == dt {
		dFieldValue.Set(sFieldValue)
		return true
	}
	return c.convert(dt, st, dFieldValue, sFieldValue)
}

// 递归指针赋值
func (c *Copy) recursionPointer(
	dt, st *reflect.Type,
	tdFieldValue, tmpFieldValue, sFieldValue *reflect.Value,
	dFieldValue reflect.Value, malloc *bool) (bool, bool) {
	for {
		if *st == *dt {
			(*tmpFieldValue).Set(*sFieldValue)
			if *malloc {
				dFieldValue.Set(*tdFieldValue)
			}
			return true, false
		}
		if (*dt).Kind() != reflect.Ptr && (*st).Kind() != reflect.Ptr {
			break
		}
		if (*dt).Kind() == reflect.Ptr {
			// 指针未初始化
			if !(*malloc) {
				*malloc = true
				*tdFieldValue = reflect.New((*dt).Elem())
				*tmpFieldValue = *tdFieldValue
			} else {
				(*tmpFieldValue).Set(reflect.New((*dt).Elem()))
			}

			*tmpFieldValue = (*tmpFieldValue).Elem()
			*dt = tmpFieldValue.Type()
		}
		if (*st).Kind() == reflect.Ptr {
			if (*sFieldValue).IsNil() {
				// 没办法赋值
				return false, false
			}
			*sFieldValue = (*sFieldValue).Elem()
			*st = (*sFieldValue).Type()
		}
	}
	return false, true
}

// 递归赋值数组
func (c *Copy) recursionSlice(
	tdt reflect.Type, dFieldValue, tdFieldValue, tmpFieldValue, sFieldValue reflect.Value, malloc bool) bool {
	sl := sFieldValue.Len()
	if sl == 0 {
		return false
	}
	if tdt.Kind() == reflect.Array {
		dl := tmpFieldValue.Len()
		if dl < sl {
			sl = dl
		}
	} else {
		tmpFieldValue.Set(reflect.MakeSlice(tdt, sl, sl))
	}
	ok := false
	for i := 0; i < sl; i++ {
		if !c.Value(tmpFieldValue.Index(i), sFieldValue.Index(i)) {
			continue
		}
		ok = true
	}
	if ok && malloc {
		dFieldValue.Set(tdFieldValue)
	}
	return ok
}

// 递归结构体
func (c *Copy) recursionStruct(dFieldValue, tdFieldValue, tmpFieldValue, sFieldValue reflect.Value, malloc bool) bool {
	ok, err := c.copySf(tmpFieldValue, sFieldValue)
	if err != nil {
		panic(errors.Wrap(err, "递归结构体赋值失败"))
	}
	if ok && malloc {
		dFieldValue.Set(tdFieldValue)
	}
	return ok
}

// 为dest在source中存在的字段自动赋值
// 结构体字段赋值函数
// 通用请调用CopyF方法
func (c *Copy) CopySF(dest interface{}) (err error) {
	if c.Source == nil {
		return errors.New("赋值体不存在")
	}
	defer func() {
		if pe := recover(); pe != nil {
			err = fmt.Errorf("赋值失败: [%#v]", pe)
		}
	}()
	// 校验类型
	dv := reflect.ValueOf(dest)
	dt := dv.Type()
	if dt.Kind() != reflect.Ptr {
		err = errors.New("被赋值的结构体必须是指针类型")
		return
	}

	// 真实数据
	dv = dv.Elem()
	if dv.Kind() != reflect.Struct {
		err = errors.New("被赋值的不是结构体")
		return
	}

	sv := reflect.Indirect(reflect.ValueOf(c.Source))
	// 赋值必须是指针类型
	stKind := sv.Kind()
	if stKind != reflect.Struct && stKind != reflect.Map {
		err = errors.New("赋值类型必须是struct或者map")
		return
	}
	_, err = c.copySf(dv, sv)
	return
}

// 单个赋值
// 通用的赋值方法
func (c *Copy) CopyF(dest interface{}) (err error) {
	if c.Source == nil {
		return errors.New("赋值体不存在")
	}
	// 强制转换可能会出现异常
	defer func() {
		if pe := recover(); pe != nil {
			err = fmt.Errorf("单体赋值失败: [%#v]", pe)
		}
	}()
	// 校验类型
	dv := reflect.ValueOf(dest)
	dt := dv.Type()
	if dt.Kind() != reflect.Ptr {
		return errors.New("被赋值的单体必须是指针类型")
	}
	// 真实数据
	dv = dv.Elem()
	sv := reflect.Indirect(reflect.ValueOf(c.Source))
	c.Value(dv, sv)
	return
}
