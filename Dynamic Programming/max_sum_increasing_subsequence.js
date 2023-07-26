/*
	Write a function that takes in a non-empty array of integers and returns the greatest sum  that can be generated
	from a strictly-increasing subsequence in the array as well as an array of the numbers in that subsequence.

	Sample Input:  = [10, 70, 20, 30, 50, 11, 30]
	Output : [110, [10, 20, 30, 50]]
	
	Explanation:
	The given code snippet implements a function called `MaxSumIncreasingSubsequence`, which finds the maximum sum increasing subsequence in a given array of integers. An increasing subsequence is a sequence of array elements where each element is strictly greater than the previous element.

	Here's a step-by-step explanation of the code:

	1. The function `MaxSumIncreasingSubsequence` takes an input array of integers called `array`.

	2. Two arrays `sums` and `sequences` are initialized with the same length as the input array. The `sums` array stores the maximum sum of increasing subsequences ending at the corresponding index, and the `sequences` array stores the previous index that contributes to the maximum sum at the current index.

	3. The `maxSumIndex` variable is used to keep track of the index with the maximum sum of an increasing subsequence.

	4. The code uses a dynamic programming approach to calculate the maximum sum of increasing subsequences. It iterates through the input array from left to right and, for each element, checks all the previous elements to find the ones that are less than the current element and can form an increasing subsequence with it. If a better sum is found for the current element, it updates the `sums` and `sequences` arrays.

	5. After iterating through the entire array, the `maxSumIndex` stores the index with the maximum sum of an increasing subsequence.

	6. The function `buildSequence` is used to reconstruct the actual increasing subsequence from the `sequences` array, starting from the `maxSumIndex` and going backward until it reaches an element with a value of `math.MinInt32`, which is used as a sentinel value to indicate the end of the sequence.

	7. The `reverse` function is a helper function used to reverse the elements in the `sequence` array since the subsequence was built backward.

	8. The function returns the maximum sum of the increasing subsequence (`sum`) and the subsequence itself (`sequence`).

	O(n^2) time | O(n) space - where n is the length of the input array
*/

function maxSumIncreasingSubsequence(array) {
  const n = array.length;
  const sums = [...array]; // Store the maximum increasing sum up to index i.
  const sequences = new Array(n).fill(-1); // Store the previous index of the increasing subsequence.

  for (let i = 0; i < n; i++) {
    for (let j = 0; j < i; j++) {
      if (array[i] > array[j] && sums[j] + array[i] > sums[i]) {
        sums[i] = sums[j] + array[i];
        sequences[i] = j;
      }
    }
  }

  // Find the index of the maximum sum in 'sums'.
  let maxSumIndex = 0;
  for (let i = 1; i < n; i++) {
    if (sums[i] > sums[maxSumIndex]) {
      maxSumIndex = i;
    }
  }
  const maxSum = sums[maxSumIndex];

  // Build the increasing subsequence using the 'sequences' array.
  const sequence = [];
  while (maxSumIndex !== -1) {
    sequence.push(array[maxSumIndex]);
    maxSumIndex = sequences[maxSumIndex];
  }
  sequence.reverse();

  return [maxSum, sequence];
}
