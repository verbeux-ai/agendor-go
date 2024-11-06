package agendor_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/verbeux-ai/agendor-go"
)

var client *agendor.Client

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")

	client = agendor.NewClient(
		agendor.WithToken(os.Getenv("TOKEN")),
	)

	os.Exit(m.Run())
}
