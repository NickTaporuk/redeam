package configuration

//nolint
// InitEnv extract configuration data from environment
func InitEnv(data map[string]string) error {
	var err error
	err = CheckEnvVar(EnvNameDatabaseType, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(EnvNameDatabaseHost, data)
	if err != nil {
		return err
	}
	err = CheckEnvVar(EnvNameDatabaseUser, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(EnvNameDatabaseName, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(EnvNameDatabasePssCode, data)
	if err != nil {
		return err
	}
	err = CheckEnvVar(EnvNameDatabasePort, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(EnvNameDatabaseSSLMode, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(EnvNameDatabaseMigrate, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(EnvNameSeeds, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(ServiceHost, data)
	if err != nil {
		return err
	}

	err = CheckEnvVar(ServicePort, data)
	if err != nil {
		return err
	}

	return nil
}
