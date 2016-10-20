package roblox_challenge_test

import (
	. "github.com/jstemen/roblox_challenge"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("solution", func() {
	Context("Merger", func() {
		It("Should be able to merge when the first slice is empty", func() {
			s1 := make([]int32, 0)
			s2 := []int32{1, 2, 3}
			merged := Merger(s1, s2)
			Expect(merged).To(Equal(s2))
		})

		It("Should be able to merge when the second slice is empty", func() {
			s1 := make([]int32, 0)
			s2 := []int32{1, 2, 3}
			merged := Merger(s2, s1)
			Expect(merged).To(Equal(s2))
		})

		It("Should be able to merge 2 empty slices", func() {
			s1 := make([]int32, 0)
			s2 := make([]int32, 0)
			merged := Merger(s2, s1)
			Expect(merged).To(Equal(s2))
		})

		It("Should be able to merge the given example", func() {
			s1 := []int32{0, 2, 3}
			s2 := []int32{1, 3, 4, 5}
			merged := Merger(s1, s2)
			Expect(merged).To(Equal([]int32{0, 1, 2, 3, 3, 4, 5}))
		})

	})
})
