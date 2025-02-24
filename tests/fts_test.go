package tests_test

import (
	"testing"
)

func TestFullTextSearch(t *testing.T) {
	//u := query.User
	//ctx := context.Background()
	//
	//// 准备测试数据
	//testUsers := []*model.User{
	//	{Name: "John", Bio: "Senior web developer with 10 years experience"},
	//	{Name: "Alice", Bio: "Full stack developer and DevOps engineer"},
	//	{Name: "Bob", Bio: "Mobile app developer"},
	//	{Name: "Carol", Description: "Web developer and UI/UX designer"},
	//}
	//err := u.WithContext(ctx).Create(testUsers)
	//assert.NoError(t, err)
	//
	//t.Run("Basic Search", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		TSQuery("web & developer", u.Bio).
	//		Find()
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, 2, len(users))
	//	assert.Contains(t, users[0].Bio, "web developer")
	//})
	//
	//t.Run("Weighted Search", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{
	//			Language: "english",
	//			Weights:  []string{"A", "B"},
	//		}).
	//		TSQuery("web & developer", u.Bio, u.Description).
	//		TSRank("web & developer", u.Bio, u.Description).
	//		Find()
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, 3, len(users))
	//	// Bio有"web developer"的应该排在前面(因为Bio的权重是A)
	//	assert.Equal(t, "John", users[0].Name)
	//})
	//
	//t.Run("Headline Search", func(t *testing.T) {
	//	type Result struct {
	//		*model.User
	//		Headline string
	//	}
	//	var results []Result
	//	err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{Language: "english"}).
	//		TSQuery("web & developer", u.Bio).
	//		TSHeadline("web & developer", u.Bio).
	//		Find(&results)
	//
	//	assert.NoError(t, err)
	//	assert.True(t, len(results) > 0)
	//	// 检查高亮标记
	//	assert.Contains(t, results[0].Headline, "<b>")
	//	assert.Contains(t, results[0].Headline, "</b>")
	//})
	//
	//t.Run("Different Language", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{Language: "simple"}).
	//		TSQuery("developer", u.Bio).
	//		Find()
	//
	//	assert.NoError(t, err)
	//	assert.True(t, len(users) > 0)
	//})
	//
	//t.Run("Empty Query", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		TSQuery("", u.Bio).
	//		Find()
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, 0, len(users))
	//})
	//
	//t.Run("No Columns", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		TSQuery("web").
	//		Find()
	//
	//	assert.NoError(t, err)
	//	assert.Equal(t, 0, len(users))
	//})
	//
	//t.Run("Complex Query", func(t *testing.T) {
	//	// 测试复杂查询语法
	//	users, err := u.WithContext(ctx).
	//		TSQuery("web & (developer | engineer) & !mobile", u.Bio).
	//		Find()
	//	assert.NoError(t, err)
	//	assert.True(t, len(users) > 0)
	//})
	//
	//t.Run("Multiple Languages", func(t *testing.T) {
	//	// 测试多语言支持
	//	testUsers := []*model.User{
	//		{Name: "Wang", Bio: "资深网页开发工程师"},
	//		{Name: "Zhang", Bio: "全栈开发工程师"},
	//	}
	//	err := u.WithContext(ctx).Create(testUsers)
	//	assert.NoError(t, err)
	//
	//	users, err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{Language: "chinese"}).
	//		TSQuery("工程师", u.Bio).
	//		Find()
	//	assert.NoError(t, err)
	//	assert.Equal(t, 2, len(users))
	//})
	//
	//t.Run("Ranking with Multiple Fields", func(t *testing.T) {
	//	// 测试多字段权重排序
	//	users, err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{
	//			Language: "english",
	//			Weights:  []string{"A", "B", "C"},
	//		}).
	//		TSQuery("developer", u.Bio, u.Description, u.Name).
	//		TSRank("developer", u.Bio, u.Description, u.Name).
	//		Find()
	//	assert.NoError(t, err)
	//	assert.True(t, len(users) > 0)
	//})
	//
	//t.Run("Headline with Options", func(t *testing.T) {
	//	// 测试文本片段提取选项
	//	type Result struct {
	//		*model.User
	//		Headline string
	//	}
	//	var results []Result
	//	err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{Language: "english"}).
	//		TSQuery("developer", u.Bio).
	//		TSHeadline("developer", u.Bio).
	//		Find(&results)
	//	assert.NoError(t, err)
	//	assert.True(t, len(results) > 0)
	//	assert.Contains(t, results[0].Headline, "<b>developer</b>")
	//})
	//
	//t.Run("Invalid Language", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		WithTSConfig(gen.TSConfig{Language: "invalid_lang"}).
	//		TSQuery("developer", u.Bio).
	//		Find()
	//	assert.Error(t, err)
	//})
	//
	//t.Run("Invalid Query Syntax", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		TSQuery("web &", u.Bio).
	//		Find()
	//	assert.Error(t, err)
	//})
	//
	//t.Run("Empty Fields", func(t *testing.T) {
	//	users, err := u.WithContext(ctx).
	//		TSQuery("web", u.Bio).
	//		Find()
	//	assert.NoError(t, err)
	//	assert.Equal(t, 0, len(users))
	//})
}
