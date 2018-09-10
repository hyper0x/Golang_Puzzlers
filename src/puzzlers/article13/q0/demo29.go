package main

import "fmt"

// 示例1。
// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

// 示例2。
type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}

// 示例3。
type Cat struct {
	name string
	Animal
}

// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func main() {
	// 示例1。
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)

	// 示例2。
	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)

	// 示例3。
	cat := Cat{
		name:   "little pig",
		Animal: animal,
	}
	fmt.Printf("The cat: %s\n", cat)
}
