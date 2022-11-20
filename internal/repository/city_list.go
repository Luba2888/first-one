package repository

import (
	"errors"
	"strconv"

	cities "github.com/kuzminprog/cities_information_service"
)

type CityListDB struct {
	db *DataBase
}

func NewCityListDB(db *DataBase) *CityListDB {
	return &CityListDB{db: db}
}

// Create places a new city in the database and assigns it an id.
// Returns the id of the city
func (r *CityListDB) Create(city cities.CityRequest) (string, error) {
	newID := r.newId()
	citySlice := []string{
		city.Name,
		city.Region,
		city.District,
		strconv.Itoa(city.Population),
		strconv.Itoa(city.Foundation),
	}
	r.db.records[newID] = citySlice
	return strconv.Itoa(newID), nil
}

// Delete deletes the city with the id from the database
func (r *CityListDB) Delete(id int) error {
	_, ok := r.db.records[id]
	if !ok {
		return errors.New(errNotFoundId)
	}

	delete(r.db.records, id)
	return nil
}

// SetPopulation updates the city population by id in the database
func (r *CityListDB) SetPopulation(id, population int) error {
	_, ok := r.db.records[id]
	if !ok {
		return errors.New(errNotFoundId)
	}

	r.db.records[id][populationIdx] = strconv.Itoa(population)
	return nil
}

// GetFromRegion returns the list of cities by region
func (r *CityListDB) GetFromRegion(region string) ([]string, error) {
	return r.findCities(regionIdx, region), nil
}

// GetFromDistrict returns the list of cities by district
func (r *CityListDB) GetFromDistrict(district string) ([]string, error) {
	return r.findCities(districtIdx, district), nil
}

// GetFromPopulation returns the list of cities by population
func (r *CityListDB) GetFromPopulation(population string) ([]string, error) {
	return r.findCities(populationIdx, population), nil
}

// GetFromFoundation returns the list of cities by foundation
func (r *CityListDB) GetFromFoundation(foundation string) ([]string, error) {
	return r.findCities(foundationIdx, foundation), nil
}

// GetFull searches for a city by id.
// If successful, it returns full information about the city
func (r *CityListDB) GetFull(id int) (*cities.City, error) {
	_, ok := r.db.records[id]
	if !ok {
		return nil, errors.New(errNotFoundId)
	}

	city := new(cities.City)
	city.Id = id
	city.Name = r.db.records[id][nameIdx]
	city.Region = r.db.records[id][regionIdx]
	city.District = r.db.records[id][districtIdx]

	population, err := strconv.Atoi(r.db.records[id][populationIdx])
	if err != nil {
		return nil, err
	}
	city.Population = population

	foundation, err := strconv.Atoi(r.db.records[id][foundationIdx])
	if err != nil {
		return nil, err
	}
	city.Foundation = foundation

	return city, nil
}

// newId returns a free identifier in the database
func (r *CityListDB) newId() int {
	r.db.lastID++
	return r.db.lastID
}

// findCities searches for cities based on idxType.
// Returns a list of found cities
func (r *CityListDB) findCities(idxType int, searchText string) []string {
	cityNames := make([]string, 0)
	for _, cityLine := range r.db.records {
		if cityLine[idxType] == searchText {
			cityNames = append(cityNames, cityLine[nameIdx])
		}
	}
	return cityNames
}
