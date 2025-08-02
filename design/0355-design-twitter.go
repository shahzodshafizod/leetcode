package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/design-twitter/

type tweet struct {
	id       int
	uniqueID int
}

type Twitter struct {
	tweets    map[int][]*tweet
	followees map[int]map[int]bool
	nextID    int
}

func NewTwitter() Twitter {
	return Twitter{
		tweets:    make(map[int][]*tweet),
		followees: make(map[int]map[int]bool),
		nextID:    1,
	}
}

func (t *Twitter) PostTweet(userID int, tweetID int) { // O(1)
	newTweet := &tweet{id: tweetID, uniqueID: t.nextID}
	t.nextID++

	if t.tweets[userID] == nil {
		t.tweets[userID] = []*tweet{newTweet}
	} else {
		t.tweets[userID] = append(t.tweets[userID], newTweet)
	}
}

func (t *Twitter) GetNewsFeed(userID int) []int { // O(N Log N)
	posts := make([]*tweet, len(t.tweets[userID]))
	for followeeID := range t.followees[userID] {
		posts = append(posts, t.tweets[followeeID]...)
	}

	copy(posts, t.tweets[userID])
	tweets := pkg.NewHeap(posts, func(x, y *tweet) bool { return x.uniqueID > y.uniqueID })
	heap.Init(tweets)

	result := make([]int, 0)

	count := 10
	for tweets.Len() > 0 && count > 0 {
		count--
		tweet, ok := heap.Pop(tweets).(*tweet)
		_ = ok

		result = append(result, tweet.id)
	}

	return result
}

func (t *Twitter) Follow(followerID int, followeeID int) { // O(1)
	if t.followees[followerID] == nil {
		t.followees[followerID] = make(map[int]bool)
	}

	t.followees[followerID][followeeID] = true
}

func (t *Twitter) Unfollow(followerID int, followeeID int) { // O(1)
	if t.followees[followerID] != nil && t.followees[followerID][followeeID] {
		delete(t.followees[followerID], followeeID)
	}
}

/*
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
