package dblayer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetBookInfoById(t *testing.T) {

	assert := assert.New(t)
	db, err := NewORM("test", gorm.Config{})

	assert.NoError(err)

	result, err := db.GetBookInfoById(1)

	assert.NoError(err)
	assert.NotEmpty(result)

	fmt.Printf("%+v\n", result)

}

func TestGetBookInfoWithRank(t *testing.T) {
	assert := assert.New(t)
	db, err := NewORM("test", gorm.Config{})

	assert.NoError(err)

	result, err := db.GetBookInfoWithRank(1)

	assert.NoError(err)
	assert.NotEmpty(result)

	for _, rank := range result {

		fmt.Printf("%+v\n", rank)
	}
}

func TestGetBookInfoWithPoint(t *testing.T) {
	assert := assert.New(t)
	db, err := NewORM("test", gorm.Config{})

	assert.NoError(err)

	pointList, err := db.GetBookInfoWithPoint(1)

	assert.NoError(err)
	assert.NotEmpty(pointList)

	for _, point := range pointList {
		fmt.Printf("%+v\n", point)
	}
}

func TestGetBookInfoWithPrice(t *testing.T) {
	assert := assert.New(t)
	db, err := NewORM("test", gorm.Config{})

	assert.NoError(err)

	pointList, err := db.GetBookInfoWithPrice(1)

	assert.NoError(err)
	assert.NotEmpty(pointList)

	for _, point := range pointList {
		fmt.Printf("%+v\n", point)
	}
}

func TestGetBookInfoWithSummary(t *testing.T) {
	assert := assert.New(t)
	db, err := NewORM("test", gorm.Config{})

	assert.NoError(err)

	pointsummary, err := db.GetBookInfoWithSummary(1)

	assert.NoError(err)
	assert.NotEmpty(pointsummary)

	fmt.Printf("%+v\n", pointsummary)
}
