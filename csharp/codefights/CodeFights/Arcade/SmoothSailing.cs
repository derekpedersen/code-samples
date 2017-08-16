using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace CodeFights.Arcade
{
    public class SmoothSailing
    {
        /// <summary>
        /// Given an array of strings, return another array containing all of its longest strings.
        /// </summary>
        /// <param name="inputArray"></param>
        /// <returns></returns>
        string[] allLongestStrings(string[] inputArray)
        {
            List<string> inputs = inputArray.ToList();

            int longestLength = 0;

            foreach (string s in inputs)
            {
                if (s.Length > longestLength)
                {
                    longestLength = s.Length;
                }
            }

            List<string> output = new List<string>();

            foreach (string s in inputs)
            {
                if (s.Length == longestLength)
                {
                    output.Add(s);
                }
            }

            return output.ToArray();
        }

        /// <summary>
        /// Given two strings, find the number of common characters between them.
        /// </summary>
        /// <param name="s1"></param>
        /// <param name="s2"></param>
        /// <returns></returns>
        int commonCharacterCount(string s1, string s2)
        {
            List<string> uniqueCharacters = new List<string>();

            foreach (char c in s1)
            {
                if (!uniqueCharacters.Contains(c.ToString()))
                {
                    uniqueCharacters.Add(c.ToString());
                }
            }

            List<Tuple<string, int>> dict = new List<Tuple<string, int>>();

            foreach (string s in uniqueCharacters)
            {
                int occurenceCount = 0;

                foreach (char c in s1)
                {
                    if (c.ToString() == s)
                    {
                        occurenceCount++;
                    }
                }

                dict.Add(new Tuple<string, int>(s, occurenceCount));
            }

            int matchingCount = 0;

            foreach (var d in dict)
            {
                int previouslyOccuring = d.Item2;
                string s = d.Item1;

                int s2Occuring = 0;

                foreach (char c in s2)
                {
                    if (c.ToString() == s)
                    {
                        s2Occuring++;
                    }
                }

                if (previouslyOccuring > s2Occuring)
                {
                    matchingCount += s2Occuring;
                }
                else
                {
                    matchingCount += previouslyOccuring;
                }
            }

            return matchingCount;
        }

        /// <summary>
        /// Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.
        /// Given a ticket number n, determine if it's lucky or not.
        /// </summary>
        /// <param name="n"></param>
        /// <returns></returns>
        bool isLucky(int n)
        {
            string s = n.ToString();
            char[] charArray = s.ToCharArray();
            int charCount = charArray.Length;
            decimal halfwayChar = (charCount / 2);

            int firstHalf = 0;
            for (int i = 0; i < halfwayChar; i++)
            {
                int num = int.Parse(charArray[i].ToString());

                firstHalf += num;
            }

            int secondHalf = 0;
            for (int i = int.Parse(halfwayChar.ToString()); i < charArray.Length; i++)
            {
                int num = int.Parse(charArray[i].ToString());

                secondHalf += num;
            }

            return (firstHalf == secondHalf);
        }

    }
}
