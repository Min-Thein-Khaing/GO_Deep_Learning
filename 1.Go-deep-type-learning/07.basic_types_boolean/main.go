package main

import (
	"fmt"
)

func main() {
	isLogging := true;
	isAdmin := false;
	hasScription := true;

	canOpenDashboard := isLogging && hasScription

	canDeletePost := isAdmin || (isLogging && hasScription)

	isAdmin = true;
	fmt.Println("Dashboard:", canOpenDashboard)
	fmt.Println("Delete Post:", canDeletePost)
}
