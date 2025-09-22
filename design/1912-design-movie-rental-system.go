package design

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/design-movie-rental-system/

type film struct {
	price int
	shop  int
	movie int
}

type MovieRentingSystem struct {
	shops     map[int][]*film
	rented    *pkg.Heap[*film]
	prices    map[int64]int
	available map[int64]bool
}

func NewMovieRentingSystem(n int, entries [][]int) MovieRentingSystem {
	_ = n
	m := MovieRentingSystem{
		shops: make(map[int][]*film),
		rented: pkg.NewHeap(
			make([]*film, 0),
			func(x, y *film) bool {
				return x.price < y.price ||
					x.price == y.price && x.shop < y.shop ||
					x.price == y.price && x.shop == y.shop && x.movie < y.movie
			},
		),
		prices:    make(map[int64]int),
		available: make(map[int64]bool),
	}

	var shop, movie, price int
	for _, entry := range entries {
		shop, movie, price = entry[0], entry[1], entry[2]
		m.shops[movie] = append(m.shops[movie], &film{price: price, shop: shop, movie: movie})
		m.prices[m.pack(shop, movie)] = price
		m.available[m.pack(shop, movie)] = true
	}

	for movie := range m.shops {
		sort.Slice(m.shops[movie], func(i int, j int) bool {
			return m.shops[movie][i].price < m.shops[movie][j].price ||
				m.shops[movie][i].price == m.shops[movie][j].price &&
					m.shops[movie][i].shop < m.shops[movie][j].shop
		})
	}

	return m
}

func (m *MovieRentingSystem) Search(movie int) []int {
	shops := make([]int, 0)

	var films []*film

	for idx, n := 0, len(m.shops[movie]); idx < n && len(shops) < 5; idx++ {
		film := m.shops[movie][idx]
		if m.available[m.pack(film.shop, movie)] {
			shops = append(shops, film.shop)
			films = append(films, film)
			m.available[m.pack(film.shop, film.movie)] = false
		}
	}

	for _, film := range films {
		m.available[m.pack(film.shop, film.movie)] = true
	}

	return shops
}

func (m *MovieRentingSystem) Rent(shop int, movie int) {
	key := m.pack(shop, movie)
	price := m.prices[key]
	heap.Push(m.rented, &film{price: price, shop: shop, movie: movie})
	m.available[key] = false
}

func (m *MovieRentingSystem) Drop(shop int, movie int) {
	m.available[m.pack(shop, movie)] = true
}

func (m *MovieRentingSystem) Report() [][]int {
	var (
		res   = make([][]int, 0)
		films = make([]*film, 0)
		key   int64
	)

	for m.rented.Len() > 0 && len(res) < 5 {
		film, ok := heap.Pop(m.rented).(*film)
		_ = ok

		key = m.pack(film.shop, film.movie)
		if !m.available[key] {
			films = append(films, film)
			res = append(res, []int{film.shop, film.movie})
		}

		m.available[key] = true
	}

	for _, film := range films {
		heap.Push(m.rented, film)
		m.available[m.pack(film.shop, film.movie)] = false
	}

	return res
}

func (m *MovieRentingSystem) pack(shop int, movie int) int64 {
	return int64(shop<<32) + int64(movie)
}

// /**
//  * Your MovieRentingSystem object will be instantiated and called as such:
//  * obj := Constructor(n, entries);
//  * param_1 := obj.Search(movie);
//  * obj.Rent(shop,movie);
//  * obj.Drop(shop,movie);
//  * param_4 := obj.Report();
//  */
