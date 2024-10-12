To learn:
- https://blog.golang.org/laws-of-reflection
- https://quii.dev/The_Web_I_Want
- https://www.csszengarden.com/
- https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/
- https://hotwired.dev/
https://go.dev/blog/examples

https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/
【原始解法】牛在吃草，草在匀速生长。利用公式，设每头牛每天吃的草量为“1”，
10头牛20天吃草的量为10×20=200份，
15头牛吃10天吃草的量为15×10=150份，
每天生长出的草量为（200-150）÷（20-10）=5份，
牧场原有草200-5×20=100份，
25头牛每天吃草量与牧场每天生长出的草量之差25-5=20份，
可供25头牛吃100÷20=5天。
典型牛吃草问题，条件假设草生长速度固定不变，不同头数的牛吃光同一片草地所需的天数各不相同，求若干头牛吃这片草地可以吃多少天？
由于吃的天数不同，草又是天天在生长的，所以草的存量随牛吃的天数不断地变化。
解决牛吃草问题用四基本公式∶
（1）草的生长速度＝（对应的牛头数×吃的较多天数－相应的牛头数×吃的较少天数）÷（吃的较多天数－吃的较少天数）；
（2）原有草量＝牛头数×吃的天数－草的生长速度×吃的天数；
（3）吃的天数＝原有草量÷（牛头数－草的生长速度）；
（4）牛头数＝原有草量÷吃的天数＋草的生长速度。