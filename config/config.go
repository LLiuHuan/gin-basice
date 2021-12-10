package config

type Server struct {
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	RateLimit RateLimit `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// oss
	Local      LocalOSS   `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      QiniuOSS   `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
}