package my_model

type Comment struct {
	ID       string     `json:"id"`
	Content  string     `json:"content"`
	PostID   int        `json:"postID"`
	UserID   int        `json:"userID"`
	ParentID *int       `json:"parentID,omitempty"`
	Comments []*Comment `json:"comments" gorm:"foreignKey:PostID"`
}
