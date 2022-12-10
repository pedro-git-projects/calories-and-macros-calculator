package calculator

import (
	"errors"
)

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

func (u User) CalculateBMR() (float64, error) {
	if u.Gender == male {
		return 66.5 + (13.75 * u.Weight) + (5.003 * u.Height) - (6.75 * float64(u.Age)), nil
	}
	if u.Gender == female {
		return 655.1 + (9.563 * u.Weight) + (1.850 * u.Height) - (4.676 * float64(u.Age)), nil
	}
	return -1, errors.New("unknown gender " + u.Gender.String())
}

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

func (u *User) calculateMacroSplit() {
	u.ProteinIntake = (u.Calories * 0.35) / float64(Protein.caloriesPerGram)
	u.FatIntake = (u.Calories * 0.15) / float64(Fat.caloriesPerGram)
	u.CarbohydrateIntake = (u.Calories * 0.50) / float64(Carbohydrate.caloriesPerGram)
}
