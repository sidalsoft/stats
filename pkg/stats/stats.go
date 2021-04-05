package stats

import "github.com/sidalsoft/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		//if p.Status == types.StatusFail {
		//	continue
		//}
		sum += p.Amount
	}
	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		//if p.Status == types.StatusFail {
		//	continue
		//}
		if p.Category != category {
			continue
		}
		sum += p.Amount
	}
	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	r := make(map[types.Category]types.Money)
	filter := make(map[types.Category][]types.Payment)
	for _, payment := range payments {
		if _, ok := filter[payment.Category]; ok {
			filter[payment.Category] = append(filter[payment.Category], payment)
		} else {
			filter[payment.Category] = []types.Payment{payment}
		}
	}
	f := func(payments []types.Payment) (types.Money, int) {
		count := 0
		m := types.Money(0)
		for _, payment := range payments {
			count++
			m += payment.Amount
		}
		return m, count
	}

	for k, v := range filter {
		m, c := f(v)
		r[k] = m / types.Money(c)
	}
	return r
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	r := map[types.Category]types.Money{}
	for k, v := range second {
		r[k] = v - first[k]
	}
	for k, v := range first {
		if _, ok := second[k]; !ok {
			r[k] = -v
		}
	}
	return r
}
