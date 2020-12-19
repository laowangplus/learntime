package main

type user struct {
	name  string
	email string
}

func main() {
	_ = createUserV1()
	_ = createUserV2()

	//println("u1", &u1, "u2", &u2)
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//println("V2", u)
	return &u
}
