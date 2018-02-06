package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The csv reader functionality", func() {
	It("should read a csv file when given", func() {
		problems := readCSV("test/testProblems.csv")
		Expect(len(problems)).To(Equal(2))
		Expect(problems[0].Question).To(Equal("5+5"))
		Expect(problems[0].Answer).To(Equal("10"))
		Expect(problems[1].Question).To(Equal("7+3"))
		Expect(problems[1].Answer).To(Equal("10"))
	})

	It("should read data/problems.csv when given an empty string", func() {
		problems := readCSV("")
		Expect(len(problems)).To(Equal(13))
	})

	It("should read data/problems.csv when given a bad csv file", func() {
		problems := readCSV("badFile.csv")
		Expect(len(problems)).To(Equal(13))
	})
})
