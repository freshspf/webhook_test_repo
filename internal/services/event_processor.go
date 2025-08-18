package services

import (
	"fmt"
	"log"
	"math"
)

// PrimeProcessor 素数处理器
type PrimeProcessor struct {
	primes []int
}

// NewPrimeProcessor 创建新的素数处理器
func NewPrimeProcessor() *PrimeProcessor {
	return &PrimeProcessor{
		primes: make([]int, 0),
	}
}

// FindPrimesInRange 查找指定范围内的所有素数
func (pp *PrimeProcessor) FindPrimesInRange(start, end int) []int {
	log.Printf("开始查找 %d 到 %d 范围内的素数", start, end)
	
	primes := make([]int, 0)
	
	for num := start; num <= end; num++ {
		if pp.isPrime(num) {
			primes = append(primes, num)
		}
	}
	
	pp.primes = primes
	log.Printf("找到 %d 个素数: %v", len(primes), primes)
	return primes
}

// isPrime 判断一个数是否为素数
func (pp *PrimeProcessor) isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	// 只需要检查到 sqrt(n)
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GetPrimes 获取已找到的素数列表
func (pp *PrimeProcessor) GetPrimes() []int {
	return pp.primes
}

// PrintPrimes 打印素数列表
func (pp *PrimeProcessor) PrintPrimes() {
	if len(pp.primes) == 0 {
		fmt.Println("还没有找到任何素数")
		return
	}
	
	fmt.Printf("找到的素数有 %d 个:\n", len(pp.primes))
	for i, prime := range pp.primes {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%3d ", prime)
	}
	fmt.Println()
}

// GetPrimeCount 获取素数总数
func (pp *PrimeProcessor) GetPrimeCount() int {
	return len(pp.primes)
}

// IsInPrimes 检查一个数是否在已找到的素数列表中
func (pp *PrimeProcessor) IsInPrimes(num int) bool {
	for _, prime := range pp.primes {
		if prime == num {
			return true
		}
	}
	return false
}

// GetLargestPrime 获取最大的素数
func (pp *PrimeProcessor) GetLargestPrime() int {
	if len(pp.primes) == 0 {
		return -1
	}
	return pp.primes[len(pp.primes)-1]
}

// ProcessPrimesFrom1To100 处理1到100范围内的素数
func ProcessPrimesFrom1To100() {
	log.Println("=== 开始查找1-100范围内的素数 ===")
	
	processor := NewPrimeProcessor()
	primes := processor.FindPrimesInRange(1, 100)
	
	fmt.Println("\n=== 1-100范围内的素数 ===")
	processor.PrintPrimes()
	
	fmt.Printf("\n总计找到 %d 个素数\n", processor.GetPrimeCount())
	fmt.Printf("最大素数: %d\n", processor.GetLargestPrime())
	
	log.Println("=== 素数查找完成 ===")
}