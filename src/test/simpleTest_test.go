package test

import "testing"


func TestHello(t *testing.T) {
    // simple test just with one function
    // got := Hello()
    // want := "Hello, world"
    // if got != want {
    //     t.Errorf("got '%s', want '%s'", got, want)
    // }
    
    assertCorrectMessage := func(t *testing.T, got, want string) {
        // 使错误追踪不会在子程序中
        // simpleTest_test.go:17: got 'Hello, 123' want 'Hello, world'
        // simpleTest_test.go:31: got 'Hello, 123' want 'Hello, world'
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }
    

    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello world when an empty string is supplied", func(t *testing.T){
        got := Hello("")
        want := "Hello, world"
        assertCorrectMessage(t, got, want)
    })
}