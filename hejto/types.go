package hejto

import (
	"net/url"
	"strconv"
	"time"

	"github.com/wombatDaiquiri/lajko/database"
	"github.com/wombatDaiquiri/lajko/ee/slices"
)

type OrderingDirection string

const (
	ASC  OrderingDirection = "asc"
	DESC OrderingDirection = "desc"
)

type PostOrdering string

const (
	PostOrderingCreatedAt   PostOrdering = "p.createdAt"
	PostOrderingLikes       PostOrdering = "p.numLikes"
	PostOrderingNumComments PostOrdering = "p.numComments"
	PostOrderingHot         PostOrdering = "p.hot"
	PostOrderingHotness     PostOrdering = "p.hotness"
	PostOrderingRand        PostOrdering = "rand"
	PostOrderingPromoted    PostOrdering = "p.promoted"
)

type PostPagination struct {
	Page     int
	Limit    int
	OrderBy  PostOrdering
	OrderDir OrderingDirection

	FromAuthor string
}

func (pp PostPagination) Query() string {
	query := url.Values{}
	if pp.Page < 1 {
		query.Set("page", "1")
	} else {
		query.Set("page", strconv.Itoa(pp.Page))
	}

	// TODO: Actual query

	return query.Encode()
}

type embedded[T any] struct {
	Items []T `json:"items"`
}

type contentLink struct {
	URL   string `json:"url"`
	Site  string `json:"site"`
	Type  string `json:"type"`
	Title string `json:"title"`
	// ???? TODO
	Audios []interface{} `json:"audios"`
	Images []struct {
		Url  string `json:"url"`
		Safe string `json:"safe"`
	} `json:"images"`
	// ???? TODO
	Videos []interface{} `json:"videos"`
	// Favicon struct {
	// 	URL  string `json:"url"`
	// 	Safe string `json:"safe"`
	// } `json:"favicon"`
	Description string `json:"description"`
}

type community struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	PrimaryColor string `json:"primary_color"`
}

type author struct {
	// technical
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Links     struct {
		Self    link `json:"self"`
		Follows link `json:"follows"`
	} `json:"_links"`

	// display
	Avatar struct {
		URLs struct {
			Tiny   string `json:"100x100"`
			Medium string `json:"250x250"`
		} `json:"urls"`
		UUID string `json:"uuid"`
	} `json:"avatar"`
	CurrentRank  string `json:"current_rank"`
	CurrentColor string `json:"current_color"`
	Userbar      struct {
		URL struct {
			Original string `json:"original"`
		} `json:"urls"`
		UUID string `json:"uuid"`
	} `json:"userbar"`

	// flags
	Status        string   `json:"status"`
	Roles         []string `json:"roles"`
	Controversial bool     `json:"controversial"`
	Verified      bool     `json:"verified"`
	Sponsor       bool     `json:"sponsor"`
	ExSponsor     bool     `json:"ex_sponsor"`
}

type tag struct {
	// the important thing
	Name  string `json:"name"`
	Links struct {
		Self    link `json:"self"`
		Follows link `json:"follows"`
		Blocks  link `json:"blocks"`
	} `json:"_links"`

	// scores
	NumFollows int `json:"num_follows"`
	NumPosts   int `json:"num_posts"`
}

type comment struct {
	Content      string        `json:"content"`
	ContentPlain string        `json:"content_plain"`
	Images       []interface{} `json:"images"`
	ContentLinks []interface{} `json:"content_links"`
	PostSlug     string        `json:"post_slug"`
	Status       string        `json:"status"`
	Author       author        `json:"author"`
	NumLikes     int           `json:"num_likes"`
	NumReports   int           `json:"num_reports"`
	IsLiked      bool          `json:"is_liked"`
	IsReported   bool          `json:"is_reported"`
	Replies      []interface{} `json:"replies"`
	LikesEnabled bool          `json:"likes_enabled"`
	CreatedAt    time.Time     `json:"created_at"`
	Uuid         string        `json:"uuid"`
	Links        struct {
		Self  link `json:"self"`
		Likes link `json:"likes"`
	} `json:"_links"`
}

type post struct {
	// technical fields
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	Slug      string    `json:"slug"`
	// this is actually an enum
	Status    string    `json:"status"`
	Hot       bool      `json:"hot"`
	Community community `json:"community"`
	Author    author    `json:"author"`
	Tags      []tag     `json:"tags"`

	// post content
	Title        string `json:"title"`
	Content      string `json:"content"`
	ContentPlain string `json:"content_plain"`
	Excerpt      string `json:"excerpt"`

	Images       []image       `json:"images"`
	ContentLinks []contentLink `json:"content_links"`
	Comments     []comment     `json:"comments"`

	// content flags
	NSFW          bool `json:"nsfw"`
	Controversial bool `json:"controversial"`
	WarContent    bool `json:"war_content"`
	Masked        bool `json:"masked"`
	// stats
	NumLikes     int `json:"num_likes"`
	NumComments  int `json:"num_comments"`
	NumFavorites int `json:"num_favorites"`
	// personal
	IsLiked     bool `json:"is_liked"`
	IsCommented bool `json:"is_commented"`
	IsFavorited bool `json:"is_favorited"`
	// looks kinda randomy/debuggy
	CommentsEnabled  bool `json:"comments_enabled"`
	FavoritesEnabled bool `json:"favorites_enabled"`
	LikesEnabled     bool `json:"likes_enabled"`
	ReportsEnabled   bool `json:"reports_enabled"`
	SharesEnabled    bool `json:"shares_enabled"`
	// no idea
	Discr string `json:"discr"`
}

type image struct {
	URLs     urls   `json:"urls"`
	UUID     string `json:"uuid"`
	Position int    `json:"position"`
}

type urls struct {
	Small  string `json:"250x250"`
	Medium string `json:"500x500"`
	Large  string `json:"1200x900"`
}

type postResponse struct {
	pagination
	Links    links          `json:"_links"`
	Embedded embedded[post] `json:"_embedded"`
}

func (ps postResponse) DatabasePosts() []database.Post {
	return slices.Map(ps.Embedded.Items, func(p post) database.Post {
		return database.Post{
			Model: database.Model{},
			Source: database.Source{
				Source: database.SourcePortalHejto,
				// TODO
				SourceData:      nil,
				SourceID:        "",
				SourceURL:       "",
				SourceCreatedAt: time.Time{},
			},
			Title: database.TextContent{
				Markdown: p.Title,
			},
			Slug: database.TextContent{
				Markdown: p.Slug,
			},
			Author: database.Author{
				Model:             database.Model{},
				Source:            database.Source{},
				Username:          p.Author.Username,
				DisplayedUsername: p.Author.Username,
				AvatarURL:         p.Author.Avatar.URLs.Medium,
			},
			Content: database.Content{
				TextContent: database.TextContent{Markdown: p.Content},
			},
			Likes: int64(p.NumLikes),
			// TODO
			Comments: nil,
			// TODO
			Tags: nil,
		}
	})
}

type pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	// how do they differ? also it's actually 50 pages xD
	Pages int `json:"pages"`
	Total int `json:"total"`
}

type links struct {
	Self  link `json:"self"`
	First link `json:"first"`
	Last  link `json:"last"`
	Next  link `json:"next"`
}

type link struct {
	HREF string `json:"href"`
}
