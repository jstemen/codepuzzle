package roblox_challenge

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"math"
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

	Context("Multiply", func() {

		It("Should return 0 when the slice is empty", func() {
			s1 := make([]int32, 0)
			product, err := FindLargestProduct(s1)
			Expect(err).To(BeNil())
			Expect(product).To(Equal(int32(0)))
		})

		It("Should return the number when this is only one", func() {
			s1 := []int32{3}
			product, err := FindLargestProduct(s1)
			Expect(err).To(BeNil())
			Expect(product).To(Equal(int32(3)))
		})

		It("Should return the product when only 2 numbers are given", func() {
			s1 := []int32{3, 2}
			product, err := FindLargestProduct(s1)
			Expect(err).To(BeNil())
			Expect(product).To(Equal(int32(6)))
		})

		It("Should work for all positives", func() {
			s1 := []int32{10, 3, 10, 2, 12, 5 }
			product, err := FindLargestProduct(s1)
			Expect(err).To(BeNil())
			Expect(product).To(Equal(int32(1200)))
		})

		It("Should work with the given example", func() {
			s1 := []int32{10, 3, 0, -2, 12, 5, 1, 1, 4}
			product, err := FindLargestProduct(s1)
			Expect(err).To(BeNil())
			Expect(product).To(Equal(int32(600)))
		})

		It("Should work for all negitives", func() {
			s1 := []int32{-3, -1, -3, -10}
			product, err := FindLargestProduct(s1)
			Expect(err).To(BeNil())
			Expect(product).To(Equal(int32(-9)))
		})

		It("Should return error when product greater than int32 max", func() {
			s1 := []int32{2147483646, 214748364, 21483647}
			_, err := FindLargestProduct(s1)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Overflow detected"))
		})
		It("Should return error when product less than int32 min", func() {
			s1 := []int32{1, -2147483648, 2147483647}
			_, err := FindLargestProduct(s1)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Underflow detected"))
		})
	})

	Context("multi", func() {

		//main logic test cases
		It("Should error on intMax * intMax", func() {
			_, err := overflowDetectingMulti(math.MaxInt32, math.MaxInt32)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Overflow detected"))
		})

		It("Should error on intMax * initMin", func() {
			_, err := overflowDetectingMulti(math.MinInt32, math.MaxInt32)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Underflow detected"))
		})

		It("Should error on initMin * intMax", func() {
			_, err := overflowDetectingMulti(math.MinInt32, math.MaxInt32)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Underflow detected"))
		})

		It("Should error on initMin * initMin", func() {
			_, err := overflowDetectingMulti(math.MinInt32, math.MinInt32)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Overflow detected"))
		})

		// 2's complement test cases

		It("Should error on initMin * -1", func() {
			_, err := overflowDetectingMulti(math.MinInt32, -1)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Overflow detected"))
		})

		It("Should error on -1 * initMin", func() {
			_, err := overflowDetectingMulti(-1, math.MinInt32)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(MatchRegexp("Overflow detected"))
		})

		//-1
		It("Should not error on -1 * -1", func() {
			_, err := overflowDetectingMulti(-1, -1)
			Expect(err).To(BeNil())
		})

		It("Should not error on -1 * 0", func() {
			_, err := overflowDetectingMulti(-1, 0)
			Expect(err).To(BeNil())
		})

		It("Should not error on -1 * -1", func() {
			_, err := overflowDetectingMulti(-1, 0)
			Expect(err).To(BeNil())
		})
		//0
		It("Should not error on 0 * -1", func() {
			_, err := overflowDetectingMulti(0, -1)
			Expect(err).To(BeNil())
		})

		It("Should not error on 0 * 0", func() {
			_, err := overflowDetectingMulti(0, 0)
			Expect(err).To(BeNil())
		})

		It("Should not error on 0 * 1", func() {
			_, err := overflowDetectingMulti(0, 1)
			Expect(err).To(BeNil())
		})

		//1

		It("Should not error on 1 * -1", func() {
			_, err := overflowDetectingMulti(1, -1)
			Expect(err).To(BeNil())
		})

		It("Should not error on 1 * 0", func() {
			_, err := overflowDetectingMulti(1, 0)
			Expect(err).To(BeNil())
		})

		It("Should not error on 1 * 1", func() {
			_, err := overflowDetectingMulti(1, 1)
			Expect(err).To(BeNil())
		})

	})
})
