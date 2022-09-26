package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
	"viper-test1/config"

	"github.com/spf13/viper"
)

func main() {
	viperLoadToml()
	initSys()
	viperLoadJson()
	viperLoadIni()
	iniLoadIni()

	// init config
	_, err := config.Init()
	if err != nil {
		fmt.Println(err)
	}
	// 注意：只能在 init 之后再次通过 viper.Get 方法读取配置，否则不生效
	for {
		cfg := &config.Config{
			Name:     viper.GetString("database.name"),
			Host:     viper.GetString("database.host"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
		}
		fmt.Println(cfg.Name)
		time.Sleep(4 * time.Second)
	}
}

func initSys() {
	fmt.Println("initSys*******")

	yamlFile, err := ioutil.ReadFile("./conf/sys.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _sys *config.Sys
	err = yaml.Unmarshal(yamlFile, &_sys)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("config.app: %#vn", _sys.App)
	fmt.Printf("config.log: %#vn", _sys.Log)
	fmt.Println()

}

// viper
// Get(key string) : interface{}
// GetBool(key string) : bool
// GetFloat64(key string) : float64
// GetInt(key string) : int
// GetIntSlice(key string) : []int
// GetString(key string) : string
// GetStringMap(key string) : map[string]interface{}
// GetStringMapString(key string) : map[string]string
// GetStringSlice(key string) : []string
// GetTime(key string) : time.Time
// GetDuration(key string) : time.Duration
// IsSet(key string) : bool
// AllSettings() : map[string]interface{}

func viperLoadJson() {
	fmt.Println("loadJson*******")

	config := viper.New()
	config.AddConfigPath("conf")
	config.SetConfigName("config")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
			// Config file was found but another error was produced
		}
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	version := config.GetString("version")
	origin := config.GetString("host.origin")
	fmt.Println(version)
	fmt.Println(origin)

	host := config.GetStringMapString("host")
	fmt.Println(host)
	fmt.Println(host["origin"])
	fmt.Println(host["port"])

	allsetting := config.AllSettings()
	fmt.Println(allsetting)
}

func viperLoadIni() {
	fmt.Println("loadIni*******")

	config := viper.New()
	config.AddConfigPath("conf")
	config.SetConfigName("app")
	config.SetConfigType("ini")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
			// Config file was found but another error was produced
		}
	}

	host := config.GetString("redis.host")
	fmt.Println("viper load ini: ", host)
}

func iniLoadIni() {
	fmt.Println("iniLoadIni******")

	file, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	mysqlUserName := file.Section("mysql").Key("username").String()
	mysqlPassWord := file.Section("mysql").Key("password").String()
	fmt.Println(mysqlUserName)
	fmt.Println(mysqlPassWord)
}

func viperLoadToml() {
	fmt.Println("viperLoadToml**********")

	config := viper.New()
	config.AddConfigPath("conf")
	config.SetConfigName("base")
	config.SetConfigType("toml")

	if err := config.ReadInConfig(); err != nil {
		panic("err")
	}

	//获取全部文件内容
	//fmt.Println("all settings: ", config.AllSettings())
	fmt.Println("--------------")
	//根据内容类型，解析出不同类型
	fmt.Println(config.GetString("database.server"))
	fmt.Println(config.GetIntSlice("database.ports"))
	fmt.Println("--------------")
	fmt.Println(config.GetString("servers.alpha.ip"))
}
