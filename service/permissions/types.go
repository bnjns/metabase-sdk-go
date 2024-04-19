package permissions

// Group represents the details of an existing permissions group returned from the Metabase API.
type Group struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Members     []GroupMember `json:"members"`
	MemberCount int64         `json:"member_count"`
}

type GroupMember struct {
	UserId         int64  `json:"user_id"`
	GroupId        int64  `json:"group_id"`
	MembershipId   int64  `json:"membership_id"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	IsGroupManager *bool  `json:"is_group_manager"`
}

// CreateGroupRequest represents the request body used to create a new permissions group.
type CreateGroupRequest struct {
	Name string `json:"name"`
}

// UpdateGroupRequest represents the request body used to update an existing permissions group.
type UpdateGroupRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
