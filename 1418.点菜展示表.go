/*
 * @lc app=leetcode.cn id=1418 lang=golang
 *
 * [1418] 点菜展示表
 */

// @lc code=start
func displayTable(orders [][]string) [][]string {
	// table : food : num
	var tm = make(map[string]map[string]int)
	var tables []string
	var foodsm = make(map[string]int)

	for _, order := range orders {
		// 用户名，table名，菜品名
		if t, ok := tm[order[1]]; ok {
			// food is in ?
			if _, ok := t[order[2]]; ok {
				// num ++
				tm[order[1]][order[2]]++
			} else {
				tm[order[1]][order[2]] = 1
			}
		} else {
			tm[order[1]] = map[string]int{
				order[2]: 1,
			}
			tables = append(tables, order[1])
		}
		foodsm[order[2]] = 0
	}
	var foods []string
	var result [][]string
	var first []string

	for food, _ := range foodsm {
		foods = append(foods, food)
	}
	sort.Strings(foods)
	first = append(first, "Table")
	for _, food := range foods {
		first = append(first, food)
	}
	result = append(result, first)

	sort.Slice(tables, func(i, j int) bool {
		t1, _ := strconv.Atoi(tables[i])
		t2, _ := strconv.Atoi(tables[j])
		return t1 < t2
	})
	for _, t := range tables {
		var col []string
		col = append(col, t)
		for _, f := range foods {
			if num, ok := tm[t][f]; ok {
				col = append(col, strconv.Itoa(num))
			} else {
				col = append(col, "0")
			}
		}
		result = append(result, col)
	}
	return result

}

// @lc code=end

