package calculator

import (
	"errors"
)

// The User type stores both user provided data
// application calculated data about the user
type User struct {
	Age                int
	Height             float64
	Weight             float64
	BMR                float64
	Calories           float64
	Gender             Gender
	ActivityLevel      ActivityLevel
	CarbohydrateIntake float64
	ProteinIntake      float64
	FatIntake          float64
}

// NewUser returns a new user instance with all fields calculated
// NewUser can fail and return an error
func NewUser(age int, height float64, weight float64, gender Gender, activity ActivityLevel) (*User, error) {
	u := User{
		Age:           age,
		Height:        height,
		Weight:        weight,
		Gender:        gender,
		ActivityLevel: activity,
	}

	var err error

	u.BMR, err = u.CalculateBMR()
	if err != nil {
		return nil, err
	}

	u.Calories, err = u.CalculateCalories()
	if err != nil {
		return nil, err
	}

	u.calculateMacroSplit()

	return &u, nil
}

// CalculateBMR calculates the user BMR according to the especified gender
// if the gneder passed is invalid it returns an error
func (u User) CalculateBMR() (float64, error) {
	if u.Gender == male {
		return 66.5 + (13.75 * u.Weight) + (5.003 * u.Height) - (6.75 * float64(u.Age)), nil
	}
	if u.Gender == female {
		return 655.1 + (9.563 * u.Weight) + (1.850 * u.Height) - (4.676 * float64(u.Age)), nil
	}
	return -1, errors.New("unknown gender " + u.Gender.String())
}

// CalculateCalories calculates calories according to the especified activity level
// if the activity level is invalid it returns a non nil error
func (u User) CalculateCalories() (float64, error) {
	switch u.ActivityLevel {
	case bmr:
		return u.BMR, nil
	case sedentary:
		return u.BMR * 1.2, nil
	case light:
		return u.BMR * 1.375, nil
	case moderate:
		return u.BMR * 1.55, nil
	case active:
		return u.BMR * 1.75, nil
	case veryActive:
		return u.BMR * 1.9, nil
	default:
		return -1, errors.New("invalid activity level" + u.ActivityLevel.String())

	}
}

// calculateMacroSplit is a mutator on the User type which populates the
// portein, fat and carb intake fields
func (u *User) calculateMacroSplit() {
	u.ProteinIntake = (u.Calories * 0.35) / float64(Protein.caloriesPerGram)
	u.FatIntake = (u.Calories * 0.20) / float64(Fat.caloriesPerGram)
	u.CarbohydrateIntake = (u.Calories * 0.45) / float64(Carbohydrate.caloriesPerGram)
}
