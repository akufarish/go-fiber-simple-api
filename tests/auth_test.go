package tests

import (
	"go-fiber/app/controllers"
	"testing"
)

func TestSayHello(t *testing.T) {
	t.Logf("Say hello to Farish %s", controllers.Login("Farish"))

	if controllers.Login("Farish") != "Hello Farish" {
		t.Errorf("Salah harusnya %s", "Hello Farish")
	}
}