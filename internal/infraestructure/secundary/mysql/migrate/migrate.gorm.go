package migrate

import (
	"velocity-technical-test/internal/infraestructure/secundary/mysql"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"
	"velocity-technical-test/pkg/logger"
)

func Migrate() {
	log := logger.NewLogger()
	db := mysql.NewDBConnection().GetDB()
	err := db.AutoMigrate(
		&models.Order{},
		&models.OrderItem{},
		&models.Product{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}

	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count == 0 {
		db.Exec(`INSERT INTO products (name, price, stock) VALUES
			('All-Season Tires (Set of 4)', 499.99, 15),
			('Car Battery - 12V', 129.99, 25),
			('LED Headlights (Pair)', 79.99, 30),
			('Engine Oil - 5W-30 Synthetic', 34.99, 50),
			('Brake Pads (Set)', 49.99, 20),
			('Windshield Wipers (Pair)', 24.99, 40),
			('Car Floor Mats (Set of 4)', 59.99, 35),
			('Spark Plugs (Set of 4)', 29.99, 60),
			('Car Jack - Hydraulic', 89.99, 10),
			('Air Filter', 19.99, 75),
			('Fuel Filter', 14.99, 80),
			('Steering Wheel Cover', 12.99, 45),
			('Portable Tire Inflator', 39.99, 25),
			('Jump Starter Kit', 99.99, 15),
			('Roof Rack for SUVs', 199.99, 10),
			('Car Wax - 500ml', 19.99, 55),
			('GPS Navigation System', 129.99, 12),
			('Seat Covers (Set of 2)', 69.99, 20),
			('Car Cleaning Kit (7 Pieces)', 49.99, 40),
			('Bluetooth Car Adapter', 15.99, 100);`)
	}

	log.Success("Database migrated")
}
