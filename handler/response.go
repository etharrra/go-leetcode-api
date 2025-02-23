package handler

// Define response structure
type GraphQLResponse struct {
	Data Data `json:"data"`
}

// Data represents the main data structure
type Data struct {
	AllQuestionsCount           []QuestionCount             `json:"allQuestionsCount"`
	MatchedUser                 MatchedUser                 `json:"matchedUser"`
	RecentSubmissions           []Submission                `json:"recentSubmissionList"`
	MatchedUserStats            UserStats                   `json:"matchedUserStats"`
	MatchedUserTagProblemCounts MatchedUserTagProblemCounts `json:"matchedUserTagProblemCounts"`
}

// QuestionCount represents the total count of questions per difficulty
type QuestionCount struct {
	Difficulty string `json:"difficulty"`
	Count      int    `json:"count"`
}

// MatchedUser represents the user's profile
type MatchedUser struct {
	Contributions      Contributions `json:"contributions"`
	Profile            Profile       `json:"profile"`
	SubmissionCalendar string        `json:"submissionCalendar"`
	SubmitStats        SubmitStats   `json:"submitStats"`
}

// Contributions represents user contribution points
type Contributions struct {
	Points int `json:"points"`
}

// Profile represents user ranking and reputation
type Profile struct {
	Reputation int `json:"reputation"`
	Ranking    int `json:"ranking"`
}

// SubmitStats represents submission statistics
type SubmitStats struct {
	AcSubmissionNum    []SubmissionStats `json:"acSubmissionNum"`
	TotalSubmissionNum []SubmissionStats `json:"totalSubmissionNum"`
}

// SubmissionStats represents statistics of accepted and total submissions
type SubmissionStats struct {
	Difficulty  string `json:"difficulty"`
	Count       int    `json:"count"`
	Submissions int    `json:"submissions"`
}

// Submission represents a recent submission by the user
type Submission struct {
	Title         string `json:"title"`
	TitleSlug     string `json:"titleSlug"`
	Timestamp     string `json:"timestamp"`
	StatusDisplay string `json:"statusDisplay"`
	Lang          string `json:"lang"`
}

// UserStats represents user submission statistics
type UserStats struct {
	SubmitStats SubmitStats `json:"submitStats"`
}

// MatchedUserTagProblemCounts represents user skills based on level
type MatchedUserTagProblemCounts struct {
	TagProblemCounts TagProblemCounts `json:"tagProblemCounts"`
}

// TagProblemCounts represents skill levels
type TagProblemCounts struct {
	Advanced     []LevelDetails `json:"advanced"`
	Intermediate []LevelDetails `json:"intermediate"`
	Fundamental  []LevelDetails `json:"fundamental"`
}

// LevelDetails represents detilas of each levels
type LevelDetails struct {
	TagName        string `json:"tagName"`
	TagSlug        string `json:"tagSlug"`
	ProblemsSolved int    `json:"problemsSolved"`
}
