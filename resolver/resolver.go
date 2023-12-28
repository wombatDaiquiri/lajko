package resolver

import (
	"context"
	"fmt"

	"github.com/wombatDaiquiri/lajko/database"
	"github.com/wombatDaiquiri/lajko/ee"
	"github.com/wombatDaiquiri/lajko/ee/slices"
	"github.com/wombatDaiquiri/lajko/hejto"
)

func New() *Resolver {
	return &Resolver{}
}

type Resolver struct{}

func (r *Resolver) Posts(ctx context.Context, args struct{ Query *PostQuery }) ([]Post, error) {
	pagination := hejto.PostPagination{
		Page:     1,
		Limit:    10,
		OrderBy:  hejto.PostOrderingHotness,
		OrderDir: hejto.DESC,
	}
	if args.Query != nil {
		if args.Query.Page != nil {
			pagination.Page = int(*args.Query.Page)
		}
		if args.Query.Limit != nil {
			pagination.Limit = int(*args.Query.Limit)
		}
		if args.Query.OrderBy != nil {
			pagination.OrderBy = hejto.PostOrdering(*args.Query.OrderBy)
		}
		if args.Query.OrderDir != nil {
			pagination.OrderDir = hejto.OrderingDirection(*args.Query.OrderDir)
		}
	}

	client := hejto.Client{}
	posts, err := client.Posts(context.Background(), pagination)
	if err != nil {
		return []Post{}, fmt.Errorf("return posts from hejto: %w", err)
	}

	fmt.Printf("number of posts returned: %v\n", len(posts))
	return newPostsResolver(posts), nil
}

func newPostsResolver(dbPosts []database.Post) []Post {
	return slices.Map(dbPosts, func(dbPost database.Post) Post {
		return Post{
			ULID:      dbPost.ULID,
			CreatedAt: int32(dbPost.CreatedAt.Unix()),
			UpdatedAt: int32(dbPost.UpdatedAt.Unix()),

			Source:          postSourceHejto,
			SourceURL:       dbPost.SourceURL,
			SourceCreatedAt: int32(dbPost.SourceCreatedAt.Unix()),

			Title: dbPost.Title.Markdown,
			// TODO: resolver for author
			// Author: dbPost.Author,
			Author: Author{},
			Slug:   dbPost.Slug.Markdown,

			Content: dbPost.Content.Markdown,

			Images: ee.EvalEager(
				len(dbPost.Attachments.Images) > 0,
				&dbPost.Attachments.Images,
				nil),
			// todo: resolvers for tags & comments
			// Tags        *[]Tag
			// Comments    *[]Comment

			Likes: int32(dbPost.Likes),
		}
	})
}

type postSource string

const postSourceHejto postSource = "HEJTO"

type direction string

const (
	directionAsc  direction = "ASC"
	directionDesc direction = "DESC"
)

type Post struct {
	ULID      string
	CreatedAt int32
	UpdatedAt int32

	Source          postSource
	SourceURL       string
	SourceCreatedAt int32

	Title  string
	Author Author
	Slug   string

	Content     string
	Images      *[]string
	Attachments *[]string
	Tags        *[]Tag
	Comments    *[]Comment

	Likes int32
}

type Author struct {
	ULID      string
	CreatedAt int32
	UpdatedAt int32

	Source          postSource
	SourceURL       string
	SourceCreatedAt int32

	Username          string
	DisplayedUsername string
	AvatarURL         string
	Description       string
}

type Tag struct {
	ULID      string
	CreatedAt int32
	UpdatedAt int32

	Name string
	URLs *[]string
}

type Comment struct {
	ULID      string
	CreatedAt int32
	UpdatedAt int32

	Author Author

	Content     string
	Images      *[]string
	Attachments *[]string
	Tags        *[]Tag

	Likes int32
}

type PostQuery struct {
	Source   *postSource
	Page     *int32
	Limit    *int32
	OrderBy  *string
	OrderDir *direction
}