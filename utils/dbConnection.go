package utils

import "fmt"

func DbConnectionString() string {

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", ReadValue("DBHOST"), ReadValue("DBPORT"), ReadValue("DBUSERNAME"), ReadValue("DBNAME"), ReadValue("DBPASSWORD"), ReadValue("DBSSL"))
}
