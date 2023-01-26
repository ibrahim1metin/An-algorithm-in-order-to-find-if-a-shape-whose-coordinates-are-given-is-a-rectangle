# An-algorithm-in-order-to-find-if-a-shape-whose-coordinates-are-given-is-a-rectangle
In this repository , there are two versions of the same algorithm. While one of them was written in Golang , the other one was written using Python. The algorithm works by following the following steps:  
1-First of all points should turn into a so called shape object.
2- Subsequently, the points need to be grouped since they might be given in a random order.
3- During the grouping, the lowest point should be deteceted firstly. In case there are more than one lowest point, the leftest one amoung these points should be detected.
4-After the detection of the lowest point , the highest point should be determined. If there are several highest points, the righest one should be found.
5-Afterwards, all points except the ones which are the highest or the lowest need to be placed in a list.
6-Then , using the coordinates of these points, the cosine values of the angles between the edges of the shape need to be calculated.
7- After that, these values also ought to be thrown in a list.
8- Since cos(90°) is 0, the boolean value of false is returned if there is value in the list other than 0 while the value true is returned  if no value other than zero is found.
