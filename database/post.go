package database

type Post struct {
	Model
	Source

	Title TextContent `prefix:"title_"`
	Slug  TextContent `prefix:"slug_"`
	// many-to-1
	Author Author

	Content
	Likes int64
	// 1-to-many
	Comments []Comment
	// many-to-many
	Tags []Tag
}

type Author struct {
	Model
	Source

	Username          string
	DisplayedUsername string
	AvatarURL         string
	Content
}

type Tag struct {
	Model
	SourceData map[SourcePortal]Content

	Name string
	URLs []string
}

type Comment struct {
	Model
	Source

	Author
	Content

	Likes int64

	ReferencesPost    *Post
	ReferencesComment *Comment
}
