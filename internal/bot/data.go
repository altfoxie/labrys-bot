package bot

import (
	"github.com/altfoxie/labrys-bot/internal/i18n"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
)

// DataRepository represents a repository, which provides
// searching and getting data with scored matches.
type DataRepository[T any, M Match[T]] interface {
	Init() error
	Search(query string) ([]M, error)
	GetN(n int) ([]T, error)
}

// Match represents a match.
type Match[T any] interface {
	Base() T
}

// InlineItem represents an item, which can be converted to inline result.
type InlineItem interface {
	InlineResult() telego.InlineQueryResult
}

func onDataQuery[T InlineItem, M Match[T]](
	bot *telego.Bot,
	locale *i18n.Locale,
	repo DataRepository[T, M],
	query *telego.InlineQuery,
	searchQuery string,
) error {
	// Shitty code because golang.
	var matches []M
	var items []T
	var err error
	if searchQuery == "" {
		items, err = repo.GetN(50)
	} else {
		matches, err = repo.Search(searchQuery)
		items = lo.Map(matches, func(match M, _ int) T {
			return match.Base()
		})
	}
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return bot.AnswerInlineQuery(
			tu.InlineQuery(query.ID).
				WithSwitchPmText(locale.NothingFound).
				WithSwitchPmParameter("ignore").
				WithCacheTime(-1).
				WithIsPersonal(),
		)
	}

	return bot.AnswerInlineQuery(
		tu.InlineQuery(
			query.ID,
			lo.Map(items, func(item T, _ int) telego.InlineQueryResult {
				return item.InlineResult()
			})...,
		).WithCacheTime(-1).WithIsPersonal(),
	)
}
