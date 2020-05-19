package main

type AllProductWithCategory map[string]interface{}

func getAllProducts() ([]AllProductWithCategory, error) {
	db := connect()
	defer db.Close()

	rows, err := db.Query(`
			select p.id,p.product_code,p.product_name,c.id as ids,c.category_name 
			from product p join category c 
			where p.category_id=c.id`)
	cols, _ := rows.Columns()

	if err != nil {
		return nil, err
	}

	var results []AllProductWithCategory

	for rows.Next() {
		current := makeResultReceiver(len(cols))

		if err := rows.Scan(current...); err != nil {
			return nil, err
		}

		result := make(AllProductWithCategory)

		for i, colName := range cols {
			val := current[i].(*interface{})

			result[colName] = string((*val).([]byte))
		}
		results = append(results, result)
	}
	return results, nil
}
