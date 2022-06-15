# Naga

![](https://www.ecured.cu/images/thumb/0/0f/Lady_vashj-wow.jpg/260px-Lady_vashj-wow.jpg)

Es un wrapper de viper que permite configurar una lista de variables de entrada, estableciendo la siguiente prioridad:

*  Flags
*  Variables de entorno
*  Archivos .yaml

Esta libreria recibe un **string** con el nombre del archivo yaml y una lista de estructuras de tipo ConfigEntry.

```
type ConfigEntry struct {
    VariableName string
    Description  string
    Shortcut     string
    DefaultValue interface{}
}
```

Esta estructura contiene los datos de las variables de configuración y se modifica cada uno de los DefaultValue según la configuración que se utilice. 

Una vez configurado, se retorna un map[string]interface{} con las llaves/valores configurados y un error. 

## Cómo usar

Primero descargamos la libreria y sus dependencias utilizando el comando:


```

go get "github.com/validatecl/commons/naga"


```

Para utilizar naga se debe importar la librería:

```
import (
    "github.com/validatecl/commons/naga"
)

```


Luego crear un VariableTypeResolver y un FlagConfigurator:

```
typeResolver := naga.NewVariableTypeResolver()
```

```
flagConfigurator := naga.NewFlagConfigurator(typeResolver)
```

Una vez creadas las estructuras, creamos un Configurator y lo ejecutamos.

```
configurator := naga.NewConfigurator(flagConfigurator, typeResolver)
values, err := configurator.Configure("filename", entries)
```
En donde *filename* es el nombre del archivo, en caso de no existir, ingresar un string vacío. Y *entries* es la lista de ConfigEntry.
Es importante señalar que los nombres en variable name de la configuración basica deben ser exactamente los mismos.

Los parametros obligatorios son: *port*, *tracing_enabled*, *logging_level*, *metrics_enabled*, *timeout*, *uri_prefix* y estos deben respetar esa sintaxis.

A continuación un ejemplo:


```
entries := []naga.ConfigEntry{
		{
			VariableName: "port",
			Description:  "Puerto a utilizar",
			Shortcut:     "p",
			DefaultValue: ":8080",
		},
		{
			VariableName: "logging_level",
			Description:  "Level de detalle de logs",
			Shortcut:     "l",
			DefaultValue: "info",
		},
		{
			VariableName: "tracing_enabled",
			Description:  "Especifica si se debe configurar tracing",
			Shortcut:     "t",
			DefaultValue: false,
		},
		{
			VariableName: "metrics_enabled",
			Description:  "Especifica si se debe configurar metrics",
			Shortcut:     "m",
			DefaultValue: false,
		},
		
		{
			VariableName: "timeout",
			Description:  "timeout por defecto ",
			Shortcut:     "",
			DefaultValue: 30,
		},
		{
			VariableName: "uri_prefix",
			Description:  "Prefijo de URL con version",
			DefaultValue: "/uri_prefix/v1",
		},
	}
```




