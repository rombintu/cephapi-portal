package utils

import "os"

func GetEnvOfDefault(envKey, defaultEnvVar string) string {
	envVar, exist := os.LookupEnv(envKey)
	if exist {
		return envVar
	} else {
		return defaultEnvVar
	}
}
