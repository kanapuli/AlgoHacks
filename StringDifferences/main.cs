using System;
using System.Linq;
using System.Collections.Generic;

class HackerRank
{
	static void Main(string[] args)
	{
		var a = Console.ReadLine();
		var b = Console.ReadLine();

		Console.WriteLine(a.Length);
		Console.WriteLine(b.Length);

		List<string> result = new List<string>();
		for (int i = a.Length; i > 0; i --)
		{


			if (b.Contains(String.Join("", a.Take(i))))
			{

				result.Add(String.Join("", a.Take(i)));
			}
		}
		if (result !=  null && result.Count > 0 )
		{
			Console.WriteLine(result.First());
		}
		var diff = result.Count > 0 ? result.First().Length : 0;
		Console.WriteLine( (b.Length - diff ) + (a.Length - diff ));
	}
}
