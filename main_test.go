package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The main file", func() {
	It("should read a csv file", func() {
		Expect(readCSV("test.csv")).ToNot(BeNil())
	})
})
