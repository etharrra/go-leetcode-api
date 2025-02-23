package handler

import (
	"encoding/json"
	"fmt"

	logger "github.com/etharrra/go-leetcode-api/loggger"
	"github.com/gofiber/fiber/v2"
)

// GraphQL request structure
type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

func FetchUserProfile(c *fiber.Ctx) error {
	username := c.AllParams()["username"]
	url := "https://leetcode.com/graphql"

	payload := GraphQLRequest{
		Query: Query,
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

	if err := logger.WriteLog(string(body)); err != nil {
		fmt.Println("Error writing log:", err)
	}

	response := GraphQLResponse{}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": fmt.Sprintf("Request errors: %v", err), "data": nil})
	}

	return c.JSON(fiber.Map{"status": code, "data": response.Data})
}
