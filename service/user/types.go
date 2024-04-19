package user

type SSOSource string

type LoginAttributes map[string]string

// User represents the details of an existing user returned from the Metabase API.
type User struct {
	Id         int64   `json:"id"`
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	CommonName *string `json:"common_name"`
	Email      string  `json:"email"`
	Locale     *string `json:"locale"`

	IsActive    bool  `json:"is_active"`
	IsQbnewb    bool  `json:"is_qbnewb"`
	IsSuperuser bool  `json:"is_superuser"`
	IsInstaller *bool `json:"is_installer"`

	LoginAttributes  *LoginAttributes  `json:"login_attributes"`
	GroupMemberships []GroupMembership `json:"user_group_memberships"`

	GoogleAuth bool       `json:"google_auth"`
	SSOSource  *SSOSource `json:"sso_source"`

	HasInvitedSecondUser    bool        `json:"has_invited_second_user"`
	HasQuestionAndDashboard bool        `json:"has_question_and_dashboard"`
	PersonalCollectionId    interface{} `json:"personal_collection_id"`

	DateJoined string  `json:"date_joined"`
	FirstLogin *string `json:"first_login"`
	LastLogin  *string `json:"last_login"`
	UpdatedAt  *string `json:"updated_at"`
}

type GroupMembership struct {
	Id             int64 `json:"id"`
	IsGroupManager bool  `json:"is_group_manager"`
}

type currentUser struct {
	User
	GroupIds []int64 `json:"group_ids"`
}

// CreateRequest represents the request body used to create a new user.
type CreateRequest struct {
	Email            string             `json:"email"`
	FirstName        *string            `json:"first_name"`
	LastName         *string            `json:"last_name"`
	GroupMemberships *[]GroupMembership `json:"user_group_memberships"`
	LoginAttributes  *LoginAttributes   `json:"login_attributes"`
}

// UpdateRequest represents the request body to update an existing user.
type UpdateRequest struct {
	Id               int64              `json:"id"`
	Email            *string            `json:"email"`
	FirstName        *string            `json:"first_name"`
	LastName         *string            `json:"last_name"`
	Locale           *string            `json:"locale"`
	IsGroupManager   *bool              `json:"is_group_manager"`
	IsSuperuser      *bool              `json:"is_superuser"`
	LoginAttributes  *LoginAttributes   `json:"login_attributes"`
	GroupMemberships *[]GroupMembership `json:"user_group_memberships"`
}
