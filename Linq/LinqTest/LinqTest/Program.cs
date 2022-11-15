using System;
using System.Linq;

namespace LinqTest
{
    public class LinqTest
    {
        public static void Main(string[] args)
        {
            string[] names = { "Peter", "Hartmut", "Jesika", "Ines" };

            var LinqName = from name in names
                           where name.Contains('t')
                           select name;

            foreach(var name in LinqName)
            {
                Console.WriteLine(name);
            }

            Console.ReadKey();
        }
    }
}
