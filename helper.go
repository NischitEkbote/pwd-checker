package main

import "unicode"

type Result struct {
	Score    int
	Strength string
	Reasons  []string
}

func checkPassword(pwd string) Result {
	score := 0
	reasons := []string{}

	if len(pwd) >= 8 {
		score++
	} else {
		reasons = append(reasons, "Password should be at least 8 Characters")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, c := range pwd {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if hasUpper {
		score++
	} else {
		reasons = append(reasons, "Password should contain at least one uppercase letter")
	}
	if hasLower {
		score++
	} else {
		reasons = append(reasons, "Password should contain at least one lowercase letter")
	}
	if hasNumber {
		score++
	} else {
		reasons = append(reasons, "Password should contain at least one number")
	}
	if hasSpecial {
		score++
	} else {
		reasons = append(reasons, "Password should contain at least one special character")
	}

	return Result{
		Score:    score,
		Strength: getStrength(score),
		Reasons:  reasons,
	}
}

func getStrength(score int) string {
	switch {
	case score <= 2:
		return "Weak"
	case score <= 4:
		return "Medium"
	default:
		return "Strong"
	}
}
