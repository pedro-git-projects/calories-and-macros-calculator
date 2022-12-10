package calculator

import "errors"

type User struct {
	Age           int
	Height        float64
	Weight        float64
	BMR           float64
	Calories      float64
	Gender        Gender
	ActivityLevel ActivityLevel
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

	return &u, nil
}

func (u User) CalculateBMR() (float64, error) {
	if u.Gender == Male {
		return 66.5 + (13.75 * u.Weight) + (5.003 * u.Height) - (6.75 * float64(u.Age)), nil
	}
	if u.Gender == Female {
		return 655.1 + (9.563 * u.Weight) + (1.850 * u.Height) - (4.676 * float64(u.Age)), nil
	}
	return 0, errors.New("unknown gender " + u.Gender.String())
}

func (u User) CalculateCalories() (float64, error) {
	switch u.ActivityLevel {
	case BMR:
		return u.BMR, nil
	case Sedentary:
		return u.BMR * 1.2, nil
	case Light:
		return u.BMR * 1.375, nil
	case Moderate:
		return u.BMR * 1.55, nil
	case Active:
		return u.BMR * 1.75, nil
	case VeryActive:
		return u.BMR * 1.9, nil
	default:
		return 0, errors.New("invalid activity level" + u.ActivityLevel.String())

	}
}
