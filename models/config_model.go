package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "gin-skeleton/utils"
)

func GetConfig() (map[string]string, error) {
    mapConfig := make(map[string]string)
    
    config, _ := utils.ReadConfig()

    db, _ := sql.Open("sqlite3", config.DB.SQLite.DBFile)
    
    rows, err := db.Query("SELECT * FROM configurations")
    
    for rows.Next() {
        var id int
        var cfgKey string
        var cfgValue string
        err = rows.Scan(&id, &cfgKey, &cfgValue)
        mapConfig[cfgKey] = cfgValue
    }
    
    return mapConfig, err
}