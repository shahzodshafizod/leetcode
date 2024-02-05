package design

import (
	"container/heap"
)

// https://leetcode.com/problems/design-twitter/

type tweet struct {
	id       int
	uniqueId int
}

type tweets []*tweet

var _ heap.Interface = &tweets{}

func (t tweets) Len() int           { return len(t) }
func (t tweets) Less(i, j int) bool { return t[i].uniqueId > t[j].uniqueId }
func (t *tweets) Swap(i, j int)     { (*t)[i], (*t)[j] = (*t)[j], (*t)[i] }
func (t *tweets) Push(x any)        { *t = append(*t, x.(*tweet)) }
func (t *tweets) Pop() any {
	if t.Len() == 0 {
		return nil
	}
	var value = (*t)[t.Len()-1]
	*t = (*t)[:t.Len()-1]
	return value
}

type Twitter struct {
	tweets    map[int][]*tweet
	followees map[int]map[int]bool
	nextId    int
}

func NewTwitter() Twitter {
	return Twitter{
		tweets:    make(map[int][]*tweet),
		followees: make(map[int]map[int]bool),
		nextId:    1,
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) { // O(1)
	var newTweet = &tweet{id: tweetId, uniqueId: t.nextId}
	t.nextId++
	if t.tweets[userId] == nil {
		t.tweets[userId] = []*tweet{newTweet}
	} else {
		t.tweets[userId] = append(t.tweets[userId], newTweet)
	}
}

func (t *Twitter) GetNewsFeed(userId int) []int { // O(N)
	var tweets = tweets(t.tweets[userId])
	for followeeId := range t.followees[userId] {
		tweets = append(tweets, t.tweets[followeeId]...)
	}

	heap.Init(&tweets)
	var result = make([]int, 0)
	var count = 10
	for tweets.Len() > 0 && count > 0 {
		count--
		tweet := heap.Pop(&tweets).(*tweet)
		result = append(result, tweet.id)
	}
	return result
}

func (t *Twitter) Follow(followerId int, followeeId int) { // O(1)
	if t.followees[followerId] == nil {
		t.followees[followerId] = make(map[int]bool)
	}
	t.followees[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) { // O(1)
	if t.followees[followerId] != nil && t.followees[followerId][followeeId] {
		delete(t.followees[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
