package database

// The Engine is the type of database (eg, MySQL, PostgreSQL, etc)
type Engine string

// Database represents the details of an existing database returned from the Metabase API.
type Database struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Engine Engine `json:"engine"`

	CreatedAt string  `json:"created_at"`
	CreatorId *int64  `json:"creator_id"`
	UpdatedAt *string `json:"updated_at"`

	AutoRunQueries   bool `json:"auto_run_queries"`
	CanUpload        bool `json:"can_upload"`
	CanManage        bool `json:"can-manage"`
	IsFullSync       bool `json:"is_full_sync"`
	IsSample         bool `json:"is_sample"`
	IsOnDemand       bool `json:"is_on_demand"`
	Refingerprint    bool `json:"refingerprint"`
	IsSavedQuestions bool `json:"is_saved_questions"`

	Caveats           *string            `json:"caveats"`
	Features          []string           `json:"features"`
	Details           *Details           `json:"details"`
	Settings          *map[string]string `json:"settings"`
	Schedules         *Schedules         `json:"schedules"`
	Timezone          string             `json:"timezone"`
	NativePermissions string             `json:"native_permissions"`
	PointsOfInterest  *string            `json:"points_of_interest"`

	CacheTTL          *string `json:"cache_ttl"`
	InitialSyncStatus string  `json:"initial_sync_status"`
}

type Details map[string]interface{}

type Schedules struct {
	MetadataSync     *ScheduleSettings `json:"metadata_sync"`
	CacheFieldValues *ScheduleSettings `json:"cache_field_values"`
}
type ScheduleType string
type ScheduleDayType string
type ScheduleFrameType string
type ScheduleSettings struct {
	Type   ScheduleType       `json:"schedule_type"`
	Day    *ScheduleDayType   `json:"schedule_day"`
	Frame  *ScheduleFrameType `json:"schedule_frame"`
	Hour   *int64             `json:"schedule_hour"`
	Minute *int64             `json:"schedule_minute"`
}

// CreateRequest represents the request body used to add a new database configuration.
type CreateRequest struct {
	Name             string     `json:"name"`
	Engine           Engine     `json:"engine"`
	Details          Details    `json:"details"`
	IsFullSync       *bool      `json:"is_full_sync"`
	IsOnDemand       *bool      `json:"is_on_demand"`
	Schedules        *Schedules `json:"schedules"`
	AutoRunQueries   *bool      `json:"auto_run_queries"`
	CacheTTL         *int64     `json:"cache_ttl"`
	ConnectionSource *string    `json:"connection_source"`
}

// UpdateRequest represents the request body used to update an existing database configuration.
type UpdateRequest struct {
	Id               int64              `json:"id"`
	Name             *string            `json:"name"`
	Engine           *Engine            `json:"engine"`
	Refingerprint    *bool              `json:"refingerprint"`
	Details          *Details           `json:"details"`
	Schedules        *Schedules         `json:"schedules"`
	Description      *string            `json:"description"`
	Caveats          *string            `json:"caveats"`
	PointsOfInterest *string            `json:"points_of_interest"`
	AutoRunQueries   *bool              `json:"auto_run_queries"`
	CacheTTL         *int64             `json:"cache_ttl"`
	Settings         *map[string]string `json:"settings"`
}
