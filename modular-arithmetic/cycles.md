## Modular Arithmetic - Cycling
In this case we will be reviewing a cycle in one dimention. Let's suppose we have steps and a cycle grouped by n steps

Space 1D
S = defines the space's span
C = defines the cycle's steps

S = 24
C = 6

In this case we can fit `x` Cycles (C) into the Space (S) = 4

>_ _ _ _ _ _ | _ _ _ _ _ _ | _ _ _ _ _ _ | _ _ _ _ _ _ |
>     1             2             3             4


Cycle Properties:
- `independency`
- `stateless`

this means each cycle is independant from other cycles, altough it does not change a possible `sequence`

## Normalization
For a given `X` we can determinate 2 follow up things:
- c: how many cycles it covered up              | c = X // C
This is the integer division flooring, resulting into how many completed cycles has, meaning no floating point is ceiled

- r: current position on the uncovered cycle    | r = X % C
This is the remanent integer of the division, on the current uncovered cycle. Its a `relative` position. Can be though as a bandswitch from min limit and max limit such as floating points 0 to 1 but in integers before reaching 1 and complete a cycle
0 . . . . . 1

`X` can be think as either an element or position, however is encomendable to think of it as `n° element` that will have its own `position`. Both

### Example
S = 24
C = 6

X = 15

**Modular Arithmetic**
c = X // C
c = 15 // 6
c = 2

r = X % C
r = 15 % 6
r = 3


> | _ _ _ _ _ _ | _ _ _ _ _ _ | _ _ _ _ _ _ | _ _ _ _ _ _ |    24 steps
>        1             2             3             4
>
>        C             C            UC             UC
> | - - - - - - | - - - - - - | - - - _ _ _ | _ _ _ _ _ _ |    15 steps subspace (sS) from 24 steps space (S)
>   1 2 3 4 5 6   1 2 3 4 5 6   1 2 3 = 15
>                               |   |_ r = module = 3° element
>                               |
>                               b = 12° element

We can calculate the init position of the 3° cycle, with the variable b which represents the init position of the current uncovered cycle

b = C * c
b = 6 * 2
b = 12

**Summary**
S = Space = 24 steps
C = Cycle = 6 steps

X = 15° step position (couting from 1...) (note: sometimes is started from 0)

c = 2 completed cycles
r = 3° element position
b = 12° element position
