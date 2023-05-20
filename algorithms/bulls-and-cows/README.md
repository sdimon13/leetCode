# 299. Bulls and Cows

Level: `Medium`

https://leetcode.com/problems/bulls-and-cows/

---

# Description

You are playing the [Bulls and Cows](https://en.wikipedia.org/wiki/Bulls_and_Cows) game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

- The number of "bulls", which are digits in the guess that are in the correct position.
- The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.

Given the secret number `secret` and your friend's guess `guess`, return the hint for your friend's guess.

The hint should be formatted as "`xAyB`", where `x` is the number of bulls and `y` is the number of cows. Note that both `secret` and `guess` may contain duplicate digits.

## Example 1:

    Input: secret = "1807", guess = "7810"
    Output: "1A3B"
    Explanation: Bulls are connected with a '|' and cows are underlined:
    "1807"
    |
    "7810"

## Example 2:

    Input: secret = "1123", guess = "0111"
    Output: "1A1B"
    Explanation: Bulls are connected with a '|' and cows are underlined:
    "1123"        "1123"
    |      or     |
    "0111"        "0111"
    Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.

## Constraints:

- `1 <= secret.length, guess.length <= 1000`
- `secret.length == guess.length`
- `secret` and `guess` consist of digits only.

---

# Solution

## Intuition
The problem is asking for the number of bulls (digits in guess that are in the correct position as in secret) and cows (digits in guess that are in the wrong position but exist in secret).

## Approach
To find the bulls is quite straightforward: just traverse both **secret** and **guess** from start to end and increment the counter for **bulls** every time we find **secret[i] == guess[i]**.

Finding the cows is trickier. We can use a single array of size 10 (since the digits can only be 0-9) to keep track of the difference in digit counts between **secret** and **guess**. We can increment the count of **cows** when we find a digit in **guess** that was a surplus in **secret** (which we kept track of in the array), and when we find a digit in **secret** that was a surplus in **guess**. This surplus is indicated by negative values for a digit from **secret** and positive values for a digit from **guess**.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the length of **secret** or **guess**. This is because we are doing a single pass over the input strings.

- Space complexity:
  The space complexity is O(1), since we only use an array of fixed size (10). This is not affected by the size of the input.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)