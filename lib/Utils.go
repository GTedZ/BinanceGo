package lib

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var utils Utils

type Utils struct{}

type CustomInterval struct {
	Rune          rune
	IntervalValue int64
}

type KlineInterval struct {
	Name       string
	Rune       rune
	BaseValue  int64
	Multiplier int64
	Value      int64
}

func (*Utils) GetIntervalFromString(intervalStr string, customIntervals ...CustomInterval) (interval *KlineInterval, isComplex bool, err error) {
	// Fetch the last character
	intervalRune := rune(intervalStr[len(intervalStr)-1])

	// Parse the rest of the string (excluding the last character) as an integer
	restOfString := intervalStr[:len(intervalStr)-1]

	var parseErr error
	multiplier, parseErr := utils.ParseInt(restOfString)
	if parseErr != nil {
		return nil, false, parseErr
	}

	if multiplier <= 0 {
		return nil, false, fmt.Errorf("multiplier value must be positive")
	}

	var baseValue int64

	found := false
	for _, customInterval := range customIntervals {
		if customInterval.IntervalValue <= 0 {
			return nil, false, fmt.Errorf("custom interval rune '%v' must have a positive value", intervalRune)
		}
		if customInterval.Rune == intervalRune {
			found = true
			baseValue = customInterval.IntervalValue
		}
	}

	if !found {
		switch intervalRune {
		case 's':
			baseValue = 1000
		case 'm':
			baseValue = 60 * 1000
		case 'h':
			baseValue = 60 * 60 * 1000
		case 'd':
			baseValue = 24 * 60 * 60 * 1000
		case 'w':
			baseValue = 7 * 24 * 60 * 60 * 1000

		default:
			return nil, true, fmt.Errorf("simple interval rune '%v' doesn't exit/not supported", intervalRune)
		}
	}

	interval = &KlineInterval{
		Name:       intervalStr,
		Rune:       intervalRune,
		BaseValue:  baseValue,
		Multiplier: multiplier,
		Value:      baseValue * multiplier,
	}

	return interval, false, nil
}

func (*Utils) GetOpenCloseTimes(currentTime int64, interval string, customIntervals ...CustomInterval) (openTime int64, closeTime int64, err error) {

	klineInterval, isComplex, err := utils.GetIntervalFromString(interval)
	if err != nil {
		return 0, 0, err
	}

	if !isComplex {
		openTime = currentTime - (currentTime % klineInterval.Value)
		closeTime = openTime + klineInterval.Value - 1
		return openTime, closeTime, nil
	}

	baseUnix_time_obj := time.Unix(0, 0)
	current_time_obj := time.Unix(0, currentTime*int64(time.Millisecond))

	switch klineInterval.Rune {
	case 'w':
		const WEEK_millis = 7 * 24 * 60 * 60 * 1000
		weekDay_offset := int(current_time_obj.Weekday() - time.Monday)
		monday_time := current_time_obj.AddDate(0, 0, weekDay_offset).UnixMilli()

		unixFirstWeek_Dayoffset := int(baseUnix_time_obj.Weekday() - time.Monday)
		unixFirstWeek_date := baseUnix_time_obj.AddDate(0, 0, unixFirstWeek_Dayoffset)
		unixFirstWeek_offset := unixFirstWeek_date.UnixMilli()

		timestamp_to_check := monday_time + unixFirstWeek_offset

		openTime = timestamp_to_check - (timestamp_to_check % (WEEK_millis * klineInterval.Multiplier))
		closeTime = openTime + (WEEK_millis * klineInterval.Multiplier) - 1

		return openTime - unixFirstWeek_offset, closeTime - unixFirstWeek_offset, nil
	case 'M':
		yearNumber := current_time_obj.Year() - 1970
		currentMonthNumber := int(current_time_obj.Month() - 1)
		monthsSinceEpoch := yearNumber*12 + currentMonthNumber

		monthsToRemoveFromCurrentTime := monthsSinceEpoch % int(klineInterval.Multiplier)

		openTime_date := time.Date(current_time_obj.Year(), current_time_obj.Month()-time.Month(monthsToRemoveFromCurrentTime), 1, 0, 0, 0, 0, time.UTC)
		closeTime_date := time.Date(current_time_obj.Year(), current_time_obj.Month()-time.Month(monthsToRemoveFromCurrentTime)+time.Month(klineInterval.Multiplier), 1, 0, 0, 0, 0, time.UTC)

		openTime = openTime_date.UnixMilli()
		closeTime = closeTime_date.UnixMilli() - 1

		return openTime, closeTime, nil
	case 'Y':
		yearNumber := current_time_obj.Year() - 1970

		yearsToRemoveFromCurrentTime := yearNumber - (yearNumber % int(klineInterval.Multiplier))

		openTime_date := time.Date(current_time_obj.Year()-yearsToRemoveFromCurrentTime, time.January, 1, 0, 0, 0, 0, time.UTC)
		closeTime_date := time.Date(current_time_obj.Year()-yearsToRemoveFromCurrentTime+int(klineInterval.Multiplier), time.January, 1, 0, 0, 0, 0, time.UTC)

		openTime = openTime_date.UnixMilli()
		closeTime = closeTime_date.UnixMilli() - 1

		return openTime, closeTime, nil
	}

	return 0, 0, fmt.Errorf("invalid interval rune of '%s' is invalid in '%s' is invalid", string(klineInterval.Rune), interval)
}

////

func (*Utils) GetStringNumberPrecision(numStr string) int {
	lastNumberIndex := 0
	dotIndex := 0

	dotFound := false

	for i, char := range numStr {
		if char == '.' {
			dotFound = true
			dotIndex = i
		} else if char != '0' {
			lastNumberIndex = i
		}
	}

	if !dotFound {
		dotIndex = len(numStr)
	}

	precision := lastNumberIndex - dotIndex

	if precision < 0 {
		precision++ // because if the number is right before the '.', then the precision must be 0, not -1 (so it's offset by 1)
	}

	return precision
}

func (*Utils) ParseInt(intStr string) (int64, error) {
	return strconv.ParseInt(intStr, 10, 64)
}

func (*Utils) ParseFloat(floatStr string) (float64, error) {
	precision := utils.GetStringNumberPrecision(floatStr)

	float, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		return float, err
	}

	return utils.ToFixed_Round(float, precision), nil
}

func (*Utils) DetectDotNumIndexes(numStr string) (dotIndex int, numIndex int) {
	dotIndex = -1
	numIndex = -1
	for i, char := range numStr {
		switch char {
		case '.':
			dotIndex = i
		case '0':
		default:
			numIndex = i
		}
	}

	return dotIndex, numIndex
}

func (*Utils) Round_priceStr(priceStr string, precision int) string {

	for i, char := range priceStr {
		if char != '0' {
			priceStr = priceStr[i:]
			break
		}
	}

	if precision == 0 {
		return strings.Split(priceStr, ".")[0]
	}

	if precision < 0 {
		abs_precision := -precision
		priceStr = strings.Split(priceStr, ".")[0]
		length := len(priceStr)
		endIndex := length - abs_precision

		if abs_precision >= length {
			return "0"
		}

		return priceStr[:endIndex] + strings.Repeat("0", abs_precision)
	} else {
		dotIndex, _ := utils.DetectDotNumIndexes(priceStr)
		if dotIndex == -1 {
			return priceStr + "." + strings.Repeat("0", precision)
		}

		arr := strings.Split(priceStr, ".")
		intStr, decimalStr := arr[0], arr[1]
		decimalLength := len(decimalStr)

		if decimalLength >= precision {
			decimalStr = decimalStr[:precision]
		} else {
			decimalStr += strings.Repeat("0", precision-decimalLength)
		}

		return intStr + "." + decimalStr
	}
}

func (*Utils) Format_TickSize_str(priceStr string, tickSize string) string {
	precision := utils.GetStringNumberPrecision(tickSize)

	return utils.Round_priceStr(priceStr, precision)
}

func (*Utils) ToFixed_Floor(price float64, precision int) float64 {
	return math.Floor(price*math.Pow10(precision)) / math.Pow10(precision)
}

func (*Utils) ToFixed_Round(price float64, precision int) float64 {
	return math.Round(price*math.Pow10(precision)) / math.Pow10(precision)
}

func (*Utils) ToFixed_Ceil(price float64, precision int) float64 {
	return math.Ceil(price*math.Pow10(precision)) / math.Pow10(precision)
}

func (*Utils) RemoveDuplicates(input []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(input))

	for _, str := range input {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}

// CreateQueryString transforms a map[string]interface{} into a query string
func (*Utils) CreateQueryString(params map[string]interface{}, sorted bool) string {
	if params == nil {
		return ""
	}

	// Extract keys to sort them if `sorted` is true
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}

	if sorted {
		sort.Strings(keys)
	}

	query := url.Values{}

	// Helper function to process values
	var addToQuery func(key string, value interface{})
	addToQuery = func(key string, value interface{}) {
		switch v := value.(type) {
		case string:
			query.Add(key, v)
		case []string:
			// Encode slices as JSON arrays
			jsonValue, err := json.Marshal(v)
			if err != nil {
				fmt.Printf("[VERBOSE] Error marshaling slice for key %s: %v\n", key, err)
				return
			}
			query.Add(key, string(jsonValue)) // Add JSON-encoded array
		case []interface{}:
			for _, item := range v {
				addToQuery(key, item) // Recursively handle each item
			}
		case map[string]interface{}:
			// Handle nested maps with dot notation
			for subKey, subValue := range v {
				addToQuery(key+"."+subKey, subValue)
			}
		case int, int64, float64, bool: // Convert basic types to string
			query.Add(key, fmt.Sprintf("%v", v))
		default:
			fmt.Printf("[VERBOSE] Error adding parameter: invalid type detected, received %v", v)
		}
	}

	// Process each key-value pair
	for _, key := range keys {
		addToQuery(key, params[key])
	}

	return query.Encode()
}
