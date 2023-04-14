using ExtensionMethods;

namespace HelloWorld
{
    class Program
    {
        static void Main(string[] args)
        {
            string s = "Hello Extension Methods";
            int i = s.WordCount();
            Console.WriteLine("{0}", i);
        }
    }
}