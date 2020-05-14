// Create a data structure to pass to a template which
// contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

// Modify this data structure to hold menu information for an unlimited number of restaurants.

package main

type restaurant struct {
	Name string
	Menu menu
}

type menu []meal

type meal struct {
	Meal   string
	Dishes []dish
}

type dish struct {
	Dish  string
	Price float64
}

var restaurants = []restaurant{
	steakhouse, pancakehouse,
}

var steakhouse = restaurant{
	"The Meats",
	steakhousemenu,
}

var pancakehouse = restaurant{
	"The Pancake House",
	pancakehousemenu,
}

var steakhousemenu = []meal{
	breakfast,
	lunch,
	dinner,
}

var pancakehousemenu = []meal{
	breakfast,
	lunch,
	dinner,
}

var breakfast = meal{
	"breakfast",
	breakfasts,
}

var lunch = meal{
	"lunch",
	lunches,
}

var dinner = meal{
	"dinner",
	dinners,
}

var breakfasts = []dish{
	{
		"pancakes", 4.00,
	},
	{
		"eggs", 5.00,
	},
	{
		"bacon", 3.00,
	},
	{
		"fruit", 5.50,
	},
	{
		"omellette", 3.00,
	},
	{
		"oatmeal", 5.50,
	},
}

var lunches = []dish{
	{
		"hoagie", 4.00,
	},
	{
		"tuna salad", 5.00,
	},
	{
		"hamburger", 3.00,
	},
	{
		"fries", 5.50,
	},
}

var dinners = []dish{
	{
		"steak", 4.00,
	},
	{
		"cobb salad", 5.00,
	},
	{
		"pasta primavera", 3.00,
	},
	{
		"clam chowder", 5.50,
	},
}
