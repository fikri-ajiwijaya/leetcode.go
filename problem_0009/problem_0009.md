# [Palindrome Number](https://leetcode.com/problems/palindrome-number)

Given an integer `x`, return `true` _if_ `x` _is a palindrome_, and `false` _otherwise_.

An integer is a **palindrome** when it reads the same forward and backward.

For example, $121$ is a palindrome while $123$ is not.

## Example 1

### Input

	x = 121

### Output

	true

### Explanation

$121$ reads as $121$ from left to right and from right to left.

## Example 2

### Input

	x = -121

### Output

	false

### Explanation

From left to right, it reads $-121$. From right to left, it becomes $121-$. Therefore it is not a palindrome.

## Example 3

### Input

	x = 10

### Output

	false

### Explanation

Reads $01$ from right to left. Therefore it is not a palindrome.

## Constraints

- $-2^{31} \leq$ `x` $\leq 2^{31} - 1$

### Follow up

Could you solve it without converting the integer to a string?
