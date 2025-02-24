package biz

import (
	"context"
	"fmt"

	"gorm.io/gen"
	"gorm.io/gen/examples/dal/query"
)

func FullTextSearchExample(ctx context.Context) {
	u := query.User

	// 基本搜索
	users, err := u.WithContext(ctx).
		TSQuery("web & developer", u.Bio, u.Description).
		Find()
	if err != nil {
		fmt.Printf("basic search error: %v\n", err)
		return
	}
	fmt.Printf("found %d users matching 'web & developer'\n", len(users))

	// 带权重的搜索
	users, err = u.WithContext(ctx).
		WithTSConfig(gen.TSConfig{
			Language: "english",
			Weights:  []string{"A", "B"}, // Bio权重A, Description权重B
		}).
		TSQuery("web & developer", u.Bio, u.Description).
		TSRank("web & developer", u.Bio, u.Description).
		Find()
	if err != nil {
		fmt.Printf("weighted search error: %v\n", err)
		return
	}
	fmt.Printf("found %d users ordered by relevance\n", len(users))

	// 获取匹配的文本片段
	type Result struct {
		*query.User
		Headline string
	}
	var results []Result
	err = u.WithContext(ctx).
		WithTSConfig(gen.TSConfig{Language: "english"}).
		TSQuery("web & developer", u.Bio).
		TSHeadline("web & developer", u.Bio).
		Find(&results)
	if err != nil {
		fmt.Printf("headline search error: %v\n", err)
		return
	}
	for _, r := range results {
		fmt.Printf("User: %s\nMatched Text: %s\n\n", r.Name, r.Headline)
	}
} 