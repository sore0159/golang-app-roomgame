package island

import (
	"fmt"
	//"log"
	"log"
	"os"
	"testing"
)

func TestIsle(t *testing.T) {
	err := os.Chdir("server")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TESTING")
	g := New("test1")
	fmt.Printf("%+v\n", g)

}
