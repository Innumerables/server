package config

type Server struct {
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Local    Local    `mapstructure:"local" json:"local" yaml:"local"`
}
