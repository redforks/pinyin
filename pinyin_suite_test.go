package pinyin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPinyin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pinyin Suite")
}
