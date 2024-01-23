package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName("db-test")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("data", "data")
}

func Load() error {
	return viper.ReadInConfig()
}
func Store() error {
	return viper.SafeWriteConfig()
}

func Register(module string, key string, value any) {
	viper.SetDefault(module+"."+key, value)
}
func GetString(module string, key string) string {
	return viper.GetString(module + "." + key)
}

func GetBool(module string, key string) bool {
	return viper.GetBool(module + "." + key)
}
func GetInt(module string, key string) int {
	return viper.GetInt(module + "." + key)
}
