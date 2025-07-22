package dto

const (
	MESSAGE_SUCCESS_GET_DASHBOARD = "succesfully get dashboard statistics"
	MESSAGE_FAILED_GET_DASHBOARD  = "Failed to get dashboard"
)

type (
	GetDashboardResponse struct {
		CountUsers    int            `json:"count_users"`
		CountReview   int            `json:"count_review"`
		TopFilms      []TopFilm      `json:"top_films"`
		TrendingFilms []TrendingFilm `json:"trending_films"`
		WeeklyUsers   []WeeklyUser   `json:"weekly_users"`
		WeeklyReviews []WeeklyReview `json:"weekly_reviews"`
	}
)
