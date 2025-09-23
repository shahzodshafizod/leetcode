package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestTwitter$
func TestTwitter(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   [][]int
	}{
		{
			commands: []string{"Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"},
			values:   [][]int{{}, {1, 5}, {1}, {1, 2}, {2, 6}, {1}, {1, 2}, {1}},
			output:   [][]int{nil, nil, {5}, nil, nil, {6, 5}, nil, {5}},
		},
		{
			commands: []string{"Twitter", "postTweet", "postTweet", "getNewsFeed"},
			values:   [][]int{{}, {1, 5}, {1, 3}, {1}},
			output:   [][]int{nil, nil, nil, {3, 5}},
		},
		{
			commands: []string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed"},
			values:   [][]int{{}, {1, 5}, {1, 3}, {1, 101}, {1, 13}, {1, 10}, {1, 2}, {1, 94}, {1, 505}, {1, 333}, {1}},
			output:   [][]int{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, {333, 505, 94, 2, 10, 13, 101, 3, 5}},
		},
	} {
		var twitter Twitter

		for index, command := range tc.commands {
			var output []int

			switch command {
			case "Twitter":
				twitter = NewTwitter()
			case "postTweet":
				userID, tweetID := tc.values[index][0], tc.values[index][1]
				twitter.PostTweet(userID, tweetID)
			case "getNewsFeed":
				userID := tc.values[index][0]
				output = twitter.GetNewsFeed(userID)
			case "follow":
				followerID, followeeID := tc.values[index][0], tc.values[index][1]
				twitter.Follow(followerID, followeeID)
			case "unfollow":
				followerID, followeeID := tc.values[index][0], tc.values[index][1]
				twitter.Unfollow(followerID, followeeID)
			default:
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
