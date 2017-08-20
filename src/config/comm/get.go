package comm

import "strings"

// Get 从map获取数据
// 支持map嵌套，每一层key之间用常量sep定义的字符串连接
func (m *Map) get(key string) (interface{}, error) {
	m.RLock()
	defer m.RUnlock()

	keys := strings.Split(key, m.KeySep)
	vals := m.Data
	l := len(keys)
	for i, k := range keys {
		v, ok := vals[k]
		if !ok {
			return nil, ErrNotExistKey
		}
		// 解析到最后一层，直接返回
		if i == l-1 {
			return v, nil
		}
		// 不是最后一层，下一层必须是一个map
		vals, ok = v.(map[string]interface{})
		if !ok {
			return nil, ErrNotExistKey
		}
	}
	return nil, ErrInvalidKey
}

// GetBool 返回指定key的bool值
func (m *Map) GetBool(key string) (bool, error) {
	val, err := m.get(key)
	if err != nil {
		return false, err
	}
	if val == nil {
		return false, ErrInvalidValue
	}
	return ParseBool(val)
}

// GetInt 返回指定key的int值
func (m *Map) GetInt(key string) (int, error) {
	val, err := m.get(key)
	if err != nil {
		return 0, err
	}
	if val == nil {
		return 0, ErrInvalidValue
	}
	return ParseInt(val)
}

// GetInt64 返回指定key的int64值
func (m *Map) GetInt64(key string) (int64, error) {
	val, err := m.get(key)
	if err != nil {
		return 0, err
	}
	if val == nil {
		return 0, ErrInvalidValue
	}
	return ParseInt64(val)
}

// GetUint 返回指定key的uint值
func (m *Map) GetUint(key string) (uint, error) {
	val, err := m.get(key)
	if err != nil {
		return 0, err
	}
	if val == nil {
		return 0, ErrInvalidValue
	}
	return ParseUint(val)
}

// GetUint64 返回指定key的uint64值
func (m *Map) GetUint64(key string) (uint64, error) {
	val, err := m.get(key)
	if err != nil {
		return 0, err
	}
	if val == nil {
		return 0, ErrInvalidValue
	}
	return ParseUint64(val)
}

// GetFloat 返回指定key的float值
func (m *Map) GetFloat(key string) (float64, error) {
	val, err := m.get(key)
	if err != nil {
		return 0, err
	}
	if val == nil {
		return 0, ErrInvalidValue
	}
	return ParseFloat(val)
}

// GetString 返回指定key的string值
func (m *Map) GetString(key string) (string, error) {
	val, err := m.get(key)
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", ErrInvalidValue
	}
	return ParseString(val), nil
}

// GetStrings 返回指定key的string数组
func (m *Map) GetStrings(key string) ([]string, error) {
	val, err := m.GetString(key)
	if err != nil {
		return nil, err
	}
	if val == "" {
		return []string{}, nil
	}
	strs := strings.Split(val, m.ArraySep)
	for i, str := range strs {
		strs[i] = strings.TrimSpace(str)
	}
	return strs, nil
}

// DefaultBool 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultBool(key string, defaultVal bool) bool {
	v, err := m.GetBool(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultInt 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultInt(key string, defaultVal int) int {
	v, err := m.GetInt(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultInt64 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultInt64(key string, defaultVal int64) int64 {
	v, err := m.GetInt64(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultUint 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultUint(key string, defaultVal uint) uint {
	v, err := m.GetUint(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultUint64 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultUint64(key string, defaultVal uint64) uint64 {
	v, err := m.GetUint64(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultFloat 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultFloat(key string, defaultVal float64) float64 {
	v, err := m.GetFloat(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultString 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultString(key string, defaultVal string) string {
	v, err := m.GetString(key)
	if err != nil {
		return defaultVal
	}
	return v
}

// DefaultStrings 返回指定key的bool值，获取失败则返回默认值
func (m *Map) DefaultStrings(key string, defaultVal []string) []string {
	v, err := m.GetStrings(key)
	if err != nil {
		return defaultVal
	}
	return v
}
