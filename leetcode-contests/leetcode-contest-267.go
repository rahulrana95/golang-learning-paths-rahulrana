package leetcodeContest267

// There are n people in a line queuing to buy tickets, where the 0th person is at the front of the line and the (n - 1)th person is at the back of the line.

// You are given a 0-indexed integer array tickets of length n where the number of tickets that the ith person would like to buy is tickets[i].

// Each person takes exactly 1 second to buy a ticket. A person can only buy 1 ticket at a time and has to go back to the end of the line (which happens instantaneously) in order to buy more tickets. If a person does not have any tickets left to buy, the person will leave the line.

// Return the time taken for the person at position k (0-indexed) to finish buying tickets.
// Link: https://leetcode.com/contest/weekly-contest-267/problems/time-needed-to-buy-tickets/

func TimeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := 0; i < len(tickets); i++ {
		if i == k || tickets[i] < tickets[k] {
			ans += tickets[i]
		} else if i > k {
			if tickets[i] > tickets[k]-1 {
				ans += tickets[k] - 1
			} else {
				ans += tickets[i]
			}

		} else {
			if tickets[i] > tickets[k] {
				ans += tickets[k]
			} else {
				ans += tickets[i]
			}

		}
	}

	return ans
}

// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
func maxDistance(colors []int) int {
	ans := -1
	for i, _ := range colors {
		for j := i + 1; j < len(colors); j++ {
			if colors[j] != colors[i] {
				if j-i > ans {
					ans = j - i
				}
			}
		}
	}

	return ans

}
