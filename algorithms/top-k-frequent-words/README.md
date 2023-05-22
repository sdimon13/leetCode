# 692. Top K Frequent Words

Level: `Medium`

https://leetcode.com/problems/top-k-frequent-words/

---

# Description

Given an array of strings `words` and an integer `k`, return the `k` most frequent strings.

Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.

## Example 1:

    Input: words = ["i","love","leetcode","i","love","coding"], k = 2
    Output: ["i","love"]
    Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.

## Example 2:

    Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
    Output: ["the","is","sunny","day"]
    Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.

## Constraints:

- `1 <= words.length <= 500`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- `k` is in the range `[1, The number of unique words[i]]`

---

# Solution

## Intuition
This problem is asking us to find the top K frequent words from a list of words. Upon reading this problem, we can clearly identify that it is a frequency counting problem and the data structure that comes to our mind for this type of problem is Hashmap. But there is an additional constraint to this problem which is if two words have the same frequency, then the word with the lower alphabetical order comes first. To deal with this additional constraint, we also need a priority queue (Heap).

## Approach
Our approach to solving this problem is quite straightforward:

1. Create a HashMap and count the frequency of each word.
2. Create a Min Heap (Priority Queue) with elements as pairs of frequency and word.
3. Push elements from the HashMap to Heap with frequency as the priority. If the Heap size exceeds K, pop the top element.
4. When all words are processed, the heap contains the K most frequent words. Pop elements from the Heap and add to the result array.

This way, we will have the K most frequent words with higher frequency words coming first. If two words have the same frequency, the one with the lower alphabetical order will come first due to the way we've defined the comparison function in the Heap.

## Complexity
- Time complexity:
  The time complexity is O(N log K), where N is the length of the words array. We iterate over the words array once to create the HashMap, which takes O(N) time. Then, we push elements from the HashMap to the Heap. In the worst case, we end up pushing all the N words into the Heap, each operation taking log(K) time because K is the size of the Heap. So, the total time complexity becomes O(N log K).

- Space complexity:
  The space complexity is O(N) for the HashMap. In the worst case, all words are unique, so we need to store all of them in the HashMap.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)