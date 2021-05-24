package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type jwtsetting struct {
	ValidIssuer            string
	ValidAudiences         string
	PlatformServiceBaseUrl string
	ClientId               string
	ClientSecret           string
}
type Setting struct {
	KafkaBrokers    string
	OktaDomain      string
	Port            string
	MaximumChannels int32
	UploadPath      string
	jwtsetting
}

func LoadSettings() Setting {
	var appSettings = Setting{}
	viper.SetConfigName("appsettings")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	viper.AutomaticEnv()

	//appSettings := &setting{}
	appSettings.KafkaBrokers = viper.GetString("KafkaBrokers")
	appSettings.OktaDomain = viper.GetString("OktaDomain")
	appSettings.ValidIssuer = viper.GetString("JwtSettings.ValidIssuer")
	appSettings.ValidAudiences = viper.GetString("JwtSettings.ValidAudiences")
	appSettings.PlatformServiceBaseUrl = viper.GetString("JwtSettings.PlatformServiceBaseUrl")
	appSettings.ClientId = viper.GetString("JwtSettings.ClientId")
	appSettings.ClientSecret = viper.GetString("JwtSettings.ClientSecret")
	appSettings.MaximumChannels = viper.GetInt32("MaximumChannels")
	appSettings.Port = viper.GetString("Port")
	appSettings.UploadPath = viper.GetString("UploadPath")
	//appSettings.AccessKey = viper.GetString("AccessKey")
	//appSettings.AccessSecret = viper.GetString("AccessSecret")
	//appSettings.Region = viper.GetString("Region")
	//appSettings.S3Bucket = viper.GetString("S3Bucket")
	return appSettings
}
