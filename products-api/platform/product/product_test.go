package product

import "testing"

func TestAdd(t *testing.T) {
	product := New()
	product.Add(Item{
		"B01DMYYCD6",
		"Canon EOS 80D DSLR Camera with EF-S 18-55mm f/3.5-5.6 IS STM Lens, Total Of 48GB SDHC along with Deluxe accessory bundle",
		"4534",
		"Canon",
		"176",
		"Cameras",
	})
	if len(product.Products) != 1 {
		t.Errorf("Item was not Added")
	}
}

func TestGetAll(t *testing.T) {
	product := New()
	product.Add(Item{
		"B01DMYYCD6",
		"Canon EOS 80D DSLR Camera with EF-S 18-55mm f/3.5-5.6 IS STM Lens, Total Of 48GB SDHC along with Deluxe accessory bundle",
		"4534",
		"Canon",
		"176",
		"Cameras",
	})
	results := product.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not Added")
	}

}

func TestFindCorrectInput(t *testing.T) {
	products := New()
	products.Add(Item{"B01DMYYCD6",
		"Canon EOS 80D DSLR Camera with EF-S 18-55mm f/3.5-5.6 IS STM Lens, Total Of 48GB SDHC along with Deluxe accessory bundle",
		"4534",
		"Canon",
		"176",
		"Cameras",
	})
	products.Add(Item{"B001EQSFRE",
		"Canon Staples P1 (2x5000)",
		"4534",
		"Canon",
		"191",
		"Home Audio",
	})
	products.Add(Item{"B004H9PAO6",
		"Brother TN330 Toner, Standard Yield (Reseller Offer)",
		"4053",
		"Brother",
		"200",
		"Printers & Scanners",
	})

	values := []string{"Canon"}
	results, err := products.Find("brandName", values)
	if err != nil || len(results) != 2 {
		t.Errorf("Find function returned an error or the wrong results")
	}

}

func TestFindIncorrectInput(t *testing.T) {
	products := New()
	products.Add(Item{"B01DMYYCD6",
		"Canon EOS 80D DSLR Camera with EF-S 18-55mm f/3.5-5.6 IS STM Lens, Total Of 48GB SDHC along with Deluxe accessory bundle",
		"4534",
		"Canon",
		"176",
		"Cameras",
	})
	products.Add(Item{"B001EQSFRE",
		"Canon Staples P1 (2x5000)",
		"4534",
		"Canon",
		"191",
		"Home Audio",
	})
	products.Add(Item{"B004H9PAO6",
		"Brother TN330 Toner, Standard Yield (Reseller Offer)",
		"4053",
		"Brother",
		"200",
		"Printers & Scanners",
	})

	values := []string{"Canon"}
	_, err := products.Find("dummyName", values)
	if err == nil {
		t.Errorf("Find Function was unable to throw error when incorrect input was given")
	}

}

func TestReturnPageCorrectInput(t *testing.T) {
	products := New()
	products.Add(Item{"B01DMYYCD6",
		"Canon EOS 80D DSLR Camera with EF-S 18-55mm f/3.5-5.6 IS STM Lens, Total Of 48GB SDHC along with Deluxe accessory bundle",
		"4534",
		"Canon",
		"176",
		"Cameras",
	})
	products.Add(Item{"B001EQSFRE",
		"Canon Staples P1 (2x5000)",
		"4534",
		"Canon",
		"191",
		"Home Audio",
	})
	products.Add(Item{"B004H9PAO6",
		"Brother TN330 Toner, Standard Yield (Reseller Offer)",
		"4053",
		"Brother",
		"200",
		"Printers & Scanners",
	})

	values := []string{"Canon"}
	results, err := products.Find("brandName", values)
	page, err := products.ReturnPage(results, 1, 1)
	if err != nil || len(page) != 1 {
		t.Errorf("ReturnPage Function had an error occur or the output was incorrect")
	}

}

func TestReturnPageIncorrectInput(t *testing.T) {
	products := New()
	products.Add(Item{"B01DMYYCD6",
		"Canon EOS 80D DSLR Camera with EF-S 18-55mm f/3.5-5.6 IS STM Lens, Total Of 48GB SDHC along with Deluxe accessory bundle",
		"4534",
		"Canon",
		"176",
		"Cameras",
	})
	products.Add(Item{"B001EQSFRE",
		"Canon Staples P1 (2x5000)",
		"4534",
		"Canon",
		"191",
		"Home Audio",
	})
	products.Add(Item{"B004H9PAO6",
		"Brother TN330 Toner, Standard Yield (Reseller Offer)",
		"4053",
		"Brother",
		"200",
		"Printers & Scanners",
	})

	values := []string{"Canon"}
	results, err := products.Find("brandName", values)
	_, err = products.ReturnPage(results, 100, 1)
	if err == nil {
		t.Errorf("ReturnPage Function wasn't able to correctly throw an error when input was invalid")
	}

}
