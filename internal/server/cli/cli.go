package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang/internal/models"
	"golang/internal/service"
)

// CLI handles command-line interface (Presentation Layer)
type CLI struct {
	service *service.Service
	scanner *bufio.Scanner
}

func NewCLI(svc *service.Service) *CLI {
	return &CLI{
		service: svc,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *CLI) Start() error {
	ctx := context.Background()

	fmt.Println("========================================")
	fmt.Println("   Contact Management System - CLI")
	fmt.Println("========================================")

	for {
		fmt.Println("1. List contacts")
		fmt.Println("2. Create contact")
		fmt.Println("3. Update contact")
		fmt.Println("4. Delete contact")
		fmt.Println("0. Exit")
		fmt.Print("\nChoice: ")

		choice := c.readInput()

		switch choice {
		case "1":
			c.listContacts(ctx)
		case "2":
			c.createContact(ctx)
		case "3":
			c.updateContact(ctx)
		case "4":
			c.deleteContact(ctx)
		case "0":
			fmt.Println("Goodbye!")
			return nil
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println()
	}
}

func (c *CLI) listContacts(ctx context.Context) {
	contacts, err := c.service.ContactService.GetAll(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\nContacts:")
	for _, contact := range contacts {
		fmt.Printf("  [%d] %s - %s\n", contact.ID, contact.FullName(), contact.Email)
	}
}

func (c *CLI) createContact(ctx context.Context) {
	fmt.Print("First Name: ")
	firstName := c.readInput()
	fmt.Print("Last Name: ")
	lastName := c.readInput()
	fmt.Print("Email: ")
	email := c.readInput()

	contact := models.Contact{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	created, err := c.service.ContactService.Create(ctx, contact)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Contact created with ID %d\n", created.ID)
}

func (c *CLI) updateContact(ctx context.Context) {
	fmt.Print("Contact ID: ")
	idStr := c.readInput()
	id, _ := strconv.Atoi(idStr)

	fmt.Print("First Name: ")
	firstName := c.readInput()
	fmt.Print("Last Name: ")
	lastName := c.readInput()
	fmt.Print("Email: ")
	email := c.readInput()

	contact := models.Contact{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if err := c.service.ContactService.UpdateAndNotify(ctx, contact); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("✓ Contact updated and notification sent")
}

func (c *CLI) deleteContact(ctx context.Context) {
	fmt.Print("Contact ID: ")
	idStr := c.readInput()
	id, _ := strconv.Atoi(idStr)

	if err := c.service.ContactService.Delete(ctx, id); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("✓ Contact deleted")
}

func (c *CLI) readInput() string {
	c.scanner.Scan()
	return strings.TrimSpace(c.scanner.Text())
}
