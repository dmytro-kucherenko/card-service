package config

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func bind(schema any) {
	schemaType := reflect.TypeOf(schema).Elem()
	for i := 0; i < schemaType.NumField(); i++ {
		field := schemaType.Field(i)
		if tag, ok := field.Tag.Lookup("mapstructure"); ok {
			viper.BindEnv(tag)
		}
	}
}

func load(path string, schema any) error {
	godotenv.Load(path)
	bind(schema)

	err := viper.Unmarshal(schema)
	if err != nil {
		return err
	}

	validate := validator.New()

	return validate.Struct(schema)
}
