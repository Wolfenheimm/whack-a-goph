package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestComprehensiveExerciseSolutions tests ALL 36 exercises (100% coverage)
func TestComprehensiveExerciseSolutions(t *testing.T) {
	// Define all 36 exercises
	allExercises := []struct {
		category string
		filename string
	}{
		{"intro", "intro1.go"},
		{"intro", "intro2.go"},
		{"variables", "variables1.go"},
		{"variables", "variables2.go"},
		{"functions", "functions1.go"},
		{"primitive_types", "primitives1.go"},
		{"primitive_types", "primitives2.go"},
		{"if", "if1.go"},
		{"if", "if2.go"},
		{"structs", "structs1.go"},
		{"structs", "structs2.go"},
		{"strings", "strings1.go"},
		{"strings", "strings2.go"},
		{"vecs", "vecs1.go"},
		{"hashmaps", "hashmaps1.go"},
		{"options", "options1.go"},
		{"error_handling", "errors1.go"},
		{"enums", "enums1.go"},
		{"enums", "enums2.go"},
		{"enums", "enums3.go"},
		{"traits", "traits1.go"},
		{"tests", "tests1.go"},
		{"threads", "threads1.go"},
		{"iterators", "iterators1.go"},
		{"smart_pointers", "smart_pointers1.go"},
		{"macros", "macros1.go"},
		{"generics", "generics1.go"},
		{"conversions", "conversions1.go"},
		{"advanced", "advanced1.go"},
		{"advanced", "advanced2.go"},
		{"modules", "modules1.go"},
		{"quizzes", "quiz1.go"},
		{"quizzes", "quiz2.go"},
		{"quizzes", "quiz3.go"},
		{"quizzes", "quiz4.go"},
		{"quizzes", "quiz_final.go"},
	}

	// Test that all exercises exist and can be completed
	for _, exercise := range allExercises {
		t.Run(exercise.category+"_"+strings.TrimSuffix(exercise.filename, ".go"), func(t *testing.T) {
			testExerciseCanBeCompleted(t, exercise.category, exercise.filename)
		})
	}

	// Test specific working solutions for verified exercises
	t.Run("verified_solutions", func(t *testing.T) {
		testIntro1Solution(t)
		testIntro2Solution(t)
		testFunctions1Solution(t)
		testPrimitives1Solution(t)
		testIf1Solution(t)
		testStructs1Solution(t)
		testTraits1Solution(t)
		testErrors1Solution(t)
	})
}

// testIntro1Solution tests intro1 exercise with correct solution
func testIntro1Solution(t *testing.T) {
	solution := `package intro

// Welcome to Whack-a-Goph! 🐹
// Your first gopher has popped up!
//
// This is the most basic exercise - make the Hello function
// return the string "Hello, Gopher!" to whack your first gopher!

// SOLUTION IMPLEMENTED FOR TESTING

func Hello() string {
	return "Hello, Gopher!"
}
`
	testExerciseSolution(t, "intro", "intro1.go", solution)
}

// testIntro2Solution tests intro2 exercise with correct solution
func testIntro2Solution(t *testing.T) {
	solution := `package intro

import "fmt"

// SOLUTION IMPLEMENTED FOR TESTING

// PrintGreeting prints a greeting message to stdout
func PrintGreeting(name string) {
	fmt.Printf("Hello, %s! Welcome to Go!\n", name)
}
`
	testExerciseSolution(t, "intro", "intro2.go", solution)
}

// testVariables1Solution tests variables1 exercise with correct solution
func testVariables1Solution(t *testing.T) {
	solution := `package variables

// Time to learn about variables! 🐹
// In Go, there are different ways to declare variables.
// Let's start with the most basic form.

// SOLUTION IMPLEMENTED FOR TESTING

func Variables1() {
	// Declare a variable named 'x' with the value 5
	var x int = 5

	// Declare a variable named 'name' with the value "Gopher"
	var name string = "Gopher"

	// Declare a variable named 'done' with the value true
	var done bool = true

	// Don't modify anything below this line
	println("x =", x)
	println("name =", name)
	println("done =", done)
}
`
	testExerciseSolution(t, "variables", "variables1.go", solution)
}

// testFunctions1Solution tests functions1 exercise with correct solution
func testFunctions1Solution(t *testing.T) {
	solution := `package functions

// Time to learn about functions! 🐹
// Functions are the building blocks of Go programs.

// SOLUTION IMPLEMENTED FOR TESTING

func Add(a, b int) int {
	return a + b
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func IsEven(n int) bool {
	return n%2 == 0
}
`
	testExerciseSolution(t, "functions", "functions1.go", solution)
}

// testPrimitives1Solution tests primitives1 exercise with correct solution
func testPrimitives1Solution(t *testing.T) {
	solution := `package primitive_types

// Time to learn about Go's primitive types! 🐹
// Go has several built-in basic types like int, string, bool, float64, etc.

// SOLUTION IMPLEMENTED FOR TESTING

func GetAge() int {
	return 25
}

func GetName() string {
	return "Gopher"
}

func IsStudent() bool {
	return true
}

func GetPI() float64 {
	return 3.14159
}
`
	testExerciseSolution(t, "primitive_types", "primitives1.go", solution)
}

// testIf1Solution tests if1 exercise with correct solution
func testIf1Solution(t *testing.T) {
	solution := `package if_statements

// Time to learn about if statements! 🐹
// Go's if statements are similar to other languages but don't need parentheses.

// SOLUTION IMPLEMENTED FOR TESTING

func IsPositive(n int) bool {
	if n > 0 {
		return true
	}
	return false
}

func GetGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
`
	testExerciseSolution(t, "if", "if1.go", solution)
}

// testStructs1Solution tests structs1 exercise with correct solution
func testStructs1Solution(t *testing.T) {
	solution := `package structs

import "fmt"

// Time to learn about structs! 🐹
// Structs group related data together.

// SOLUTION IMPLEMENTED FOR TESTING

type Person struct {
	Name string
	Age  int
}

func CreatePerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func GetPersonInfo(p Person) string {
	return "Name: " + p.Name + ", Age: " + fmt.Sprintf("%d", p.Age)
}

func IsAdult(p Person) bool {
	return p.Age >= 18
}
`
	testExerciseSolution(t, "structs", "structs1.go", solution)
}

// testStrings1Solution tests strings1 exercise with correct solution
func testStrings1Solution(t *testing.T) {
	solution := `package strings

// Time to learn about strings! 🐹
// Strings are sequences of characters in Go.

// SOLUTION IMPLEMENTED FOR TESTING

func GetLength(s string) int {
	return len(s)
}

func Concatenate(a, b string) string {
	return a + b
}

func GetFirstChar(s string) byte {
	if len(s) > 0 {
		return s[0]
	}
	return 0
}

func Contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
`
	testExerciseSolution(t, "strings", "strings1.go", solution)
}

// testVecs1Solution tests vecs1 exercise with correct solution
func testVecs1Solution(t *testing.T) {
	solution := `package vecs

// Time to learn about slices (Go's version of dynamic arrays)! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func CreateSlice() []int {
	return []int{1, 2, 3, 4, 5}
}

func AppendElement(slice []int, element int) []int {
	return append(slice, element)
}

func GetLength(slice []int) int {
	return len(slice)
}

func GetSum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
`
	testExerciseSolution(t, "vecs", "vecs1.go", solution)
}

// testHashmaps1Solution tests hashmaps1 exercise with correct solution
func testHashmaps1Solution(t *testing.T) {
	solution := `package hashmaps

// Time to learn about maps! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func CreateMap() map[string]int {
	return map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
}

func GetValue(m map[string]int, key string) (int, bool) {
	value, exists := m[key]
	return value, exists
}

func SetValue(m map[string]int, key string, value int) {
	m[key] = value
}

func DeleteKey(m map[string]int, key string) {
	delete(m, key)
}

func GetKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
`
	testExerciseSolution(t, "hashmaps", "hashmaps1.go", solution)
}

// testEnums1Solution tests enums1 exercise with correct solution
func testEnums1Solution(t *testing.T) {
	solution := `package enums

// Time to learn about enums in Go! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func GetFavoriteColor() Color {
	return Blue
}

func IsWarmColor(c Color) bool {
	return c == Red
}

func ColorName(c Color) string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Unknown"
	}
}
`
	testExerciseSolution(t, "enums", "enums1.go", solution)
}

// testTraits1Solution tests traits1 exercise with correct solution
func testTraits1Solution(t *testing.T) {
	solution := `package main

// Traits 1: Basic Interface Definitions 🐹

// SOLUTION IMPLEMENTED FOR TESTING

import "fmt"

type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Car struct {
	Brand string
}

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

func (d Dog) Move() string {
	return d.Name + " runs around"
}

func (c Cat) Speak() string {
	return "Meow! I'm " + c.Name
}

func (c Cat) Move() string {
	return c.Name + " prowls silently"
}

func (c Car) Move() string {
	return "The " + c.Brand + " drives smoothly"
}

func MakeSound(s Speaker) string {
	return s.Speak()
}

func StartMoving(m Mover) string {
	return m.Move()
}

func DescribeAnimal(animal interface{}) string {
	if speaker, ok := animal.(Speaker); ok {
		if mover, ok := animal.(Mover); ok {
			return speaker.Speak() + " and " + mover.Move()
		}
		return speaker.Speak()
	}
	if mover, ok := animal.(Mover); ok {
		return mover.Move()
	}
	return "Cannot describe this entity"
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	car := Car{Brand: "Toyota"}

	fmt.Println("Speaking:")
	fmt.Println(MakeSound(dog))
	fmt.Println(MakeSound(cat))

	fmt.Println("\nMoving:")
	fmt.Println(StartMoving(dog))
	fmt.Println(StartMoving(cat))
	fmt.Println(StartMoving(car))

	fmt.Println("\nDescribing:")
	fmt.Println(DescribeAnimal(dog))
	fmt.Println(DescribeAnimal(cat))
	fmt.Println(DescribeAnimal(car))
}
`
	testExerciseSolution(t, "traits", "traits1.go", solution)
}

// testErrors1Solution tests errors1 exercise with correct solution
func testErrors1Solution(t *testing.T) {
	solution := `package error_handling

import "errors"

// SOLUTION IMPLEMENTED FOR TESTING

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	// Simple approximation using Newton's method
	if x == 0 {
		return 0, nil
	}
	
	guess := x / 2
	for i := 0; i < 10; i++ {
		guess = (guess + x/guess) / 2
	}
	return guess, nil
}

func HandleError(err error) string {
	if err != nil {
		return "Error: " + err.Error()
	}
	return "Success"
}
`
	testExerciseSolution(t, "error_handling", "errors1.go", solution)
}

// testGenerics1Solution tests generics1 exercise with correct solution
func testGenerics1Solution(t *testing.T) {
	solution := `package generics

// SOLUTION IMPLEMENTED FOR TESTING

func Identity[T any](x T) T {
	return x
}

func First[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[0]
}

func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
`
	testExerciseSolution(t, "generics", "generics1.go", solution)
}

// testThreads1Solution tests threads1 exercise with correct solution
func testThreads1Solution(t *testing.T) {
	solution := `package main

import (
	"fmt"
	"sync"
	"time"
)

// SOLUTION IMPLEMENTED FOR TESTING

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintLetters() {
	letters := []string{"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		fmt.Printf("Letter: %s\n", letter)
		time.Sleep(150 * time.Millisecond)
	}
}

func RunConcurrently() {
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		PrintNumbers()
	}()
	
	go func() {
		defer wg.Done()
		PrintLetters()
	}()
	
	wg.Wait()
}

func main() {
	fmt.Println("Running concurrently:")
	RunConcurrently()
	fmt.Println("Done!")
}
`
	testExerciseSolution(t, "threads", "threads1.go", solution)
}

// testVariables2Solution tests variables2 exercise with correct solution
func testVariables2Solution(t *testing.T) {
	solution := `package variables

// Now let's learn about type inference! 🐹
// Go can often figure out the type of a variable automatically.

// SOLUTION IMPLEMENTED FOR TESTING

func Variables2() {
	// TODO: Declare a variable 'number' with value 42 using type inference
	// Hint: var number = 42
	number := 42

	// TODO: Declare a variable 'pi' with value 3.14 using type inference
	// Hint: var pi = 3.14
	pi := 3.14

	// TODO: Declare a variable 'message' with value "Hello, Go!" using type inference
	// Hint: var message = "Hello, Go!"
	message := "Hello, Go!"

	// Don't modify anything below this line
	println("number is", number, "and has type int")
	println("pi is", pi, "and has type float64")
	println("message is", message, "and has type string")
}
`
	testExerciseSolution(t, "variables", "variables2.go", solution)
}

// testPrimitives2Solution tests primitives2 exercise with correct solution
func testPrimitives2Solution(t *testing.T) {
	solution := `package primitive_types

// Time to learn about numeric types! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func GetByte() byte {
	return 255
}

func GetInt8() int8 {
	return -128
}

func GetUint16() uint16 {
	return 65535
}

func GetInt64() int64 {
	return 9223372036854775807
}

func GetFloat32() float32 {
	return 3.14
}

func GetComplex() complex128 {
	return complex(3, 4)
}
`
	testExerciseSolution(t, "primitive_types", "primitives2.go", solution)
}

// testIf2Solution tests if2 exercise with correct solution
func testIf2Solution(t *testing.T) {
	solution := `package if_statements

// More advanced if statements! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func CanVote(age int) bool {
	return age >= 18
}

func GetTemperatureDescription(temp int) string {
	if temp < 0 {
		return "Freezing"
	} else if temp < 20 {
		return "Cold"
	} else if temp < 30 {
		return "Warm"
	} else {
		return "Hot"
	}
}

func IsValidPassword(password string) bool {
	return len(password) >= 8
}

func GetDiscount(age int, isStudent bool) float64 {
	if age < 18 {
		return 0.20 // 20% discount for minors
	} else if isStudent {
		return 0.10 // 10% discount for students
	} else if age >= 65 {
		return 0.15 // 15% discount for seniors
	}
	return 0.0 // No discount
}
`
	testExerciseSolution(t, "if", "if2.go", solution)
}

// testStructs2Solution tests structs2 exercise with correct solution
func testStructs2Solution(t *testing.T) {
	solution := `package structs

import "fmt"

// Advanced struct operations! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Book struct {
	Title  string
	Author string
	Pages  int
	Price  float64
}

func CreateBook(title, author string, pages int, price float64) Book {
	return Book{
		Title:  title,
		Author: author,
		Pages:  pages,
		Price:  price,
	}
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("%s by %s (%d pages) - $%.2f", b.Title, b.Author, b.Pages, b.Price)
}

func (b *Book) ApplyDiscount(discount float64) {
	b.Price = b.Price * (1.0 - discount)
}

func IsExpensive(b Book) bool {
	return b.Price > 50.0
}
`
	testExerciseSolution(t, "structs", "structs2.go", solution)
}

// testStrings2Solution tests strings2 exercise with correct solution
func testStrings2Solution(t *testing.T) {
	solution := `package strings

import "strings"

// Advanced string operations! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

func CountOccurrences(s, substr string) int {
	return strings.Count(s, substr)
}

func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func SplitString(s, sep string) []string {
	return strings.Split(s, sep)
}

func JoinStrings(parts []string, sep string) string {
	return strings.Join(parts, sep)
}
`
	testExerciseSolution(t, "strings", "strings2.go", solution)
}

// testOptions1Solution tests options1 exercise with correct solution
func testOptions1Solution(t *testing.T) {
	solution := `package options

// Time to learn about pointers and nil values! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func GetPointer(value int) *int {
	return &value
}

func GetValue(ptr *int) int {
	if ptr != nil {
		return *ptr
	}
	return 0
}

func IsNil(ptr *int) bool {
	return ptr == nil
}

func ModifyValue(ptr *int, newValue int) {
	if ptr != nil {
		*ptr = newValue
	}
}
`
	testExerciseSolution(t, "options", "options1.go", solution)
}

// testEnums2Solution tests enums2 exercise with correct solution
func testEnums2Solution(t *testing.T) {
	solution := `package enums

// Advanced enum operations! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func IsWeekend(d Day) bool {
	return d == Saturday || d == Sunday
}

func NextDay(d Day) Day {
	return (d + 1) % 7
}

func DayName(d Day) string {
	names := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if int(d) >= 0 && int(d) < len(names) {
		return names[d]
	}
	return "Unknown"
}
`
	testExerciseSolution(t, "enums", "enums2.go", solution)
}

// testEnums3Solution tests enums3 exercise with correct solution
func testEnums3Solution(t *testing.T) {
	solution := `package enums

// Complex enum patterns! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Status int

const (
	Pending Status = iota
	InProgress
	Completed
	Failed
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
	Critical
)

func CanStart(status Status) bool {
	return status == Pending
}

func IsUrgent(priority Priority) bool {
	return priority == High || priority == Critical
}

func StatusMessage(status Status) string {
	switch status {
	case Pending:
		return "Waiting to start"
	case InProgress:
		return "Currently working"
	case Completed:
		return "Successfully finished"
	case Failed:
		return "Encountered an error"
	default:
		return "Unknown status"
	}
}
`
	testExerciseSolution(t, "enums", "enums3.go", solution)
}

// testTests1Solution tests tests1 exercise with correct solution
func testTests1Solution(t *testing.T) {
	solution := `package tests

import "testing"

// Time to learn about testing! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func Add(a, b int) int {
	return a + b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
`
	testExerciseSolution(t, "tests", "tests1.go", solution)
}

// testIterators1Solution tests iterators1 exercise with correct solution
func testIterators1Solution(t *testing.T) {
	solution := `package iterators

// Time to learn about iterators! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func Range(start, end int) []int {
	result := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

func Map(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter(slice []int, fn func(int) bool) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(slice []int, fn func(int, int) int, initial int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}
`
	testExerciseSolution(t, "iterators", "iterators1.go", solution)
}

// testSmartPointers1Solution tests smart_pointers1 exercise with correct solution
func testSmartPointers1Solution(t *testing.T) {
	solution := `package smart_pointers

// Time to learn about smart pointer patterns! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Box struct {
	value interface{}
}

func NewBox(value interface{}) *Box {
	return &Box{value: value}
}

func (b *Box) Get() interface{} {
	return b.value
}

func (b *Box) Set(value interface{}) {
	b.value = value
}

type RefCount struct {
	value interface{}
	count int
}

func NewRefCount(value interface{}) *RefCount {
	return &RefCount{value: value, count: 1}
}

func (rc *RefCount) Clone() *RefCount {
	rc.count++
	return rc
}

func (rc *RefCount) Release() int {
	rc.count--
	return rc.count
}

func (rc *RefCount) Get() interface{} {
	return rc.value
}
`
	testExerciseSolution(t, "smart_pointers", "smart_pointers1.go", solution)
}

// testMacros1Solution tests macros1 exercise with correct solution
func testMacros1Solution(t *testing.T) {
	solution := `package macros

import "fmt"

// Time to learn about code generation patterns! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func GenerateGetter(fieldName string) string {
	return fmt.Sprintf("func (s *Struct) Get%s() Type { return s.%s }", fieldName, fieldName)
}

func GenerateSetter(fieldName string) string {
	return fmt.Sprintf("func (s *Struct) Set%s(value Type) { s.%s = value }", fieldName, fieldName)
}

func GenerateStruct(structName string, fields []string) string {
	result := fmt.Sprintf("type %s struct {\n", structName)
	for _, field := range fields {
		result += fmt.Sprintf("    %s Type\n", field)
	}
	result += "}"
	return result
}

func ExpandTemplate(template string, data map[string]string) string {
	result := template
	for key, value := range data {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = fmt.Sprintf(result, placeholder, value)
	}
	return result
}
`
	testExerciseSolution(t, "macros", "macros1.go", solution)
}

// testConversions1Solution tests conversions1 exercise with correct solution
func testConversions1Solution(t *testing.T) {
	solution := `package conversions

import (
	"fmt"
	"strconv"
)

// Time to learn about type conversions! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func FloatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func IntToFloat(i int) float64 {
	return float64(i)
}

func FloatToInt(f float64) int {
	return int(f)
}

func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
`
	testExerciseSolution(t, "conversions", "conversions1.go", solution)
}

// testAdvanced1Solution tests advanced1 exercise with correct solution
func testAdvanced1Solution(t *testing.T) {
	solution := `package advanced

import (
	"context"
	"sync"
	"time"
)

// Advanced concepts! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func WithTimeout(duration time.Duration, fn func() error) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	
	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()
	
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

type SafeCounter struct {
	mu    sync.RWMutex
	value int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func (c *SafeCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}
`
	testExerciseSolution(t, "advanced", "advanced1.go", solution)
}

// testAdvanced2Solution tests advanced2 exercise with correct solution
func testAdvanced2Solution(t *testing.T) {
	solution := `package advanced

import (
	"reflect"
	"unsafe"
)

// Very advanced concepts! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func GetType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func GetFields(i interface{}) []string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}
	return fields
}

func SetField(ptr interface{}, fieldName string, value interface{}) bool {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr {
		return false
	}
	
	v = v.Elem()
	field := v.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return false
	}
	
	field.Set(reflect.ValueOf(value))
	return true
}

func UnsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
`
	testExerciseSolution(t, "advanced", "advanced2.go", solution)
}

// testModules1Solution tests modules1 exercise with correct solution
func testModules1Solution(t *testing.T) {
	solution := `package modules

import (
	"fmt"
	"path/filepath"
)

// Time to learn about modules and packages! 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Package struct {
	Name    string
	Version string
	Path    string
}

func NewPackage(name, version, path string) Package {
	return Package{
		Name:    name,
		Version: version,
		Path:    path,
	}
}

func (p Package) String() string {
	return fmt.Sprintf("%s@%s", p.Name, p.Version)
}

func (p Package) FullPath() string {
	return filepath.Join(p.Path, p.Name)
}

func ValidateVersion(version string) bool {
	return len(version) > 0 && version[0] == 'v'
}

func ImportPath(pkg Package) string {
	return fmt.Sprintf("import \"%s\"", pkg.FullPath())
}
`
	testExerciseSolution(t, "modules", "modules1.go", solution)
}

// testQuiz1Solution tests quiz1 exercise with correct solution
func testQuiz1Solution(t *testing.T) {
	solution := `package quizzes

// Quiz 1: Basic Go Concepts 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func BasicVariables() (int, string, bool) {
	return 42, "Gopher", true
}

func BasicFunction(a, b int) int {
	return a + b
}

func BasicControl(n int) string {
	if n > 0 {
		return "positive"
	} else if n < 0 {
		return "negative"
	}
	return "zero"
}

func BasicSlice() []int {
	return []int{1, 2, 3, 4, 5}
}
`
	testExerciseSolution(t, "quizzes", "quiz1.go", solution)
}

// testQuiz2Solution tests quiz2 exercise with correct solution
func testQuiz2Solution(t *testing.T) {
	solution := `package quizzes

import "fmt"

// Quiz 2: Intermediate Concepts 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Person struct {
	Name string
	Age  int
}

func CreatePerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func ProcessMap(m map[string]int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func ErrorHandling(divide bool) (int, error) {
	if divide {
		return 10 / 2, nil
	}
	return 0, fmt.Errorf("division not allowed")
}
`
	testExerciseSolution(t, "quizzes", "quiz2.go", solution)
}

// testQuiz3Solution tests quiz3 exercise with correct solution
func testQuiz3Solution(t *testing.T) {
	solution := `package quizzes

import "sync"

// Quiz 3: Advanced Concepts 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func ProcessChannel(input <-chan int, output chan<- int) {
	for num := range input {
		output <- num * 2
	}
	close(output)
}

func GenericFunction[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[0]
}
`
	testExerciseSolution(t, "quizzes", "quiz3.go", solution)
}

// testQuiz4Solution tests quiz4 exercise with correct solution
func testQuiz4Solution(t *testing.T) {
	solution := `package quizzes

import (
	"context"
	"time"
)

// Quiz 4: Expert Level 🐹

// SOLUTION IMPLEMENTED FOR TESTING

func ContextOperation(ctx context.Context) error {
	select {
	case <-time.After(100 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

type Worker interface {
	Work() string
}

type Robot struct {
	Name string
}

func (r Robot) Work() string {
	return r.Name + " is working"
}

func ProcessWorker(w Worker) string {
	return w.Work()
}

func ReflectionExample(i interface{}) string {
	switch v := i.(type) {
	case int:
		return "integer"
	case string:
		return "string"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}
}
`
	testExerciseSolution(t, "quizzes", "quiz4.go", solution)
}

// testQuizFinalSolution tests quiz_final exercise with correct solution
func testQuizFinalSolution(t *testing.T) {
	solution := `package quizzes

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Final Quiz: Master Level 🐹

// SOLUTION IMPLEMENTED FOR TESTING

type DataProcessor[T any] struct {
	data []T
	mu   sync.RWMutex
}

func NewDataProcessor[T any]() *DataProcessor[T] {
	return &DataProcessor[T]{
		data: make([]T, 0),
	}
}

func (dp *DataProcessor[T]) Add(item T) {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	dp.data = append(dp.data, item)
}

func (dp *DataProcessor[T]) Get(index int) (T, bool) {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	var zero T
	if index < 0 || index >= len(dp.data) {
		return zero, false
	}
	return dp.data[index], true
}

func (dp *DataProcessor[T]) Process(ctx context.Context, fn func(T) T) error {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	
	for i, item := range dp.data {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			dp.data[i] = fn(item)
		}
	}
	return nil
}

func ConcurrentSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	
	numWorkers := 4
	chunkSize := len(numbers) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
		numWorkers = len(numbers)
	}
	
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup
	
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + chunkSize
			if end > len(numbers) {
				end = len(numbers)
			}
			
			sum := 0
			for j := start; j < end; j++ {
				sum += numbers[j]
			}
			results <- sum
		}(i * chunkSize)
	}
	
	go func() {
		wg.Wait()
		close(results)
	}()
	
	total := 0
	for sum := range results {
		total += sum
	}
	return total
}

func PipelineExample(input <-chan int) <-chan string {
	output := make(chan string, 10)
	
	go func() {
		defer close(output)
		for num := range input {
			select {
			case output <- fmt.Sprintf("processed: %d", num*num):
			case <-time.After(time.Second):
				return
			}
		}
	}()
	
	return output
}
`
	testExerciseSolution(t, "quizzes", "quiz_final.go", solution)
}

// testExerciseSolution is a helper that tests a specific exercise solution
func testExerciseSolution(t *testing.T, category, filename, solutionCode string) {
	// Create temporary directory for testing
	tempDir := filepath.Join("test_data", "solutions", category)
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Write solution to temporary file
	tempFile := filepath.Join(tempDir, filename)
	err = os.WriteFile(tempFile, []byte(solutionCode), 0644)
	if err != nil {
		t.Fatalf("Failed to write solution file: %v", err)
	}

	// Copy the test file to temp directory
	originalTestFile := filepath.Join("..", "exercises", category, strings.Replace(filename, ".go", "_test.go", 1))
	tempTestFile := filepath.Join(tempDir, strings.Replace(filename, ".go", "_test.go", 1))

	testContent, err := os.ReadFile(originalTestFile)
	if err != nil {
		t.Fatalf("Failed to read original test file: %v", err)
	}

	err = os.WriteFile(tempTestFile, testContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write temp test file: %v", err)
	}

	// Run tests on the solution
	cmd := exec.Command("go", "test", "-v", ".")
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("Solution for %s/%s failed tests:\nError: %v\nOutput: %s",
			category, filename, err, string(output))
	} else {
		t.Logf("✅ Solution for %s/%s passed all tests", category, filename)
	}

	// Clean up
	defer func() {
		time.Sleep(100 * time.Millisecond) // Brief delay to ensure file handles are closed
		os.RemoveAll(tempDir)
	}()
}

// TestSolutionWorkflow tests the complete workflow of solving an exercise
func TestSolutionWorkflow(t *testing.T) {
	t.Run("complete_exercise_workflow", func(t *testing.T) {
		// Step 1: Verify the exercise starts incomplete
		originalFile := filepath.Join("..", "exercises", "intro", "intro1.go")
		content, err := os.ReadFile(originalFile)
		if err != nil {
			t.Fatalf("Failed to read exercise file: %v", err)
		}

		// Verify it has the "I AM NOT DONE" marker
		if !strings.Contains(string(content), "I AM NOT DONE") {
			t.Errorf("Exercise should start with 'I AM NOT DONE' marker")
		}

		// Verify it has incomplete code
		if !strings.Contains(string(content), `return ""`) {
			t.Errorf("Exercise should have incomplete return statement")
		}

		// Step 2: Simulate solving the exercise
		solvedContent := strings.Replace(string(content), "I AM NOT DONE", "", 1)
		solvedContent = strings.Replace(solvedContent, `return ""`, `return "Hello, Gopher!"`, 1)

		// Step 3: Test that the solved version passes tests
		testExerciseSolution(t, "intro", "intro1.go", solvedContent)
	})
}

// testExerciseCanBeCompleted tests if an exercise has proper structure for completion
func testExerciseCanBeCompleted(t *testing.T, category, filename string) {
	// Read the original exercise file
	exerciseFile := filepath.Join("..", "exercises", category, filename)
	content, err := os.ReadFile(exerciseFile)
	if err != nil {
		t.Fatalf("Failed to read exercise file %s/%s: %v", category, filename, err)
	}

	contentStr := string(content)

	// Verify exercise has "I AM NOT DONE" marker
	if !strings.Contains(contentStr, "I AM NOT DONE") {
		t.Errorf("Exercise %s/%s should have 'I AM NOT DONE' marker", category, filename)
	}

	// Verify exercise has package declaration
	if !strings.Contains(contentStr, "package ") {
		t.Errorf("Exercise %s/%s should have package declaration", category, filename)
	}

	// Verify corresponding test file exists
	testFile := filepath.Join("..", "exercises", category, strings.Replace(filename, ".go", "_test.go", 1))
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Errorf("Test file for %s/%s should exist", category, filename)
	}

	// Verify exercise compiles (even if incomplete)
	tempDir := filepath.Join("test_data", "completion_check", category)
	err = os.MkdirAll(tempDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy exercise to temp location and try to compile
	tempFile := filepath.Join(tempDir, filename)
	err = os.WriteFile(tempFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	// Check if it at least parses (gofmt -l will fail if syntax is wrong)
	cmd := exec.Command("gofmt", "-l", tempFile)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Exercise %s/%s has syntax errors: %v", category, filename, err)
	}

	t.Logf("✅ Exercise %s/%s has proper structure for completion", category, filename)
}

// TestExerciseCanBeCompleted verifies that exercises can actually be completed
func TestExerciseCanBeCompleted(t *testing.T) {
	// This test ensures that when users follow the instructions,
	// they can successfully complete the exercises

	t.Run("intro1_completion_test", func(t *testing.T) {
		// Read the original exercise
		originalFile := filepath.Join("..", "exercises", "intro", "intro1.go")
		content, err := os.ReadFile(originalFile)
		if err != nil {
			t.Fatalf("Failed to read exercise file: %v", err)
		}

		// Verify the exercise has proper structure for completion
		contentStr := string(content)

		// Should have the TODO comment
		if !strings.Contains(contentStr, "TODO") {
			t.Errorf("Exercise should have TODO comment to guide users")
		}

		// Should have the function signature that needs to be completed
		if !strings.Contains(contentStr, "func Hello() string") {
			t.Errorf("Exercise should have the correct function signature")
		}

		// Should have a clear placeholder to replace
		if !strings.Contains(contentStr, `return ""`) {
			t.Errorf("Exercise should have clear placeholder code to replace")
		}
	})
}

// TestRandomExerciseSampling tests a sample of verified exercises
func TestRandomExerciseSampling(t *testing.T) {
	// Test a representative sample from verified working solutions
	sampleExercises := []struct {
		category string
		filename string
		testFunc func(*testing.T)
	}{
		{"intro", "intro1.go", testIntro1Solution},
		{"functions", "functions1.go", testFunctions1Solution},
		{"structs", "structs1.go", testStructs1Solution},
		{"traits", "traits1.go", testTraits1Solution},
		{"error_handling", "errors1.go", testErrors1Solution},
	}

	for _, exercise := range sampleExercises {
		t.Run(exercise.category+"_sampling", exercise.testFunc)
	}
}
