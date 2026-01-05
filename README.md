# The Anarchy Solo Challenge Generator
I got the game for the holidays in 2025 and I really love it. I noticed that there hasn't been a solo challenge set up (maybe because it would require a lot of spoiler tags/information), so I decided to make a program that would generate everything.

# How it works
The `components.toml` file has a list of all the components that are necessary for generating the solo challenge. Since the domain cards are hidden (and not uniquely identifiable), I need to generate a complete deck. The below table shows how many possible domain cards need to exist. They probably wouldn't all be used, but this program generates them all.

102 cards - 24 cards per shuffle

| Area                     | Num of Cards |
| ------------------------ | ------------ |
| Knight's Training        | 18           |
| St. Valentine's Festival | 12           |
| Tournaments              | 36           |
| Brewhouse                | 12           |
| Michaelmas               | 12           |
| Lammas                   | 12           |

