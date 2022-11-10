package main

import "Spotlas/Database"

func main() {
	db := Database.DBSetup()
	//Database.Question1(db)
	Database.Q1(db)
}
