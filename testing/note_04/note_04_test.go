// source: https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


// To bring into existence a web server I need ListenAndServe function:
// 		func ListenAndServe(addr string, handler Handler) error 
// This amazing little fella starts a web server on a port (not a place near the sea) 
// It creates a goroutine for every request it receives and runs that 
// goroutine against Handler	


// Here Handler is an interface:
// type Handler interface {
//		  ServeHTTP(ResponseWriter, *Request)
// }
// Since interface handles what happens to incoming request I can say that 
// any type that implements the Handler interface will now take the burden of
// handling the request depending on how it implements ServerHTTP method.
// Because, if I look at ServeHTTP method, it takes two arguments:
// 1) ResponseWriter which is in charge of crafting a response to be send off
// 2) *Request represents a client request the server receives, pointer to Request implies
// that it is instantiated well before it passed down into our ServeHTTP method


// TestGETPlayers supposedly tests how well my PlayerServer function is gonna do
// I pass *testing.T to my test function, because it comes with all good stuff in it
//, which brings benefit when testing things
// So what I am going to test here anyway?
// Let me imagine that I am a client called Pepper and I want to know my score   
// Secretly, I know my score, but I am suspicious if PlayerServer tells me the 
// truth or not 
// Let me get inside now 
func TestGETPlayers(t *testing.T) {

	// according to Wikipedia, this is called sub-test
	// joking not Wikipedia but official testing docs
	t.Run("returns Pepper's score", func(t *testing.T) {
		// Here, I am, I am creating a necessary *http.Request 
		// (I'm pretty sure no error pops up)
		request, _  := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		
		// Here I am fabricating a *httptest.ResponseRecorder 
		// this thing is a fake alternative to http.ResponseWriter   
		response := httptest.NewRecorder() 

		// Here I passing response and request down into the function 
		// which is under a serious investigation
		// wait ...
		// I am passing request which PlayerServer promised to accept.
		// but response has a type *httptest.ResponseRecorder which is 
		// different from what PlayerServer promised to accept. 
		// The culmination of my self talk is here. Here I am making fun of 
		// PlayerServer thanks to the power of interfaces in golang.
		// I gonna scroll down to bottom to see explanation of this 
		// mocking mechanism
		PlayerServer(response, request)
		// if I look carefully I passed response which was a pointer type
		// the reason is that I am gonna see what happened to him 
		// while he was inside PlayerServer

		// Now I am seeing that response got a body
		// but i don't know the value in it
		got := response.Body.String()
		want := "20"

		// here I will compare got with want (the real score I know)
		// if they are not the same, then I can blame PlayerServer 
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	
}



// Mocking mechanism explanation:
// Let me say, Initially there is a type called BrandWatch:

// type BrandWatch struct {
//		Name 		  string
//      SellingPrice  float64
//      costPrice 	  float32          // do not export cost price
// }

// Now let me say again, after a few months, a new type 
// is brought into existence and its name is MockBrandWatch: 
// (I could call it FakeBrandWatch, but MockBrandWatch justifies it name better
// since it is not only fake but also makes fun of real BrandWatch)

// type MockBrandWatch struct {
//		Name			string	           
//		SellingPrice    float64
//      constPrice      float32         // do not export cost price
// }

// From now on I have to have a little chat with a gopher version of me
// Dialog between me and †gopher   
// †gopher = gopher version of me   
// Disclaimer! gopher version of me is not a expert at being a gopher yet.

// me:     What gives MockBrandWatch a power to mock BrandWatch?
// gopher: Hmm let me think...     
// me:     is that they have exactly the same fields?
// gopher: Look! there must be a rule to accept MockBrandWatch as BrandWatch, right?
// 		   so, people who created golang made a rule for that occasions
//         The rule says there should be another third party thing beetween those
//         types: BrandWatch and MockBrandWatch. 
//         That thing is called interface, Interface is also a type.
//         It declares a list of zero or more method signatures and says
//         any type that implements these methods can inherit my type. 
// me:     got it! Let me tell what I got from your answer. So interface type says
//         that whoever wants to refer his name, must prove that they can do the actions
//         the interface declares.
// gopher: correct!

// Now I am done with talking to gopher
// Let me go back to my explanation. So I am creating an interface time now:

// type Watch interface {
//      HasUniqueShape() (Shape, bool)                      	
// } 
// Ahhh I am bored writing...
// Lemme say BrandWatch has a unique shape:
// func (bw BrandWatch) HasUniqueShape() (Shape, bool) { 
//      // hidden implementation 	
// }
// Now I am saying MockBrandWatch implements that method too
// func (bw MockBrandWatch) HasUniqueShape() (Shape, bool) { 
//      // hidden implementation 	
// }
// they both implemented methods in Watch interface, now they can mock each other
// I am done, let me scroll up inside test function