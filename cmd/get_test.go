package cmd

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("go-cli-app get commad execution", func() {

	Describe("Positive Test", func() {
		Context("When file available for download", func() {
			It("Should get downloaded", func() {
				// The Ginkgo test runner takes over os.Args and fills it with
				// its own flags.  This makes the cobra command arg parsing
				// fail because of unexpected options.  Work around this.

				// Save a copy of os.Args
				origArgs := os.Args[:]
				fmt.Println(origArgs)

				// Trim os.Args to only the first arg + get
				os.Args[1] = "get" // trim to only the first arg + get, which is the command itself
				os.Args = os.Args[:2]
				err := Execute()
				// Restore os.Args
				os.Args = origArgs[:]
				Expect(err).NotTo(HaveOccurred())
				//fmt.Println("Checking")

			})
		})

	})

})

func TestGinkoSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "go-cli-app Ginkgo Suites")

}
