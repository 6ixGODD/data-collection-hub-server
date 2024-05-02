package config

import (
	"reflect"
	"strconv"
	"time"

	"data-collection-hub-server/internal/pkg/config/modules"
)

type Config struct {
	BaseConfig    modules.BaseConfig    `mapstructure:"base" yaml:"base"`
	CasbinConfig  modules.CasbinConfig  `mapstructure:"casbin" yaml:"casbin"`
	JWTConfig     modules.JWTConfig     `mapstructure:"jwt" yaml:"jwt"`
	MongoConfig   modules.MongoConfig   `mapstructure:"mongo" yaml:"mongo"`
	RedisConfig   modules.RedisConfig   `mapstructure:"redis" yaml:"redis"`
	ZapConfig     modules.ZapConfig     `mapstructure:"zap" yaml:"zap"`
	FiberConfig   modules.FiberConfig   `mapstructure:"fiber" yaml:"fiber"`
	LimiterConfig modules.LimiterConfig `mapstructure:"limiter" yaml:"limiter"`
	CorsConfig    modules.CorsConfig    `mapstructure:"cors" yaml:"cors"`
}

// NewConfig returns a new instance of Config
func NewConfig() (config *Config, err error) {
	config = &Config{}
	err = Init(config)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}

func Init(config *Config) (err error) {
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		subStructValue := configValue.Field(i)
		subStructType := subStructValue.Type()

		if subStructValue.Kind() == reflect.Struct {
			for j := 0; j < subStructType.NumField(); j++ {
				subField := subStructType.Field(j)
				subFieldValue := subStructValue.Field(j)

				if subFieldValue.CanSet() {
					defaultTag := subField.Tag.Get("default")
					if defaultTag != "" {
						if err := setFieldWithDefault(subFieldValue, defaultTag); err != nil {
							return err
						}
					}
				}
			}
		} else {
			defaultTag := subStructType.Field(0).Tag.Get("default")
			if defaultTag != "" {
				if err := setFieldWithDefault(subStructValue, defaultTag); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func setFieldWithDefault(field reflect.Value, defaultTag string) (err error) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(defaultTag)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v, err := strconv.ParseInt(defaultTag, 10, field.Type().Bits()); err == nil {
			field.SetInt(v)
		} else if d, err := time.ParseDuration(defaultTag); err == nil {
			field.SetInt(int64(d))
		} else {
			return err
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v, err := strconv.ParseUint(defaultTag, 10, field.Type().Bits()); err == nil {
			field.SetUint(v)
		} else {
			return err
		}
	case reflect.Float32, reflect.Float64:
		if v, err := strconv.ParseFloat(defaultTag, field.Type().Bits()); err == nil {
			field.SetFloat(v)
		} else {
			return err
		}
	case reflect.Bool:
		if v, err := strconv.ParseBool(defaultTag); err == nil {
			field.SetBool(v)
		} else {
			return err
		}
	case reflect.Array, reflect.Slice:
		switch field.Type().Elem().Kind() {
		case reflect.String:
			field.Set(reflect.ValueOf([]string{defaultTag}))
		case reflect.Int:
			if v, err := strconv.ParseInt(defaultTag, 10, field.Type().Elem().Bits()); err == nil {
				field.Set(reflect.ValueOf([]int{int(v)}))
			} else {
				return err
			}
		case reflect.Uint:
			if v, err := strconv.ParseUint(defaultTag, 10, field.Type().Elem().Bits()); err == nil {
				field.Set(reflect.ValueOf([]uint{uint(v)}))
			} else {
				return err
			}
		case reflect.Float64:
			if v, err := strconv.ParseFloat(defaultTag, field.Type().Elem().Bits()); err == nil {
				field.Set(reflect.ValueOf([]float64{v}))
			} else {
				return err
			}
		case reflect.Bool:
			if v, err := strconv.ParseBool(defaultTag); err == nil {
				field.Set(reflect.ValueOf([]bool{v}))
			} else {
				return err
			}
		default:
			field.SetString(defaultTag)
		}
	default:
		field.SetString(defaultTag)
	}
	return nil
}
