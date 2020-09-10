package services

import (
	"sort"
	"sync"
	"errors"

	"demo-user/dao"
	"demo-user/models"
)

// UserList ...
func UserList() ([]models.UserDetail, error) {
	var (
		result = make([]models.UserDetail, 0)
		wg     sync.WaitGroup
	)

	// Find
	users, err := dao.UserList()
	total := len(users)

	// Return if not found
	if total == 0 {
		return result, err
	}

	// Add process
	wg.Add(total)

	for index := range users {
		go func(index int) {
			defer wg.Done()

			// Convert to UserDetail
			user := convertToUserDetail(users[index])

			// Append
			result = append(result, user)
		}(index)
	}

	// Wait process
	wg.Wait()

	// Sort again
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})
	return result, err
}

// UserCreate ...
func UserCreate(body models.UserCreatePayload) (models.UserBSON, error) {
	var (
		user = userCreatePayloadToBSON(body)
	)

	//Create user
	doc, err := dao.UserCreate(user)
	if err != nil {
		err = errors.New("Khong the tao user")
		return doc, err
	}

	return doc, err
}
