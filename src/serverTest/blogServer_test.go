package serverTest

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestServer(t *testing.T)  {
	store := StubPlayerStore{
        map[string]int{
            "Pepper": 20,
            "Floyd":  10,
        },
    }
	server := &PlayerServer{&store}
	fmt.Println(server)
	t.Run("returns Pepper's score", func (t *testing.T)  {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func assertResponseBody(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
    }
}

// // 争取不直接修改struct本身
// func GetPlayerScore(name string) string {
//     if name == "Pepper" {
//         return "20"
//     }

//     if name == "Floyd" {
//         return "10"
//     }

//     return ""
// }



type StubPlayerStore struct {
    scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}