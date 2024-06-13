package naga

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	configError = "Error reading ConfigFile"

	flagError = "Error with FlagConfigurator"

	noEntriesError = "No entries given"

	noFilenameError = "No filename given"

	noExtensionError = "No extension given"

	valueError = "No value given to the variable through config file, enviroment variables or flag"

	typeResolverError = "Error on ResolveType"
)

// ConfigEntry entrada de configuración
type ConfigEntry struct {
	VariableName string
	Description  string
	Shortcut     string
	DefaultValue interface{}
}

// CfgBase configuración basica de la API
type CfgBase struct {
	Port           string
	LoggingLevel   string
	Timeout        time.Duration
	EnabledTracing bool
	EnabledMetrics bool
	URIPrefix      string
}

// Configurator Interfaz configurador de Naga.
type Configurator interface {
	Configure(configFileName string, extension string, entries []ConfigEntry) (map[string]interface{}, error)
}

// configurator type
type configurator struct {
	flagConfigurator FlagConfigurator
	typeResolver     VariableTypeResolver
}

// NewConfigurator constructor
func NewConfigurator(flagConfigurator FlagConfigurator, typeResolver VariableTypeResolver) Configurator {
	return &configurator{
		flagConfigurator: flagConfigurator,
		typeResolver:     typeResolver,
	}
}

// Configure method
func (c *configurator) Configure(configFileName string, extension string, entries []ConfigEntry) (map[string]interface{}, error) {

	if len(entries) == 0 {
		return nil, errors.New(noEntriesError)
	}

	// Establecer valores predeterminados si no se proporcionan configFileName o extension
	if len(configFileName) == 0 {
		configFileName = "config"
	}
	if len(extension) == 0 {
		extension = "env"
	}

	// Configuration for file with specific extensions
	if len(extension) > 0 {
		if err := configFileConfiguration(configFileName, extension); err != nil {
			return nil, err
		}
	}

	// Configuration for environmental variables on the system
	viper.AutomaticEnv()

	//Configuration for flags
	if err := flagConfiguration(entries, c.flagConfigurator); err != nil {
		return nil, err
	}

	values := make(map[string]interface{})

	for _, entry := range entries {
		valType, err := validateEntry(entry, c.typeResolver)
		if err != nil {
			return nil, errors.New(typeResolverError)
		}

		name := entry.VariableName
		val := viperConfiguration(name, entry, valType)

		switch valType {
		case TypeInt:
			values[name] = val.(int)
		case TypeBool:
			values[name] = val.(bool)
		case TypeString:
			values[name] = val.(string)
		}
	}
	res := MapValues(values, entries)
	return res, nil
}

/*
Validates if the entry is accepted by the typeResolver
Returns type and an error
*/
func validateEntry(entry ConfigEntry, typeResolver VariableTypeResolver) (VariableType, error) {
	val, err := typeResolver.ResolveType(entry.DefaultValue)

	if err != nil {
		return TypeNone, errors.New(typeResolverError)
	}
	return val, nil
}

/*
Sets the value by getting the viper value
If there's no viper value, sets the default value
*/
func viperConfiguration(name string, entry ConfigEntry, valType VariableType) interface{} {
	val := viper.Get(name)
	if val == nil {
		val = entry.DefaultValue
	}

	if strType := reflect.TypeOf(val).String(); valType != TypeString && strType == "string" {
		strValue := val.(string)
		if intValue, err := strconv.Atoi(strValue); err == nil {
			return intValue
		}
		if boolValue, err := strconv.ParseBool(strValue); err == nil {
			return boolValue
		}
	}

	//checks for bool/int values that are set as string by enviroment variables
	return val
}

/*
Uses viper to read the yaml
*/
func configFileConfiguration(filename string, extension string) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")
	viper.SetConfigType(extension)
	if err := viper.ReadInConfig(); err != nil {
		return errors.New(configError + " " + err.Error())
	}
	return nil
}

/*
Sets flags using flagConfigurator
*/
func flagConfiguration(entries []ConfigEntry, flagConfigurator FlagConfigurator) error {
	for _, entry := range entries {
		if err := flagConfigurator.ConfigureFlag(entry); err != nil {
			return errors.New(flagError)
		}
	}
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	return nil
}
