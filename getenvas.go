package getenvas

import (
	"os"
	"strconv"
	"time"
)

// Retrieve a string value from an environment variable or fall-back to a default value if not set
// String values don't require conversion so you'll never receive any errors
// If the environment variable doesn't exist, the default value is returned
func String(key string, default_value string) string {
	if value, present := os.LookupEnv(key); present {
		return value
	}
	return default_value
}

// Retrieve a string value from an environment variable or fall-back to a default value if not set
// String values don't require conversion so you'll never receive any errors
// If the environment variable doesn't exist, the default value is assigned
func StringVar(out *string, key string, default_value string) {
	if value, present := os.LookupEnv(key); present {
		*out = value
	} else {
		*out = default_value
	}
}

// Retrieve a boolean value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Bool(key string, default_value bool) (bool, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseBool(env_value); err != nil {
			return default_value, err
		} else {
			return value, nil
		}
	}
	return default_value, nil
}

// Retrieve a boolean value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func BoolVar(out *bool, key string, default_value bool) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseBool(env_value); err != nil {
			*out = default_value
			return err
		} else {
			*out = value
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a 32-bit integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Int(key string, default_value int) (int, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseInt(env_value, 10, 32); err != nil {
			return default_value, err
		} else {
			return int(value), nil
		}
	}
	return default_value, nil
}

// Retrieve a 32-bit integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func IntVar(out *int, key string, default_value int) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseInt(env_value, 10, 32); err != nil {
			*out = default_value
			return err
		} else {
			*out = int(value)
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a 32-bit unsigned integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Uint(key string, default_value uint) (uint, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseUint(env_value, 10, 32); err != nil {
			return default_value, err
		} else {
			return uint(value), nil
		}
	}
	return default_value, nil
}

// Retrieve a 32-bit unsigned integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func UintVar(out *uint, key string, default_value uint) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseUint(env_value, 10, 32); err != nil {
			*out = default_value
			return err
		} else {
			*out = uint(value)
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a 64-bit integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Int64(key string, default_value int64) (int64, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseInt(env_value, 10, 64); err != nil {
			return default_value, err
		} else {
			return value, nil
		}
	}
	return default_value, nil
}

// Retrieve a 64-bit integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func Int64Var(out *int64, key string, default_value int64) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseInt(env_value, 10, 64); err != nil {
			*out = default_value
			return err
		} else {
			*out = value
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a 64-bit unsigned integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Uint64(key string, default_value uint64) (uint64, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseUint(env_value, 10, 64); err != nil {
			return default_value, err
		} else {
			return value, nil
		}
	}
	return default_value, nil
}

// Retrieve a 64-bit unsigned integer value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func Uint64Var(out *uint64, key string, default_value uint64) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseUint(env_value, 10, 64); err != nil {
			*out = default_value
			return err
		} else {
			*out = value
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a 32-bit floating point value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Float32(key string, default_value float32) (float32, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseFloat(env_value, 32); err != nil {
			return default_value, err
		} else {
			return float32(value), nil
		}
	}
	return default_value, nil
}

// Retrieve a 32-bit floating point value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func Float32Var(out *float32, key string, default_value float32) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseFloat(env_value, 32); err != nil {
			*out = default_value
			return err
		} else {
			*out = float32(value)
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a 64-bit floating point value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Float64(key string, default_value float64) (float64, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseFloat(env_value, 64); err != nil {
			return default_value, err
		} else {
			return float64(value), nil
		}
	}
	return default_value, nil
}

// Retrieve a 64-bit floating point value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func Float64Var(out *float64, key string, default_value float64) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := strconv.ParseFloat(env_value, 64); err != nil {
			*out = default_value
			return err
		} else {
			*out = float64(value)
		}
	} else {
		*out = default_value
	}
	return nil
}

// Retrieve a time duration value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is returned with no error
// If a parsing error occurs, the default value is returned along with the parsing error
func Duration(key string, default_value time.Duration) (time.Duration, error) {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := time.ParseDuration(env_value); err != nil {
			return default_value, err
		} else {
			return value, nil
		}
	}
	return default_value, nil
}

// Retrieve a time duration value from an environment variable or fall-back to a default value if not set
// If the environment variable doesn't exist, the default value is assigned and no error is returned
// If a parsing error occurs, the default value is assigned and the parsing error is returned
func DurationVar(out *time.Duration, key string, default_value time.Duration) error {
	if env_value, present := os.LookupEnv(key); present {
		if value, err := time.ParseDuration(env_value); err != nil {
			*out = default_value
			return err
		} else {
			*out = value
		}
	} else {
		*out = default_value
	}
	return nil
}
