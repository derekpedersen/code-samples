using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace CodeFights.Arcade
{
    public class TheJourneyBegins
    {
        /// <summary>
        /// Write a function that returns the sum of two numbers.
        /// </summary>
        /// <param name="param1"></param>
        /// <param name="param2"></param>
        /// <returns></returns>
        int add(int param1, int param2)
        {
            return (param1 + param2);
        }

        /// <summary>
        /// Given a year, return the century it is in. The first century spans from the year 1 up to and including the year 100, the second - from the year 101 up to and including the year 200, etc.
        /// </summary>
        /// <param name="year"></param>
        /// <returns></returns>
        int centuryFromYear(int year)
        {

            var divided = year / 100;

            var multiplied = divided * 100;

            if (year > multiplied)
            {
                return divided + 1;
            }
            else
            {
                return divided;
            }
        }

        /// <summary>
        /// Given the string, check if it is a palindrome.
        /// </summary>
        /// <param name="inputString"></param>
        /// <returns></returns>
        bool checkPalindrome(string inputString)
        {
            var ogArray = inputString.ToCharArray();
            var revArray = inputString.ToCharArray();

            Array.Reverse(revArray);

            for (int i = 0; i < ogArray.Length; i++)
            {
                if (ogArray[i] != revArray[i])
                {
                    return false;
                }
            }

            return true;
        }

    }
}
