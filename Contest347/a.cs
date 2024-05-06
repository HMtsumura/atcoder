int[] NK = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();
int N = NK[0];
int K = NK[1];
int [] a = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();
Console.Writeline(string.Join(" ", a.Where(x => x % K == 0).Select(y => y % K).ToArray()));