package database

const (
	EngineAmazonAthena    Engine = "athena"
	EngineAmazonRedshift  Engine = "redshift"
	EngineBigQuery        Engine = "bigquery"
	EngineDruid           Engine = "druid"
	EngineGoogleAnalytics Engine = "googleanalytics"
	EngineMongoDB         Engine = "mongo"
	EngineMySQL           Engine = "mysql"
	EnginePostgres        Engine = "postgres"
	EnginePresto          Engine = "presto-jdbc"
	EngineSnowflake       Engine = "snowflake"
	EngineSparkSQL        Engine = "sparksql"
	EngineSQLServer       Engine = "sqlserver"
	EngineSQLite          Engine = "sqlite"
)

const (
	ScheduleTypeHourly  ScheduleType = "hourly"
	ScheduleTypeDaily   ScheduleType = "daily"
	ScheduleTypeWeekly  ScheduleType = "weekly"
	ScheduleTypeMonthly ScheduleType = "monthly"
)

const (
	ScheduleDayTypeSun ScheduleDayType = "sun"
	ScheduleDayTypeMon ScheduleDayType = "mon"
	ScheduleDayTypeTue ScheduleDayType = "tue"
	ScheduleDayTypeWed ScheduleDayType = "wed"
	ScheduleDayTypeThu ScheduleDayType = "thu"
	ScheduleDayTypeFri ScheduleDayType = "fri"
	ScheduleDayTypeSat ScheduleDayType = "sat"
)

const (
	ScheduleFrameTypeFirst ScheduleFrameType = "first"
	ScheduleFrameTypeMid   ScheduleFrameType = "mid"
	ScheduleFrameTypeLast  ScheduleFrameType = "last"
)
