package main_test

import (
	"testing"

	. "github.com/dfreilich/leetcode-go-solutions/1_two_sum"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestTwoSum(t *testing.T) {
	spec.Run(t, "(1) Two Sum: https://leetcode.com/problems/two-sum/", testTwoSum, spec.Report(report.Terminal{}))
}

func testTwoSum(t *testing.T, when spec.G, it spec.S) {
	var Expect func(interface{}, ...interface{}) Assertion
	it.Before(func() {
		Expect = NewWithT(t).Expect
	})

	when("runs", func() {
		it("succesfully passes", func() {
			Expect(Run()).To(Succeed())
		})
	})

}
