/*
Copyright 2016 Mark Chenoweth
Copyright 2018 Alessandro Sanino
Licensed under terms of MIT license (see LICENSE)
*/

package talib

import (
	"testing"
	"github.com/kjx98/cgo-talib"
)

// Test all the functions

func TestCgoSma(t *testing.T) {
	result := Sma(testClose, 20)
	cgo_result := talib.Sma(testClose, 20)
	equals(t, result, cgo_result)
}

func TestCgoEma(t *testing.T) {
	result := Ema(testClose, 5)
	cgo_result := talib.Ema(testClose, 5)
	equals(t, result, cgo_result)
	result = Ema(testClose, 20)
	cgo_result = talib.Ema(testClose, 20)
	equals(t, result, cgo_result)
	result = Ema(testClose, 50)
	cgo_result = talib.Ema(testClose, 50)
	equals(t, result, cgo_result)
	result = Ema(testClose, 100)
	cgo_result = talib.Ema(testClose, 100)
	equals(t, result, cgo_result)
}

func TestCgoRsi(t *testing.T) {
	result := Rsi(testClose, 10)
	cgo_result := talib.Rsi(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoAdd(t *testing.T) {
	result := Add(testHigh, testLow)
	cgo_result := talib.Add(testHigh, testLow)
	equals(t, result, cgo_result)
}

func TestCgoDiv(t *testing.T) {
	result := Div(testHigh, testLow)
	cgo_result := talib.Div(testHigh, testLow)
	equals(t, result, cgo_result)
}

func TestCgoMax(t *testing.T) {
	result := Max(testClose, 10)
	cgo_result := talib.Max(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMaxIndex(t *testing.T) {
	result := MaxIndex(testClose, 10)
	cgo_result := talib.MaxIndex(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMin(t *testing.T) {
	result := Min(testClose, 10)
	cgo_result := talib.Min(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMinIndex(t *testing.T) {
	result := MinIndex(testClose, 10)
	cgo_result := talib.MinIndex(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMult(t *testing.T) {
	result := Mult(testHigh, testLow)
	cgo_result := talib.Mult(testHigh, testLow)
	equals(t, result, cgo_result)
}

func TestCgoSub(t *testing.T) {
	result := Sub(testHigh, testLow)
	cgo_result := talib.Sub(testHigh, testLow)
	equals(t, result, cgo_result)
}

func TestCgoRocp(t *testing.T) {
	result := Rocp(testClose, 10)
	cgo_result := talib.Rocp(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoObv(t *testing.T) {
	result := Obv(testClose, testVolume)
	cgo_result := talib.Obv(testClose, testVolume)
	equals(t, result, cgo_result)
}

func TestCgoAtr(t *testing.T) {
	result := Atr(testHigh, testLow, testClose, 14)
	cgo_result := talib.Atr(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoNatr(t *testing.T) {
	result := Natr(testHigh, testLow, testClose, 14)
	cgo_result := talib.Natr(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

/*
func TestCgoTRange(t *testing.T) {
	result := TRange(testHigh, testLow, testClose)
	cgo_result := talib.Trange(testHigh, testLow, testClose)
	equals(t, result, cgo_result)
}
*/

func TestCgoAvgPrice(t *testing.T) {
	result := AvgPrice(testOpen, testHigh, testLow, testClose)
	cgo_result := talib.AvgPrice(testOpen, testHigh, testLow, testClose)
	equals(t, result, cgo_result)
}

func TestCgoMedPrice(t *testing.T) {
	result := MedPrice(testHigh, testLow)
	cgo_result := talib.MedPrice(testHigh, testLow)
	equals(t, result, cgo_result)
}

func TestCgoTypPrice(t *testing.T) {
	result := TypPrice(testHigh, testLow, testClose)
	cgo_result := talib.TypPrice(testHigh, testLow, testClose)
	equals(t, result, cgo_result)
}

func TestCgoWclPrice(t *testing.T) {
	result := WclPrice(testHigh, testLow, testClose)
	cgo_result := talib.WclPrice(testHigh, testLow, testClose)
	equals(t, result, cgo_result)
}

/*
func TestCgoAcos(t *testing.T) {
	result := Acos(testRand)
	cgo_result := talib.Acos(testRand)
	equals(t, result, cgo_result)
}

func TestCgoAsin(t *testing.T) {
	result := Asin(testRand)
	cgo_result := talib.Asin(testRand)
	equals(t, result, cgo_result)
}

func TestCgoAtan(t *testing.T) {
	result := Atan(testRand)
	cgo_result := talib.Atan(testRand)
	equals(t, result, cgo_result)
}
*/

func TestCgoCeil(t *testing.T) {
	result := Ceil(testClose)
	cgo_result := talib.Ceil(testClose)
	equals(t, result, cgo_result)
}

/*
func TestCgoCos(t *testing.T) {
	result := Cos(testRand)
	cgo_result := talib.Cos(testRand)
	equals(t, result, cgo_result)
}

func TestCgoCosh(t *testing.T) {
	result := Cosh(testRand)
	cgo_result := talib.Cosh(testRand)
	equals(t, result, cgo_result)
}

func TestCgoExp(t *testing.T) {
	result := Exp(testRand)
	cgo_result := talib.Exp(testRand)
	equals(t, result, cgo_result)
}
*/

func TestCgoFloor(t *testing.T) {
	result := Floor(testClose)
	cgo_result := talib.Floor(testClose)
	equals(t, result, cgo_result)
}

/*
func TestCgoLn(t *testing.T) {
	result := Ln(testClose)
	cgo_result := talib.Ln(testClose)
	equals(t, result, cgo_result)
}

func TestCgoLog10(t *testing.T) {
	result := Log10(testClose)
	cgo_result := talib.Log10(testClose)
	equals(t, result, cgo_result)
}

func TestCgoSin(t *testing.T) {
	result := Sin(testRand)
	cgo_result := talib.Sin(testRand)
	equals(t, result, cgo_result)
}

func TestCgoSinh(t *testing.T) {
	result := Sinh(testRand)
	cgo_result := talib.Sinh(testRand)
	equals(t, result, cgo_result)
}
*/

func TestCgoSqrt(t *testing.T) {
	result := Sqrt(testClose)
	cgo_result := talib.Sqrt(testClose)
	equals(t, result, cgo_result)
}

/*
func TestCgoTan(t *testing.T) {
	result := Tan(testRand)
	cgo_result := talib.Tan(testRand)
	equals(t, result, cgo_result)
}

func TestCgoTanh(t *testing.T) {
	result := Tanh(testRand)
	cgo_result := talib.Tanh(testRand)
	equals(t, result, cgo_result)
}

func TestCgoSum(t *testing.T) {
	result := Sum(testClose, 10)
	cgo_result := talib.Sum(testClose, 10)
	equals(t, result, cgo_result)
}
*/

func TestCgoVar(t *testing.T) {
	result := Var(testClose, 10)
	cgo_result := talib.Var(testClose, 10, 1.0)
	equals(t, result, cgo_result)
}

/*
func TestCgoTsf(t *testing.T) {
	result := Tsf(testClose, 10)
	cgo_result := talib.Tsf(testClose, 10)
	equals(t, result, cgo_result)
}
*/

func TestCgoStdDev(t *testing.T) {
	result := StdDev(testRand, 10, 1.0)
	cgo_result := talib.StdDev(testRand, 10, 1.0)
	equals(t, result, cgo_result)
}

/*
func TestCgoLinearRegSlope(t *testing.T) {
	result := LinearRegSlope(testClose, 10)
	cgo_result := talib.LinearRegSlope(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoLinearRegIntercept(t *testing.T) {
	result := LinearRegIntercept(testClose, 10)
	cgo_result := talib.LinearRegIntercept(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoLinearRegAngle(t *testing.T) {
	result := LinearRegAngle(testClose, 10)
	cgo_result := talib.LinearRegAngle(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoLinearReg(t *testing.T) {
	result := LinearReg(testClose, 10)
	cgo_result := talib.LinearReg(testClose, 10)
	equals(t, result, cgo_result)
}
*/

func TestCgoCorrel(t *testing.T) {
	result := Correl(testHigh, testLow, 10)
	cgo_result := talib.Correl(testHigh, testLow, 10)
	equals(t, result, cgo_result)
}

func TestCgoBeta(t *testing.T) {
	result := Beta(testHigh, testLow, 5)
	cgo_result := talib.Beta(testHigh, testLow, 5)
	equals(t, result, cgo_result)
}

func TestCgoHtDcPeriod(t *testing.T) {
	result := HtDcPeriod(testClose)
	cgo_result := talib.HtDcPeriod(testClose)
	equals(t, result, cgo_result)
}

func TestCgoHtPhasor(t *testing.T) {
	result1, result2 := HtPhasor(testClose)
	cgo_result1, cgo_result2 := talib.HtPhasor(testClose)
	equals(t, result1, cgo_result1)
	equals(t, result2, cgo_result2)
}

func TestCgoHtSine(t *testing.T) {
	result1, result2 := HtSine(testClose)
	cgo_result1, cgo_result2 := talib.HtSine(testClose)
	equals(t, result1, cgo_result1)
	equals(t, result2, cgo_result2)
}

func TestCgoHtTrendline(t *testing.T) {
	result := HtTrendline(testClose)
	cgo_result := talib.HtTrendLine(testClose)
	equals(t, result, cgo_result)
}

func TestCgoHtTrendMode(t *testing.T) {
	result := HtTrendMode(testClose)
	cgo_result := talib.HtTrendMode(testClose)
	equals(t, result, cgo_result)
}

func TestCgoWillR(t *testing.T) {
	result := WillR(testHigh, testLow, testClose, 9)
	cgo_result := talib.Willr(testHigh, testLow, testClose, 9)
	equals(t, result, cgo_result)
}

func TestCgoAdx(t *testing.T) {
	result := Adx(testHigh, testLow, testClose, 14)
	cgo_result := talib.Adx(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoAdxR(t *testing.T) {
	result := AdxR(testHigh, testLow, testClose, 5)
	cgo_result := talib.Adxr(testHigh, testLow, testClose, 5)
	equals(t, result, cgo_result)
}

func TestCgoCci(t *testing.T) {
	result := Cci(testHigh, testLow, testClose, 14)
	cgo_result := talib.Cci(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoRoc(t *testing.T) {
	result := Roc(testClose, 10)
	cgo_result := talib.Roc(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoRocr(t *testing.T) {
	result := Rocr(testClose, 10)
	cgo_result := talib.Rocr(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoRocr100(t *testing.T) {
	result := Rocr100(testClose, 10)
	cgo_result := talib.Rocr100(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMom(t *testing.T) {
	result := Mom(testClose, 10)
	cgo_result := talib.Mom(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoBBands(t *testing.T) {
	upper, middle, lower := BBands(testClose, 5, 2.0, 2.0, SMA)
	cgo_upper, cgo_middle, cgo_lower := talib.BBands(testClose, 5, 2.0, 2.0, int32(SMA))
	equals(t, upper, cgo_upper)
	equals(t, middle, cgo_middle)
	equals(t, lower, cgo_lower)
}

func TestCgoDema(t *testing.T) {
	result := Dema(testClose, 10)
	cgo_result := talib.Dema(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoTema(t *testing.T) {
	result := Tema(testClose, 10)
	cgo_result := talib.Tema(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoWma(t *testing.T) {
	result := Wma(testClose, 10)
	cgo_result := talib.Wma(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMa(t *testing.T) {
	result := Ma(testClose, 10, DEMA)
	cgo_result := talib.Ma(testClose, 10, int32(DEMA))
	equals(t, result, cgo_result)
}

func TestCgoTrima(t *testing.T) {
	result := Trima(testClose, 10)
	cgo_result := talib.TriMa(testClose, 10)
	equals(t, result, cgo_result)
	result = Trima(testClose, 11)
	cgo_result = talib.TriMa(testClose, 11)
	equals(t, result, cgo_result)
}

func TestCgoMidPoint(t *testing.T) {
	result := MidPoint(testClose, 10)
	cgo_result := talib.MidPoint(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMidPrice(t *testing.T) {
	result := MidPrice(testHigh, testLow, 10)
	cgo_result := talib.MidPrice(testHigh, testLow, 10)
	equals(t, result, cgo_result)
}

func TestCgoT3(t *testing.T) {
	result := T3(testClose, 5, 0.7)
	cgo_result := talib.T3(testClose, 5, 0.7)
	equals(t, result, cgo_result)
}

func TestCgoKama(t *testing.T) {
	result := Kama(testClose, 10)
	cgo_result := talib.Kama(testClose, 10)
	equals(t, result, cgo_result)
}

func TestCgoMaVp(t *testing.T) {
	periods := make([]float64, len(testClose))
	for i := range testClose {
		periods[i] = 5.0
	}
	result := MaVp(testClose, periods, 2, 10, SMA)
	cgo_result := talib.Mavp(testClose, periods, 2, 10, int32(SMA))
	equals(t, result, cgo_result)
}

func TestCgoMinusDM(t *testing.T) {
	result := MinusDM(testHigh, testLow, 14)
	cgo_result := talib.MinusDm(testHigh, testLow, 14)
	equals(t, result, cgo_result)
}

func TestCgoPlusDM(t *testing.T) {
	result := PlusDM(testHigh, testLow, 14)
	cgo_result := talib.PlusDm(testHigh, testLow, 14)
	equals(t, result, cgo_result)
}

func TestCgoSar(t *testing.T) {
	result := Sar(testHigh, testLow, 0.0, 0.0)
	cgo_result := talib.Sar(testHigh, testLow, 0.0, 0.0)
	equals(t, result, cgo_result)
}

func TestCgoSarExt(t *testing.T) {
	result := SarExt(testHigh, testLow, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	cgo_result := talib.SarExt(testHigh, testLow, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	equals(t, result, cgo_result)
}

func TestCgoMama(t *testing.T) {
	mama, fama := Mama(testClose, 0.5, 0.05)
	cgo_mama, cgo_fama := talib.Mama(testClose, 0.5, 0.05)
	equals(t, mama, cgo_mama)
	equals(t, fama, cgo_fama)
}

func TestCgoMinMax(t *testing.T) {
	min, max := MinMax(testClose, 10)
	cgo_min, cgo_max := talib.MinMax(testClose, 10)
	equals(t, min, cgo_min)
	equals(t, max, cgo_max)
}

func TestCgoMinMaxIndex(t *testing.T) {
	minidx, maxidx := MinMaxIndex(testClose, 10)
	cgo_minidx, cgo_maxidx := talib.MinMaxIndex(testClose, 10)
	equals(t, minidx, cgo_minidx)
	equals(t, maxidx, cgo_maxidx)
}

func TestCgoApo(t *testing.T) {
	result := Apo(testClose, 12, 26, SMA)
	cgo_result := talib.Apo(testClose, 12, 26, int32(SMA))
	equals(t, result, cgo_result)
	result = Apo(testClose, 26, 12, SMA)
	cgo_result = talib.Apo(testClose, 26, 12, int32(SMA))
	equals(t, result, cgo_result)
}

func TestCgoPpo(t *testing.T) {
	result := Ppo(testClose, 12, 26, SMA)
	cgo_result := talib.Ppo(testClose, 12, 26, int32(SMA))
	equals(t, result, cgo_result)
	result = Ppo(testClose, 26, 12, SMA)
	cgo_result = talib.Ppo(testClose, 26, 12, int32(SMA))
	equals(t, result, cgo_result)
}

func TestCgoAroon(t *testing.T) {
	dn, up := Aroon(testHigh, testLow, 14)
	cgo_dn, cgo_up := talib.AroOn(testHigh, testLow, 14)
	equals(t, dn, cgo_dn)
	equals(t, up, cgo_up)
}

func TestCgoAroonOsc(t *testing.T) {
	result := AroonOsc(testHigh, testLow, 14)
	cgo_result := talib.AroOnOsc(testHigh, testLow, 14)
	equals(t, result, cgo_result)
}

func TestCgoBop(t *testing.T) {
	result := Bop(testOpen, testHigh, testLow, testClose)
	cgo_result := talib.Bop(testOpen, testHigh, testLow, testClose)
	equals(t, result, cgo_result)
}

func TestCgoCmo(t *testing.T) {
	result := Cmo(testClose, 14)
	cgo_result := talib.Cmo(testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoDx(t *testing.T) {
	result := Dx(testHigh, testLow, testClose, 14)
	cgo_result := talib.Dx(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoMinusDI(t *testing.T) {
	result := MinusDI(testHigh, testLow, testClose, 14)
	cgo_result := talib.MinusDi(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoPlusDI(t *testing.T) {
	result := PlusDI(testHigh, testLow, testClose, 14)
	cgo_result := talib.PlusDi(testHigh, testLow, testClose, 14)
	equals(t, result, cgo_result)
}

func TestCgoMfi(t *testing.T) {
	result := Mfi(testHigh, testLow, testClose, testVolume, 14)
	cgo_result := talib.Mfi(testHigh, testLow, testClose, testVolume, 14)
	equals(t, result, cgo_result)
}

func TestCgoUltOsc(t *testing.T) {
	result := UltOsc(testHigh, testLow, testClose, 7, 14, 28)
	cgo_result := talib.UltOsc(testHigh, testLow, testClose, 7, 14, 28)
	equals(t, result, cgo_result)
}

func TestCgoStoch(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 5, 3, SMA, 3, SMA)
	cgo_slowk, cgo_slowd := talib.Stoch(testHigh, testLow, testClose, 5, 3, int32(SMA), 3, int32(SMA))
	equals(t, slowk, cgo_slowk)
	equals(t, slowd, cgo_slowd)
}

func TestCgoStoch2(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 12, 3, SMA, 3, SMA)
	cgo_slowk, cgo_slowd := talib.Stoch(testHigh, testLow, testClose, 12, 3, int32(SMA), 3, int32(SMA))
	equals(t, slowk, cgo_slowk)
	equals(t, slowd, cgo_slowd)
}

func TestCgoStoch3(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 12, 3, SMA, 15, SMA)
	cgo_slowk, cgo_slowd := talib.Stoch(testHigh, testLow, testClose, 12, 3, int32(SMA), 15, int32(SMA))
	equals(t, slowk, cgo_slowk)
	equals(t, slowd, cgo_slowd)
}

func TestCgoStochF(t *testing.T) {
	fastk, fastd := StochF(testHigh, testLow, testClose, 5, 3, SMA)
	cgo_fastk, cgo_fastd := talib.Stochf(testHigh, testLow, testClose, 5, 3, int32(SMA))
	equals(t, fastk, cgo_fastk)
	equals(t, fastd, cgo_fastd)
}

func TestCgoStochRsi(t *testing.T) {
	fastk, fastd := StochRsi(testClose, 14, 5, 2, SMA)
	cgo_fastk, cgo_fastd := talib.StochRsi(testClose, 14, 5, 2, int32(SMA))
	equals(t, fastk, cgo_fastk)
	equals(t, fastd, cgo_fastd)
}

func TestCgoMacdExt(t *testing.T) {
	macd, macdsignal, macdhist := MacdExt(testClose, 12, SMA, 26, SMA, 9, SMA)
	cgo_macd, cgo_macdsignal, cgo_macdhist := talib.MacdExt(testClose, 12, int32(SMA), 26, int32(SMA), 9, int32(SMA))
	equals(t, macd, cgo_macd)
	equals(t, macdsignal, cgo_macdsignal)
	equals(t, macdhist, cgo_macdhist)
}

func TestCgoTrix(t *testing.T) {
	result := Trix(testClose, 5)
	cgo_result := talib.Trix(testClose, 5)
	equals(t, result, cgo_result)
	result = Trix(testClose, 30)
	cgo_result = talib.Trix(testClose, 30)
	equals(t, result, cgo_result)
}

func TestCgoMacd(t *testing.T) {
	macd, macdsignal, macdhist := Macd(testClose, 12, 26, 9)
	cgo_macd, cgo_macdsignal, cgo_macdhist := talib.Macd(testClose, 12, 26, 9)
	//unstable := 100
	equals(t, macd, cgo_macd)
	equals(t, macdsignal, cgo_macdsignal)
	equals(t, macdhist, cgo_macdhist)
}

func TestCgoMacdFix(t *testing.T) {
	macd, macdsignal, macdhist := MacdFix(testClose, 9)
	cgo_macd, cgo_macdsignal, cgo_macdhist := talib.MacdFix(testClose, 9)
	//unstable := 100
	equals(t, macd, cgo_macd)
	equals(t, macdsignal, cgo_macdsignal)
	equals(t, macdhist, cgo_macdhist)
}

func TestCgoAd(t *testing.T) {
	result := Ad(testHigh, testLow, testClose, testVolume)
	cgo_result := talib.Ad(testHigh, testLow, testClose, testVolume)
	equals(t, result, cgo_result)
}

func TestCgoAdOsc(t *testing.T) {
	result := AdOsc(testHigh, testLow, testClose, testVolume, 3, 10)
	cgo_result := talib.AdOsc(testHigh, testLow, testClose, testVolume, 3, 10)
	equals(t, result, cgo_result)
}
