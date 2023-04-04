package main


import (
	"fmt"
	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	pref := cid.Prefix{
		Version: 1,
		Codec: uint64(mc.Raw),
		MhType: mh.SHA2_256,
		MhLength: -1, // default length
	}
	
	// And then feed it some data
	c, err := pref.Sum([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}

	// convert c to string
	fmt.Println("Created CID: ", c.String())
}