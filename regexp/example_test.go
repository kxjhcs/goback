package regexp_test

import (
	"fmt"

	"github.com/h2so5/goback/regexp"
)

func Example() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^(\w)\w+\k{1}\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
	// Output:
	// false
	// true
	// false
	// false
}

func Example_BackReference() {
	re := regexp.MustCompile(`^(\w)\w+\k{1}$`)
	fmt.Println(re.MatchString("acca"))
	fmt.Println(re.MatchString("accccab"))
	fmt.Println(re.MatchString("AA"))
	// Output:
	// true
	// false
	// false
}

func Example_PossessiveQualifiers() {
	re := regexp.MustCompile(`^[0-9]++[0-9a]`)
	fmt.Println(re.MatchString("1234a"))
	fmt.Println(re.MatchString("1234"))
	// Output:
	// true
	// false
}

func Example_AtomicGroup() {
	re := regexp.MustCompile(`^(?>[0-9]+)[0-9a]`)
	fmt.Println(re.MatchString("1234a"))
	fmt.Println(re.MatchString("1234"))
	// Output:
	// true
	// false
}

func Example_Lookahead() {
	re := regexp.MustCompile(`a(?=[0-9]{3})1`)
	fmt.Println(re.MatchString("a123"))
	fmt.Println(re.MatchString("a12a"))
	// Output:
	// true
	// false
}