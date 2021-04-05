package stats

import (
	"fmt"
	"github.com/sidalsoft/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 1_000_00,
			Status: types.StatusFail,
		},
		{
			Amount: 1_000_00,
		},
		{
			Amount: 1_000_00,
		},
		{
			Amount: 1_000_00,
		},
	}
	fmt.Println(Avg(payments))
	//Output: 100000
}

func ExampleTotalInCategoryAvg() {
	payments := []types.Payment{
		{
			Amount:   2_000_00,
			Category: "salom",
			Status:   types.StatusFail,
		},
		{
			Amount:   1_000_00,
			Category: "alek",
		},
		{
			Amount:   5_000_00,
			Category: "salom",
		},
		{
			Amount:   1_000_00,
			Category: "paka",
		},
	}
	fmt.Println(TotalInCategory(payments, "salom"))
	//Output: 500000
}
