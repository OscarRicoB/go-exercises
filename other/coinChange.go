package other

// A vending machine has the following denominations: 1c, 5c, 10c, 25c, 50c, and $1. Your task is to write a
// program that will be used in a vending machine to return change. Assume that the vending machine will always
//  want to return the least number of coins or notes. Devise a function CoinChange(M, P) where M is how much money
//  was inserted into the machine and P the price of the item selected, that returns an array of integers representing
//  the number of each denomination to return.
//
// 	fmt.Println(CoinChange(5, 0.99))		// should return [1,0,0,0,0,4]
// 	fmt.Println(CoinChange(3.14, 1.99))  // should return [0,1,1,0,0,1]
// 	fmt.Println(CoinChange(3, 0.01))     // should return [4,0,2,1,1,2]

// 	fmt.Println(CoinChange(4, 3.14))     // should return [1,0,1,1,1,0]
// 	fmt.Println(CoinChange(0.45, 0.34))  // should return [1,0,1,0,0,0]
//

func CoinChange(M float64, P float64) []int {
	change := make([]int, 6)
	coinValues := []int{1, 5, 10, 25, 50, 100}
	currentChange, currentCoin, money, price := 0, 5, int(M*100), int(P*100)
	c := money - price
	for currentChange < c {
		if currentChange+coinValues[currentCoin] <= c {
			currentChange += coinValues[currentCoin]
			change[currentCoin]++
		} else {
			currentCoin--
		}
		if currentCoin < 0 {
			break
		}
	}
	return change
}
