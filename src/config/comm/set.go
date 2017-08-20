package comm

import "strings"

// Set 向map设置数据
// 支持map嵌套，每一层key之间用常量sep定义的字符串连接
func (m *Map) set(key string, val interface{}) error {
	m.Lock()
	defer m.Unlock()

	keys := strings.Split(key, m.KeySep)
	vals := m.Data
	l := len(keys)
	for i, key := range keys {
		// 当前是最后一层，则直接设置
		if i == l-1 {
			vals[key] = val
			return nil
		}
		// 当前不是最后一层
		v, ok := vals[key]
		if !ok {
			// 不存在，则新建一个map
			mm := make(map[string]interface{})
			vals[key] = mm
			vals = mm
		} else {
			// 存在，则必须为map类型
			vals, ok = v.(map[string]interface{})
			if !ok {
				return ErrNotExistKey
			}
		}
	}
	return ErrInvalidKey
}

// SetBool 设置指定key的bool值
func (m *Map) SetBool(key string, val bool) error {
	return m.set(key, val)
}

// SetInt 设置指定key的int值
func (m *Map) SetInt(key string, val int) error {
	return m.set(key, val)
}

// SetInt64 设置指定key的int64值
func (m *Map) SetInt64(key string, val int64) error {
	return m.set(key, val)
}

// SetUint 设置指定key的uint值
func (m *Map) SetUint(key string, val uint) error {
	return m.set(key, val)
}

// SetUint64 设置指定key的uint64值
func (m *Map) SetUint64(key string, val uint64) error {
	return m.set(key, val)
}

// SetFloat 设置指定key的float64值
func (m *Map) SetFloat(key string, val float64) error {
	return m.set(key, val)
}

// SetString 设置指定key的string值
func (m *Map) SetString(key string, val string) error {
	return m.set(key, val)
}

// SetStrings 设置指定key的string数组
func (m *Map) SetStrings(key string, val []string) error {
	return m.set(key, strings.Join(val, m.ArraySep))
}
