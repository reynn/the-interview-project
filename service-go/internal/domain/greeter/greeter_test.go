package greeter

import (
	"fmt"
	"testing"
)

func TestGreeter(t *testing.T) {
	type TestResults struct {
		greeting string
	}

	type TestParameters struct {
		greeterName string
		beginFn     func(t *testing.T)
		expectFn    func(t *testing.T, results TestResults)
	}

	type TestScenario struct {
		description string
		setupFn     func(t *testing.T) TestParameters
	}

	testScenarios := []TestScenario{
		{
			description: "formats greeting with given name",
			setupFn: func(t *testing.T) TestParameters {
				greeterName := "foo"
				return TestParameters{
					greeterName: greeterName,
					beginFn:     func(t *testing.T) {},
					expectFn: func(t *testing.T, results TestResults) {
						expectedGreeting := fmt.Sprintf(greetingFormatter, greeterName)
						if results.greeting != expectedGreeting {
							t.Errorf(
								"Unexpected greeting: \"%s\" != \"%s\"",
								results.greeting,
								expectedGreeting,
							)
						}
					},
				}
			},
		},
		{
			description: "formats greeting with \"stranger\" when no name is given",
			setupFn: func(t *testing.T) TestParameters {
				greeterName := ""
				return TestParameters{
					greeterName: greeterName,
					beginFn:     func(t *testing.T) {},
					expectFn: func(t *testing.T, results TestResults) {
						expectedGreeting := fmt.Sprintf(greetingFormatter, "stranger")
						if results.greeting != expectedGreeting {
							t.Errorf(
								"Unexpected greeting: \"%s\" != \"%s\"",
								results.greeting,
								expectedGreeting,
							)
						}
					},
				}
			},
		},
	}

	for _, scenario := range testScenarios {
		parameters := scenario.setupFn(t)
		t.Run(
			scenario.description,
			func(t *testing.T) {
				parameters.beginFn(t)

				result := Greet(parameters.greeterName)

				parameters.expectFn(
					t,
					TestResults{
						greeting: result,
					},
				)
			},
		)
	}
}
