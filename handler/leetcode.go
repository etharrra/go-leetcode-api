package handler

import (
	"encoding/json"
	"fmt"

	logger "github.com/etharrra/go-leetcode-api/loggger"
	"github.com/gofiber/fiber/v2"
)

// GraphQL query
const query = `
	query getUserProfile($username: String!) {
		allQuestionsCount {
			difficulty
			count
		}
		matchedUser(username: $username) {
			contributions {
				points
			}
			profile {
				reputation
				ranking
			}
			submissionCalendar
			submitStats {
				acSubmissionNum {
					difficulty
					count
					submissions
				}
				totalSubmissionNum {
					difficulty
					count
					submissions
				}
			}
		}
		recentSubmissionList(username: $username) {
			title
			titleSlug
			timestamp
			statusDisplay
			lang
		}
		matchedUserStats: matchedUser(username: $username) {
			submitStats: submitStatsGlobal {
				acSubmissionNum {
					difficulty
					count
					submissions
				}
				totalSubmissionNum {
					difficulty
					count
					submissions
				}
			}
		}
	}
`

// GraphQL request structure
type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

// Define response structure
type Response struct {
	MatchedUser struct {
		Profile struct {
			Reputation int `json:"reputation"`
			Ranking    int `json:"ranking"`
		} `json:"profile"`
	} `json:"matchedUser"`

	RecentSubmissionList []struct {
		Title         string `json:"title"`
		TitleSlug     string `json:"titleSlug"`
		Timestamp     string `json:"timestamp"`
		StatusDisplay string `json:"statusDisplay"`
		Lang          string `json:"lang"`
	} `json:"recentSubmissionList"`
}

func FetchUserProfile(c *fiber.Ctx) error {
	username := c.AllParams()["username"]
	url := "https://leetcode.com/graphql"

	payload := GraphQLRequest{
		Query: query,
		Variables: map[string]any{
			"username": username,
		},
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": fmt.Sprintf("Failed to Marshal request %v", err), "data": nil})
	}

	a := fiber.AcquireAgent()
	a.Referer("https://leetcode.com")
	a.ContentType("application/json")
	a.Body(jsonData)
	req := a.Request()
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)

	if err := a.Parse(); err != nil {
		panic(err)
	}

	code, body, errs := a.Bytes()
	if len(errs) > 0 {
		return c.JSON(fiber.Map{"status": 500, "message": fmt.Sprintf("Request errors: %v", errs), "data": nil})
	}

	logger.WriteLog(string(body))
	return c.JSON(fiber.Map{"status": code, "message": fmt.Sprintf("Hello, I'm %s", username), "data": nil})
}
