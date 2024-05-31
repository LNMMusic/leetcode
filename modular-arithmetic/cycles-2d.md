The following content presents an interesting concept of cycle or periodic patterns and functions that by themself are independant / stateless, but its sequence and contiguity leads to some non-standard values that can be normalized or re-scale into an arrangement of values of the statless cycle in order to fetch the positions and have an staless functions such as a linear one where we can fetch values or data as in a table

## Structure
### Constraints
- D (Dimentionality): 2d
- S (Space): the expansion of the space along the x and y axis

- Pattern: Zig-Zag
- Cycle: an stateless pattern with some init variables values that is repeated in a sequence periodically along the space S

#### Example
D = 2
S = 4 = 4x4 = 16 positions

C = 6 positions (↓ down and ↑diag)

space
```
[
    [-, -, -, -],
    [-, -, -, -],
    [-, -, -, -],
    [-, -, -, -],
]
```

pattern
```
[
    [↓, -, -, -], [1, -, -, -],
    [↓, -, ↑, -], [2, -, 6, -],
    [↓, ↑, -, -], [3, 5, -, -],
    [↓, -, -, -], [4, -, -, -],
]
```

Considering this C pattern, there is a clear `pivot` on (3, -) when it start raising diagonally towards the ceiling and reach the position (-, 3) where the cycle repeats

pattern
```
[
    [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -],
    [↓, -, ↑, -], [↓, -, ↑, -], [↓, -, ↑, -], [↓, -, ↑, -], [↓, -, ↑, -], [↓, -, ↑, -], [↓, -, ↑, -],
    [↓, ↑, -, -], [↓, ↑, -, -], [↓, ↑, -, -], [↓, ↑, -, -], [↓, ↑, -, -], [↓, ↑, -, -], [↓, ↑, -, -],
    [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -], [↓, -, -, -],
]
```

This cycle can be breakdown into 2 steps in an `stateless` manner
```
[       down     &    diag
    [↓, -, -, -], [-, -, -, -],
    [↓, -, -, -], [-, -, ↑, -],
    [↓, -, -, -], [-, ↑, -, -],
    [↓, -, -, -], [-, -, -, -],
]
```
meaning we can have a SYSTEM OF EQUATIONS with 2 `boundaries`. We can consider the input of our system of equation as the element and give its position along the cycle. Based on this stateless cycle, is NOT SUPPOSE to be beyond the limits
- N° Element: 1, 2, 3, 4, 5, 6 
- Positions : 0, 1, 2, 3, 4, 5

b = fixed init position of the cycle along cols (positions are counted from 0)
x = n° element (starting from 1 ...)
should be considered as position in order to fit the space S
x = n° position = x - 1


S normalization => now S will represent the limit of the S expansion into a position
S = S - 1
{
    -------------------------
    Down
    f(x) = [x, b]                       | IF x <= S

    -------------------------
    Diagonal
    This parts requires normalization, since we are working with a new pivot with linear values, and the first drop goes S elements and so the diagonal, we can reset the indexing back to 0 in the arrangement of 0,...,S

    Also since the diagonal always goes from bottom left to top right, it goes up, but since the array we do not count in an inverted manner of the cartesian graph, we should inverse the positions

    • Normalization
    x_n = (x - S)

    • Diagonal Inversion
    x_ni = (S - x_n)

    f(x_n) = [S-x_n, b+x_n]        | IF x > S
}

THIS IS STATELESS SYSTEM EQUATION WITH B INIT POSITION TO CREATE THE ZIG ZAG PATTERN FROM ANY B BIAS POINT

now since this pattern can be adjacent, we can create a sequence of repeated cycles giving a certain consistency in the position of X, so we might have to apply `modular-arithmetic` in order to normalize this position based on the consumed cycles and remanent steps


---

### Sequential Cycling - Modular Arithmetic

Given an X nº of element, we can try to figure out in which cycle falls, what's the remanent, and where is the exact position of that number of element along the pathway of the zigzag, in the 2D Space

#### Formula
##### Normalization
We need to first normalize the X n° of element in order to determinate the `relative position` around the uncovered cycle to work in an `stateless` manner with our previous system of equations and so for the `b` position

• c = X // C        | n° completed cycles
• r = X % C         | n° remanent steps along the incompleted cycle or n° element in the incompleted cycle

**Special Case**
In case X == C, it means a cycle is fully completed but no more remanent steps on an uncovered cycle. Meaning i should consider the last step of completed circle, relatively

if r == 0; r == C (in order to have the last )

now we have the exact number of cycles (c) and where we currently are along the uncovered cycle (r)

**r, x_p and b parameter**
now we need to calculate b to know the starting position of the current cycle. This will also affect r in an special case where there is no remanent.

In case of `r` we can normalize it into the x position called `x_p`
> {
>     x_p = C-1      | IF !r (no remanent)
>     x_p = r-1      | IF r  (remanent)
> }

In case of `b`
B defines where along the space, is the searched position, it depends on the number of cycles (c) and remanent (r) steps. Its a measurement of position in this case and based on the pattern, it will indicate the `axis x` or cols

Full Cycle Cols
> CS = 1 + (S-2)
> {
>   b = CS * (c-1)   | IF !r (no remanent)
>   b = CS * (c)     | IF r  (remanent)
> }

After we get c, r, and b we can send b, but x should be along the range or predispose limits of our stateless pattern function, so x is not literal x, but the relative position, that is represented in this case by r, the r that counts x from the beggining of the cycle

we send b, and finally r as x

b = position measure
r as x = position measure