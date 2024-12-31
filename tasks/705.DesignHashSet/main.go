package main

//https://leetcode.com/problems/design-hashset/description/

type MyHashSet struct {
	data map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		data: make(map[int]struct{}),
	}
}

func (s *MyHashSet) Add(key int) {
	_, exist := s.data[key]
	if !exist {
		s.data[key] = struct{}{}
	}
}

func (s *MyHashSet) Remove(key int) {
	delete(s.data, key)
}

func (s *MyHashSet) Contains(key int) bool {
	_, exist := s.data[key]
	return exist
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
