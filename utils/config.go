package utils

import(
    "os"
    "fmt"
    "github.com/joho/godotenv"
)

type Config struct {
    DefaultDBHost     string
    DefaultDBUser     string
    DefaultDBPassword string
    DefaultDBName     string
    DefaultDBPort     string

    BillingDBHost     string
    BillingDBUser     string
    BillingDBPassword string
    BillingDBName     string
    BillingDBPort     string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
	if(err != nil){
		fmt.Println("Not founded env file");
	}

    // Чтение переменных окружения
    config := &Config{
        BillingDBHost:     getEnv("DB_HOST", "172.18.0.7"),
        BillingDBUser:     getEnv("DB_PORT", "5432"),
        BillingDBPassword: getEnv("DB_DATABASE", "billing"),
        BillingDBName:     getEnv("DB_USERNAME", "db_user"),
        BillingDBPort:     getEnv("DB_PASSWORD", "dbpass"),

        DefaultDBHost:     getEnv("COMMON_DB_HOST", "172.18.0.7"),
        DefaultDBUser:     getEnv("COMMON_DB_PORT", "postgres"),
        DefaultDBPassword: getEnv("COMMON_DB_DATABASE", "defaultdb"),
        DefaultDBName:     getEnv("COMMON_DB_USERNAME", "doadmin"),
        DefaultDBPort:     getEnv("COMMON_DB_PASSWORD", "dbpass"),
    }

    return config, err
}


func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return fallback
    }
    return value
}