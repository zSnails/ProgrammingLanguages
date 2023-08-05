package main

func main() {
	store := Store{}
	store.Add(
		NewFootwear("Nike", "Run To Zero", 60000, 42, 1),
		NewFootwear("Adidas", "Ultraboost 21", 56000, 42, 20),
		NewFootwear("Nike", "Air Force 1Low", 120000, 42, 30),
		NewFootwear("Converse", "Chuck Taylor All Star", 30000, 42, 30),
		NewFootwear("Puma", "RS-X³", 40000, 42, 15),
		NewFootwear("Reebok", "Classic Leather", 45000, 42, 50),
		NewFootwear("Vans", "Old Skool", 34000, 42, 10),
		NewFootwear("New Balance", "990v5", 23000, 30, 1),
		NewFootwear("ASICS", "Gel-Kayano 28", 45000, 39, 2),
		NewFootwear("Timberland", "6-Inch Premium Waterproof Boot", 30000, 40, 100),
		NewFootwear("Brooks", "Ghost 14", 34000, 40, 40),
	)

	// Ahora quiero vender las Nike Run To Zero
	element, err := store.FindByName("Nike Run To Zero")
	if err != nil {
		panic(err)
	}

	// Aquí todo bien
	err = store.Sell(element)
	if err != nil {
		panic(err)
	}

	// Aquí ya hay error
	err = store.Sell(element)
	if err != nil {
		panic(err)
	}

}
