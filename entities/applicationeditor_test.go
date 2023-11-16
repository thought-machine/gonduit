package entities_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/samwestmoreland/gonduit/entities"
)

var _ = Describe("ObjectIdentifier", func() {
	Describe("serialized to JSON", func() {
		It("should correctly serialize the ID", func() {
			objId := ObjectIdentifier{ID: 1337}

			bytes, err := json.Marshal(objId)
			Ω(err).Should(BeNil())

			actual := string(bytes)
			expected := `1337`
			Ω(actual).Should(Equal(expected))
		})
		It("should correctly serialize the PHID", func() {
			objId := ObjectIdentifier{PHID: "PHID-TASK-muqrhdnsnovqs5c6qiqc"}

			bytes, err := json.Marshal(objId)
			Ω(err).Should(BeNil())

			actual := string(bytes)
			expected := `"PHID-TASK-muqrhdnsnovqs5c6qiqc"`
			Ω(actual).Should(Equal(expected))
		})
		It("should correctly serialize the Monogram", func() {
			objId := ObjectIdentifier{Monogram: "T1337"}

			bytes, err := json.Marshal(objId)
			Ω(err).Should(BeNil())

			actual := string(bytes)
			expected := `"T1337"`
			Ω(actual).Should(Equal(expected))
		})
		It("should prefer the PHID", func() {
			objId := ObjectIdentifier{ID: 1337, PHID: "PHID-TASK-muqrhdnsnovqs5c6qiqc", Monogram: "T1337"}

			bytes, err := json.Marshal(objId)
			Ω(err).Should(BeNil())

			actual := string(bytes)
			expected := `"PHID-TASK-muqrhdnsnovqs5c6qiqc"`
			Ω(actual).Should(Equal(expected))
		})
	})
})
