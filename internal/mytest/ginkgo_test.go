package mytest_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	example "github.com/niqdev/gopher-labs/pkg"
)

var _ = Describe("example", func() {

	It("verify empty", func() {
		_, err := example.Hello("")

		Expect(err.Error()).Should(Equal("empty name"))
	})

	It("verify hello", func() {
		result, err := example.Hello("test")

		Expect(result).Should(Equal("hello test"))
		Expect(err).Should(BeNil())
	})
})
