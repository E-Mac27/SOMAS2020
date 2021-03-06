package rules

import "fmt"

// init Registers all veriables defined in Static variables list
func init() {
	for _, v := range StaticVariables {
		e := RegisterNewVariable(v)
		if e != nil {
			panic(fmt.Sprintf("variable registration gone wrong, variable: '%v' has been registered multiple times", v.VariableName))
		}
	}
}

// StaticVariables holds all globally defined variables
var StaticVariables = []VariableValuePair{
	{
		VariableName: "number_of_islands_contributing_to_common_pool",
		Values:       []float64{5},
	},
	{
		VariableName: "number_of_failed_forages",
		Values:       []float64{0.5},
	},
	{
		VariableName: "number_of_broken_agreements",
		Values:       []float64{1},
	},
	{
		VariableName: "max_severity_of_sanctions",
		Values:       []float64{2},
	},
	{
		VariableName: "no_islands_alive",
		Values:       []float64{6},
	},
	{
		VariableName: "no_ballots_cast",
		Values:       []float64{6},
	},
	{
		VariableName: "no_allocations_sent",
		Values:       []float64{6},
	},
	{
		VariableName: "islands_alive",
		Values:       []float64{0, 1, 2, 3, 4, 5},
	},
	{
		VariableName: "speakerSalary",
		Values:       []float64{50},
	},
	{
		VariableName: "judgeSalary",
		Values:       []float64{50},
	},
	{
		VariableName: "presidentSalary",
		Values:       []float64{50},
	},
	{
		VariableName: "expected_tax_contribution",
		Values:       []float64{0},
	},
	{
		VariableName: "expected_allocation",
		Values:       []float64{0},
	},
	{
		VariableName: "island_tax_contribution",
		Values:       []float64{0},
	},
	{
		VariableName: "island_allocation",
		Values:       []float64{0},
	},
}
