package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	kubus1           Kubus1  = Kubus1{4}
	expectedVolume   float64 = 64
	expectedLuas     float64 = 96
	expectedKeliling float64 = 48
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus1.Volume(), expectedVolume)
	//t.Logf("volume: %.2f", kubus1.Volume())
	//if kubus1.Volume() != expectedVolume {
	//	t.Errorf("SALAH! expected %.2f", expectedVolume)
	//}
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus1.Luas(), expectedLuas)
	//t.Logf("luas: %.2f", kubus1.Luas())
	//if kubus1.Luas() != expectedLuas {
	//	t.Errorf("SALAH! expected %.2f", expectedLuas)
	//}
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus1.Keliling(), expectedKeliling)
	//t.Logf("keliling: %.2f", kubus1.Keliling())
	//if kubus1.Keliling() != expectedKeliling {
	//	t.Errorf("SALAH! expected %.2f", expectedKeliling)
	//}
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus1.Luas()
	}
}

/*
go test basic/kubus.go basic/kubus_test.go
ok      command-line-arguments  0.314s

go test basic/kubus.go basic/kubus_test.go -v
=== RUN   TestHitungVolume
    kubus_test.go:13: volume: 64.00
--- PASS: TestHitungVolume (0.00s)
=== RUN   TestHitungLuas
    kubus_test.go:20: luas: 96.00
--- PASS: TestHitungLuas (0.00s)
=== RUN   TestHitungKeliling
    kubus_test.go:27: keliling: 48.00
--- PASS: TestHitungKeliling (0.00s)
PASS
ok      command-line-arguments  0.080s

go test basic/kubus.go basic/kubus_test.go -bench=.
goos: darwin
goarch: arm64
BenchmarkHitungLuas-10          117888075                9.984 ns/op
PASS
ok      command-line-arguments  2.152s

go test basic/kubus.go basic/kubus_test.go -bench=. -v
=== RUN   TestHitungVolume
    kubus_test.go:13: volume: 64.00
--- PASS: TestHitungVolume (0.00s)
=== RUN   TestHitungLuas
    kubus_test.go:20: luas: 96.00
--- PASS: TestHitungLuas (0.00s)
=== RUN   TestHitungKeliling
    kubus_test.go:27: keliling: 48.00
--- PASS: TestHitungKeliling (0.00s)
goos: darwin
goarch: arm64
BenchmarkHitungLuas
BenchmarkHitungLuas-10          117942189                9.992 ns/op
PASS
ok      command-line-arguments  2.157s

117942189  9.992 ns/op  --> di test sebanyak 117942189 kali, membutuhkan waktu rata-rata 9.992 nano detik untuk run satu fungsi
*/
