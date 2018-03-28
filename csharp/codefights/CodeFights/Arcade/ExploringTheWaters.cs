using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace CodeFights.Arcade
{
    public class ExploringTheWaters
    {
        public string[] addBorder(string[] picture)
        {
            var l = picture[0].Length + 2;
            var s = "";

            for (int i = 0; i < l; i++)
            {
                s += "*";
            }

            string[] sArray = new string[] { s };
            
            
        }
    }
}
