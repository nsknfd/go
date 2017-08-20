package config

// 支持 key.key.key 获取子对象， 分隔符 "." 在不同配置器中可能不同
// 支持 str,str,str 表示的字符串数组，分隔符 "," 在不同配置器中可能不同
type Configer interface {
	SaveFile(filename string) error

	SetBool(key string, val bool) error
	SetInt(key string, val int) error
	SetInt64(key string, val int64) error
	SetUint(key string, val uint) error
	SetUint64(key string, val uint64) error
	SetFloat(key string, val float64) error
	SetString(key string, val string) error
	SetStrings(key string, val []string) error

	GetBool(key string) (bool, error)
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetUint(key string) (uint, error)
	GetUint64(key string) (uint64, error)
	GetFloat(key string) (float64, error)
	GetString(key string) (string, error)
	GetStrings(key string) ([]string, error)

	DefaultBool(key string, defaultVal bool) bool
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultUint(key string, defaultVal uint) uint
	DefaultUint64(key string, defaultVal uint64) uint64
	DefaultFloat(key string, defaultVal float64) float64
	DefaultString(key string, defaultVal string) string
	DefaultStrings(key string, defaultVal []string) []string
}
