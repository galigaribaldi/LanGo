package conversions

var yardInKilometer float32
var yardInMetter float32
var yardInCentimeters float32

//Return yards in Kilometter
func GetYardKilometter() float32 {
	return yardInKilometer
}

//Return yards in metter
func GetYardMetter() float32 {
	return yardInMetter
}

//Return yards in Centimetter
func GetYardCentimetter() float32 {
	return yardInCentimeters
}

var featInKilometter float32
var featInMetter float32
var featInCentimetter float32

//Return Feat in Metter
func GetFeatInKilometter() float32 {
	return featInKilometter
}

//Return Feat in Metter
func GetFeatInMetter() float32 {
	return featInMetter
}

//Return Feat in Metter
func GetFeatInCentimetter() float32 {
	return featInCentimetter
}

var milleInKilometter float32
var milleInMetter float32
var milleInCentimetter float32

//Return Mille in KiloMetter
func GetMilleInKilometter() float32 {
	return milleInKilometter
}

//Return Mille in Metter
func GetMilleInMetter() float32 {
	return milleInMetter
}

//Return Mille in Metter
func GetMilleInCentiMetter() float32 {
	return milleInCentimetter
}

var incheInKilometter float32
var incheInMetter float32
var incheInCentimetter float32

//Return Inche in Metter
func GetIncheInKilometter() float32 {
	return incheInKilometter
}

//Return Inche in Metter
func GetIncheInMetter() float32 {
	return incheInMetter
}

//Return Inche in Metter
func GetIncheInCentimetter() float32 {
	return incheInCentimetter
}

//Variable Iniciatilizating with Default Values
func init() {
	yardInMetter = 0.9144
	yardInKilometer = 0.0009144
	yardInCentimeters = 91.44

	featInKilometter = 0.0003048
	featInMetter = 0.3048
	featInCentimetter = 30.48

	milleInMetter = 1609.34
	milleInKilometter = 1.60934
	milleInCentimetter = 160934

	incheInKilometter = 0.0000254
	incheInMetter = 0.0254
	incheInCentimetter = 2.54
}
