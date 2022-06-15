package naga

import "time"

//MapValues validate map
func MapValues(values map[string]interface{}, entries []ConfigEntry) map[string]interface{} {
	if values == nil || entries == nil {
		panic("config not exists")
	}
	for _, entry := range entries {
		_, exists := values[entry.VariableName]
		if !exists {
			panic(entry.VariableName + " config not exists")
		}

	}
	return values
}

//GetBaseCfg lee y configura la base de variables de entorno base de la api
func GetBaseCfg(cfg map[string]interface{}) *CfgBase {

	return &CfgBase{

		Port:           cfg["port"].(string),
		EnabledMetrics: cfg["metrics_enabled"].(bool),
		LoggingLevel:   cfg["logging_level"].(string),
		Timeout:        time.Duration(cfg["timeout"].(int)) * time.Second,
		EnabledTracing: cfg["tracing_enabled"].(bool),
		URIPrefix:      cfg["uri_prefix"].(string),
	}

}
