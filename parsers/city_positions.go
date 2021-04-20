package main

import (
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"golang.org/x/net/html/charset"
)

type vertex struct {
	x float64
	y float64
}

func convertToUtf(s string, encoding string) []byte {
	r, err := charset.NewReader(strings.NewReader(s), encoding)
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

/*
func main() {
	file, err := os.Open("../sources/positions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	provinces := make([]string, 0)
	coordinates := make([]vertex, 0)
	for scanner.Scan() {
		if (i-3)%12 == 0 {
			coordinate := strings.Split(strings.TrimSpace(scanner.Text()), " ")
			x, _ := strconv.ParseFloat(coordinate[0], 8)
			y, _ := strconv.ParseFloat(coordinate[1], 8)
			coordinates = append(coordinates, vertex{x, y})
		}
		if i == 0 || i%12 == 0 {
			provinces = append(provinces, strings.TrimSpace(scanner.Text())[1:])
		}
		//if i > 200 {
		//break
		//}
		i++
	}

	//write to the database
	dbPass := os.Getenv("db_pass")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgresql://localhost/eu4?user=postgres&password=%s&sslmode=disable", dbPass))
	for index, val := range provinces {

		db.MustExec(`DELETE FROM province.coordinates
					where province_name=$1`, convertToUtf(val, "cp1250"))
		db.MustExec(`INSERT INTO province.coordinates
					(province_name, x_coordinate, y_coordinate)
						VALUES ($1, $2, $3)`,
			convertToUtf(val, "cp1250"), coordinates[index].x, coordinates[index].y)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
*/
