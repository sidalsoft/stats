package stats

import "github.com/sidalsoft/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		sum += p.Amount
	}
	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		if p.Category != category {
			continue
		}
		sum += p.Amount
	}
	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	r := make(map[types.Category]types.Money)
	for _, payment := range payments {
		if _, ok := r[payment.Category]; ok {
			r[payment.Category] += payment.Amount
		} else {
			r[payment.Category] = payment.Amount
		}
	}
	return r
}
