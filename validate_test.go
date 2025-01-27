package validate_test

import (
	"errors"
	"testing"

	"github.com/craiggwilson/validate"
)

type testCase struct {
	name     string
	instance interface{}
	err      error
}

func TestValidate_And(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"string of length 3 and not empty (pass)",
			struct {
				Name string `validate:"len(3),notempty"`
			}{
				Name: "ABC",
			},
			nil,
		},
		{
			"string of length 3 and not empty (fail)",
			struct {
				Name string `validate:"len(3),notempty"`
			}{
				Name: "A",
			},
			errors.New(`"Name" must be of length 3`),
		},
		{
			"string of length 3 and empty",
			struct {
				Name string `validate:"len(3),empty"`
			}{
				Name: "A",
			},
			errors.New(`"Name" must be of length 3 and must be empty`),
		},
	})
}

func TestValidate_Equals(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"eq(3)"`
			}{
				Age: 3,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"eq(3)"`
			}{
				Age: 2,
			},
			errors.New(`"Age" must be equal to 3`),
		},
	})
}

func TestValidate_GreaterThan(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"gt(3)"`
			}{
				Age: 4,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"gt(3)"`
			}{
				Age: 2,
			},
			errors.New(`"Age" must be greater than 3`),
		},
	})
}

func TestValidate_GreaterThanOrEqual(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"gte(3)"`
			}{
				Age: 3,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"gte(3)"`
			}{
				Age: 2,
			},
			errors.New(`"Age" must be greater than or equal to 3`),
		},
	})
}

func TestValidate_In(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"in(1,2,3)"`
			}{
				Age: 2,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"in(1,2,3)"`
			}{
				Age: 4,
			},
			errors.New(`"Age" must be one of [1 2 3]`),
		},
	})
}

func TestValidate_Items(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"top-level pass",
			struct {
				Ages []int `validate:"items" validateItems:"gt(2)"`
			}{
				Ages: []int{3, 4, 5},
			},
			nil,
		},
		{
			"top-level fail single",
			struct {
				Ages []int `validate:"items" validateItems:"gt(3)"`
			}{
				Ages: []int{3, 4, 5},
			},
			errors.New(`"Ages" [0] must be greater than 3`),
		},
		{
			"top-level fail multiple",
			struct {
				Ages []int `validate:"items" validateItems:"gt(4)"`
			}{
				Ages: []int{3, 4, 5},
			},
			errors.New(`"Ages" [0] must be greater than 4 and [1] must be greater than 4`),
		},
		{
			"2nd-level pass",
			struct {
				Ages [][]int `validate:"items(vi)" vi:"items(vi2)" vi2:"gt(2)"`
			}{
				Ages: [][]int{{3, 4, 5}, {5, 7, 9}},
			},
			nil,
		},
		{
			"2nd-level fail single",
			struct {
				Ages [][]int `validate:"items(vi)" vi:"items(vi2)" vi2:"gt(3)"`
			}{
				Ages: [][]int{{9, 4, 5}, {3, 7, 9}},
			},
			errors.New(`"Ages" [1] [0] must be greater than 3`),
		},
		{
			"2nd-level fail multiple",
			struct {
				Ages [][]int `validate:"items(vi)" vi:"items(vi2)" vi2:"gt(3)"`
			}{
				Ages: [][]int{{3, 4, 5}, {4, 5, 2}},
			},
			errors.New(`"Ages" [0] [0] must be greater than 3 and [1] [2] must be greater than 3`),
		},
		{
			"top-level map pass",
			struct {
				Ages map[string]int `validate:"items" validateItems:"gt(2)"`
			}{
				Ages: map[string]int{"uno": 3, "dos": 4, "tres": 5},
			},
			nil,
		},
		{
			"top-level map fail single",
			struct {
				Ages map[string]int `validate:"items" validateItems:"gt(3)"`
			}{
				Ages: map[string]int{"uno": 3, "dos": 4, "tres": 5},
			},
			errors.New(`"Ages" [uno] must be greater than 3`),
		},
		{
			"top-level map fail multiple",
			struct {
				Ages map[string]int `validate:"items" validateItems:"gt(4)"`
			}{
				Ages: map[string]int{"uno": 3, "dos": 4, "tres": 5},
			},
			errors.New(`"Ages" [uno] must be greater than 4 and [dos] must be greater than 4`),
		},
	})
}

func TestValidate_Length(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"string of length 3 (pass)",
			struct {
				Name string `validate:"len(3)"`
			}{
				Name: "ABC",
			},
			nil,
		},
		{
			"string of length 3 (fail)",
			struct {
				Name string `validate:"len(3)"`
			}{
				Name: "A",
			},
			errors.New(`"Name" must be of length 3`),
		},
		{
			"string ptr of length 3 (pass)",
			struct {
				Name *string `validate:"len(3)"`
			}{
				Name: stringPtr("ABC"),
			},
			nil,
		},
		{
			"string ptr of length 3 (fail)",
			struct {
				Name *string `validate:"len(3)"`
			}{
				Name: stringPtr("A"),
			},
			errors.New(`"Name" must be of length 3`),
		},
		{
			"nil string ptr of length 3 (fail)",
			struct {
				Name *string `validate:"len(3)"`
			}{},
			errors.New(`"Name" must be of length 3`),
		},
	})
}

func TestValidate_LessThan(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"lt(3)"`
			}{
				Age: 2,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"lt(3)"`
			}{
				Age: 3,
			},
			errors.New(`"Age" must be less than 3`),
		},
	})
}

func TestValidate_LessThanOrEqual(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"lte(3)"`
			}{
				Age: 3,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"lte(3)"`
			}{
				Age: 4,
			},
			errors.New(`"Age" must be less than or equal to 3`),
		},
	})
}

func TestValidate_NotEmpty(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"not empty string",
			struct {
				Name string `validate:"notempty"`
			}{
				Name: "A",
			},
			nil,
		},
		{
			"nil string ptr",
			struct {
				Name *string `validate:"notempty"`
			}{},
			errors.New(`"Name" must not be empty`),
		},
		{
			"not empty string pointer",
			struct {
				Name *string `validate:"notempty"`
			}{
				Name: stringPtr("funny"),
			},
			nil,
		},
		{
			"empty string",
			struct {
				Name string `validate:"notempty~AHAHAHAH"`
			}{},
			errors.New(`AHAHAHAH`),
		},
		{
			"empty string ptr",
			struct {
				Name *string `validate:"notempty"`
			}{
				Name: stringPtr(""),
			},
			errors.New(`"Name" must not be empty`),
		},
		{
			"bool",
			struct {
				False *bool `validate:"notempty"`
			}{},
			errors.New("(notempty) only pointers to/or strings, arrays, slices, and maps are supported"),
		},
		{
			"bool ptr",
			struct {
				False *bool `validate:"notempty"`
			}{
				False: boolPtr(true),
			},
			errors.New("(notempty) only pointers to/or strings, arrays, slices, and maps are supported"),
		},
	})
}

func TestValidate_NotEqual(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"pass",
			struct {
				Age int `validate:"neq(3)"`
			}{
				Age: 2,
			},
			nil,
		},
		{
			"fail",
			struct {
				Age int `validate:"neq(3)"`
			}{
				Age: 3,
			},
			errors.New(`"Age" must not be equal to 3`),
		},
	})
}

func TestValidate_NotNil(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"nil pointer to a string",
			struct {
				Name *string `validate:"notnil"`
			}{},
			errors.New(`"Name" must not be nil`),
		},
		{
			"non-nil pointer to a string",
			struct {
				Name *string `validate:"notnil"`
			}{
				Name: stringPtr("funny"),
			},
			nil,
		},
		{
			"string",
			struct {
				Name string `validate:"notnil"`
			}{
				Name: "A",
			},
			errors.New("(notnil) only nil-able types are supported"),
		},
	})
}

func TestValidate_Or(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"string of length 3 and not empty or empty",
			struct {
				Name string `validate:"len(3),notempty|empty"`
			}{
				Name: "A",
			},
			errors.New(`"Name" must be of length 3 or must be empty`),
		},
		{
			"nil string ptr of length 3 or empty (pass)",
			struct {
				Name *string `validate:"nil|len(3)"`
			}{},
			nil,
		},
	})
}

type selfValidator struct {
	v int
}

func (sv *selfValidator) Validate(ctx validate.Context) error {
	if sv.v > 3 {
		return errors.New("AHAHAH")
	}

	return nil
}

func TestValidate_Self(t *testing.T) {
	runTestCases(t, []testCase{
		{
			"self validate success",
			selfValidator{
				v: 0,
			},
			nil,
		},
		{
			"self validate fail",
			selfValidator{
				v: 4,
			},
			errors.New("AHAHAH"),
		},
		{
			"self validate pointer success",
			&selfValidator{
				v: 0,
			},
			nil,
		},
		{
			"self validate pointer fail",
			&selfValidator{
				v: 4,
			},
			errors.New("AHAHAH"),
		},
	})
}

func TestValidate_StructCycle(t *testing.T) {
	type cycle struct {
		c *cycle `validate:"notnil, struct"`
	}

	runTestCases(t, []testCase{
		{
			"struct cycle",
			cycle{
				c: &cycle{},
			},
			errors.New(`"c" "c" must not be nil`),
		},
	})
}

func runTestCases(t *testing.T, testCases []testCase) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Validate(&tc.instance)
			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("expected error %v, but got %v", tc.err, err)
			} else if err != nil && tc.err == nil {
				t.Fatalf("expected no error, but got %v", err)
			} else if err == nil && tc.err != nil {
				t.Fatalf("expected error %v, but got none", tc.err)
			}
		})
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}
