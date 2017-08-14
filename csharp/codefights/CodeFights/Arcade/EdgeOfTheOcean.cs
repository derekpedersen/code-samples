using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace CodeFights.Arcade
{
    public class EdgeOfTheOcean
    {
        /// <summary>
        /// Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.
        /// </summary>
        /// <param name="inputArray"></param>
        /// <returns></returns>
        int adjacentElementsProduct(int[] inputArray)
        {
            int highestProduct = 0;
            for (int i = 0; i < inputArray.Length; i++)
            {
                int farthestElement = (i + 1);

                if (!(farthestElement >= inputArray.Length))
                {
                    int productResult = inputArray[i] * inputArray[farthestElement];

                    if (highestProduct == 0 || productResult > highestProduct)
                    {
                        highestProduct = productResult;
                    }
                }
            }

            return highestProduct;
        }

        /// <summary>
        /// Below we will define an n-interesting polygon. Your task is to find the area of a polygon for a given n.
        /// A 1-interesting polygon is just a square with a side of length 1. An n-interesting polygon is obtained by taking the n - 1-interesting polygon and appending 1-interesting polygons to its rim, side by side. You can see the 1-, 2-, 3- and 4-interesting polygons in the picture below.
        /// </summary>
        /// <param name="n"></param>
        /// <returns></returns>
        int shapeArea(int n)
        {
            var first = n * n;
            var second = (n - 1) * (n - 1);
            return first + second;
        }

        /// <summary>
        /// Ratiorg got statues of different sizes as a present from CodeMaster for his birthday, each statue having an non-negative integer size. Since he likes to make things perfect, he wants to arrange them from smallest to largest so that each statue will be bigger than the previous one exactly by 1. He may need some additional statues to be able to accomplish that. Help him figure out the minimum number of additional statues needed.
        /// </summary>
        /// <param name="statues"></param>
        /// <returns></returns>
        int makeArrayConsecutive2(int[] statues)
        {
            Array.Sort(statues);

            var first = statues[0];
            var last = statues[statues.Length - 1];

            var difference = last - first;

            return difference - (statues.Length - 1);
        }

        /// <summary>
        /// Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.
        /// </summary>
        /// <param name="sequence"></param>
        /// <returns></returns>
        bool almostIncreasingSequence(int[] sequence)
        {
            bool foundOne = false;

            for (int i = -1, j = 0, k = 1; k < sequence.Length; k++)
            {
                bool deleteCurrent = false;
                if (sequence[j] >= sequence[k])
                {
                    if (foundOne)
                    {
                        return false;
                    }
                    foundOne = true;

                    if (k > 1 && sequence[i] >= sequence[k])
                    {
                        deleteCurrent = true;
                    }
                }

                if (!foundOne)
                {
                    i = j;
                }

                if (!deleteCurrent)
                {
                    j = k;
                }
            }

            return true;
        }

        /// <summary>
        /// After becoming famous, CodeBots decided to move to a new building and live together.The building is represented by a rectangular matrix of rooms, each cell containing an integer - the price of the room.Some rooms are free (their cost is 0), but that's probably because they are haunted, so all the bots are afraid of them. That is why any room that is free or is located anywhere below a free room in the same column is not considered suitable for the bots.
        /// </summary>
        /// <param name="matrix"></param>
        /// <returns></returns>
        int matrixElementsSum(int[][] matrix)
        {
            int totalCost = 0;
            List<int> zeroPositions = new List<int>();
            for (int i = 0; i < matrix.Length; i++)
            {
                // grab row array
                int[] row = matrix[i];
                for (int j = 0; j < row.Length; j++)
                {
                    if (row[j] == 0)
                    {
                        zeroPositions.Add(j);
                    }

                    if (!zeroPositions.Contains(j))
                    {
                        totalCost += row[j];
                    }
                }
            }

            return totalCost;
        }
    }
}
