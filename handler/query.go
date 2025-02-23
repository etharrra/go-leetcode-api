package handler

// GraphQL query
const Query = `
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
		matchedUserTagProblemCounts: matchedUser(username: $username) {
			tagProblemCounts {
				advanced {
					tagName
					tagSlug
					problemsSolved
				}
				intermediate {
					tagName
					tagSlug
					problemsSolved
				}
				fundamental {
					tagName
					tagSlug
					problemsSolved
				}
			}
		}
	}
`
