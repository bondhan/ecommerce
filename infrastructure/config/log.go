package config

const (
	Production = "production"
)

// NewLogConf will return the production status and mapping for logging
func NewLogConf(env string, appName string) (bool, map[string]interface{}) {
	isProd := false
	if env == Production {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = env
	m["service"] = appName

	return isProd, m
}
