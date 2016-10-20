package roblox_challenge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSimpleDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roblox Challenge Suite")
}
