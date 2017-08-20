package yaml

import (
	"io/ioutil"

	"github.com/nsknfd/go/src/config"
	"github.com/nsknfd/go/src/config/comm"

	"gopkg.in/yaml.v2"
)

type Config struct {
}

type Configer struct {
	comm.Map
}

// ParseData 从数据解析配置
func (c *Config) ParseData(data []byte) (config.Configer, error) {
	var m map[interface{}]interface{}
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	cnf := Configer{Map: comm.NewMap()}
	err = cnf.Parse(m)
	if err != nil {
		return nil, err
	}
	return &cnf, nil
}

// ParseData 从文件解析配置
func (c *Config) ParseFile(filename string) (config.Configer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return c.ParseData(data)
}

// SaveFile 将配置写入文件
func (c *Configer) SaveFile(filename string) error {
	c.RLock()
	m := c.Data
	c.RUnlock()
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func init() {
	config.Register("yaml", &Config{})
}
