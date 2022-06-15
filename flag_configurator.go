package naga

import (
	"errors"

	pflag "github.com/spf13/pflag"
)

const (
	//ErrorType corresponde a un error al ocupar el servicio ResolveType
	ErrorType = "Error al determinar el tipo de dato "
	//ErrorConfig corresponde a un error configurando la bandera
	ErrorConfig = "Error configurando el flag"
	//ErrorStruct Error en la estructira del ConfigEntry
	ErrorStruct = "Error en la estructura del ConfigEntry"
	//ErrorInt corresponde a un error convirtiendo a entero el value
	ErrorInt = "Error convirtiendo a entero"
	//ErrorConvert corresponde a un error en la conversion de la interface
	ErrorConvert = "Error convirtiendo el tipo de archivo"
)

//FlagConfigurator Interfaz que define configuración de flags
type FlagConfigurator interface {
	ConfigureFlag(ConfigEntry) error
}

//ConfigInit variable de configuracion por defecto
type flagConfigurator struct {
	variableTypeResolver VariableTypeResolver
}

//NewFlagConfigurator constructor de la implementación de la interfaz
func NewFlagConfigurator(typeResolver VariableTypeResolver) FlagConfigurator {

	return &flagConfigurator{
		variableTypeResolver: typeResolver,
	}
}

//ConfigureFlag recibe una structura ConfigEntry y ocupando el servicio de tipo configura un flag
func (f *flagConfigurator) ConfigureFlag(config ConfigEntry) error {

	//Interfaz Value del flag
	configValue := config.DefaultValue

	if len(config.VariableName) < 1 || len(config.Description) < 1 {
		return errors.New(ErrorStruct)
	}
	//Ocupo el servicio ResolveType para determinar el tipo de la interface
	variableType, err := f.variableTypeResolver.ResolveType(configValue)

	if err != nil || variableType == TypeNone {
		return errors.New(ErrorConvert)
	}

	//SWITCH PRINCIPAL
	//Configuro flag segun tipo de variable otorgada por ResolveType

	switch variableType {

	case TypeInt:

		if intVal, ok := configValue.(int); ok {

			pflag.IntP(config.VariableName, config.Shortcut, intVal, config.Description)

		}

	case TypeBool:
		if boolVal, ok := configValue.(bool); ok {

			pflag.BoolP(config.VariableName, config.Shortcut, boolVal, config.Description)
		}

	case TypeString:

		if stringVal, ok := configValue.(string); ok {

			pflag.StringP(config.VariableName, config.Shortcut, stringVal, config.Description)
		}

	}

	return nil
}
