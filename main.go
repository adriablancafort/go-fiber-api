package main

import (
    "github.com/gofiber/fiber/v2"
)

type Product struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// In-memory storage
var products = make(map[string]*Product)

func main() {
    app := fiber.New()

    app.Get("/products", getAllProducts)
    app.Get("/products/:id", getProduct)
    app.Post("/products", createProduct)
    app.Put("/products/:id", updateProduct)
    app.Delete("/products/:id", deleteProduct)

    app.Listen(":3000")
}

// Get all products
func getAllProducts(c *fiber.Ctx) error {
    productList := make([]*Product, 0, len(products))
    for _, product := range products {
        productList = append(productList, product)
    }
    return c.JSON(productList)
}

// Get single product
func getProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    if product, exists := products[id]; exists {
        return c.JSON(product)
    }
    return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
}

// Create new product
func createProduct(c *fiber.Ctx) error {
    product := new(Product)
    if err := c.BodyParser(product); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }
    products[product.ID] = product
    return c.Status(201).JSON(product)
}

// Update product
func updateProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    if _, exists := products[id]; !exists {
        return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
    }

    product := new(Product)
    if err := c.BodyParser(product); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }
    product.ID = id
    products[id] = product
    return c.JSON(product)
}

// Delete product
func deleteProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    if _, exists := products[id]; !exists {
        return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
    }
    delete(products, id)
    return c.SendStatus(204)
}