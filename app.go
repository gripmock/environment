package environment

type App struct {
	// Application logging level.
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	// Ignore the service package.
	Simpler bool `envconfig:"SERVICE_SIMPLER" default:"true"`

	// Logical matching mode.
	// Since version 3.x, all conditions are checked together, which adds more flexibility in sampling.
	// To disable functionality, turn this option off.
	Matcher bool `envconfig:"MATCHER" default:"false"`
}
