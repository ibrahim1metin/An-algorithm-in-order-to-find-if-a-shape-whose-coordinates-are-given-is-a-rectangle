class Point:
  def __init__(self,x,y):
    self.x=x
    self.y=y
  def __str__(self):
    return f"({self.x},{self.y})"
  def __repr__(self):
    return str(self)
  def __add__(self,point):
    return Point(self.x+point.x,self.y+point.y)
  def __sub__(self,point):
    return Point(self.x-point.x,self.y-point.y)
  def pythogram(self):
    return (self.x**2+self.y**2)**0.5

class Shape:
  def __init__(self,*p):
    def check_if_point(obj):
      examplePoint=Point(0,0)
      if type(obj)==type(examplePoint):
        return obj
      else:
        raise ValueError("A shape can only consist of points")
    self.__points=[check_if_point(point) for point in p]
    self.xs=[i.x for i in self.__points]
    self.ys=[i.y for i in self.__points]
    return None
  def __call__(self):
    return self.__points
  def __str__(self):
    string_val_without_parantesis=""
    for point in self.__points:
      string_val_without_parantesis+=str(point)+","
    return f"[{string_val_without_parantesis[:-1:]}]"
  def getPoints(self):
    return self.__points
  def __repr__(self):
    return str(self)
  def getHighest(self):
    highPoints=list()
    maxy=max(self.ys)
    for point in self.__points:
      if point.y==maxy:
        highPoints.append(point)
    return highPoints if len(highPoints)!=1 else highPoints[0]
  def getLowest(self):
    lowPoints=list()
    miny=min(self.ys)
    for point in self.__points:
      if point.y==miny:
        lowPoints.append(point)
    return lowPoints if len(lowPoints)!=1 else lowPoints[0]
  def getLeftest(self):
    leftPoints=list()
    minx=min(self.xs)
    for point in self.__points:
      if point.x==minx:
        leftPoints.append(point)
    return leftPoints if len(leftPoints)!=1 else leftPoints[0]
  def getRightest(self):
    rightPoints=list()
    maxx=max(self.xs)
    for point in self.__points:
      if point.x==maxx:
        rightPoints.append(point)
    return rightPoints if len(rightPoints)!=1 else rightPoints[0]
  def group_points(self):
    lowest=self.getLowest()
    if type(lowest)==type([1]):
      lowest=Shape(*lowest)
      lowest=lowest.getLeftest()
    highest=self.getHighest()
    if type(highest)==type([0]):
      highest=Shape(*highest)
      highest=highest.getRightest()
    others=[point for point in self.getPoints() if not(point==lowest or point==highest)]
    return lowest, highest, others

def isRectangle(shape:Shape):
  low,high,others=shape.group_points()
  for i in range(len(others)):
    others[i]=others[i]-low
  def Cosine_of_the_angle(p1:Point,p2:Point)->float:
    return (p1.x*p2.x+p1.y*p2.y)/(p1.pythogram()*p2.pythogram())
  cosangles=[]
  cosangles.append(Cosine_of_the_angle(*others))
  for i in range(len(others)):
    others[i]=others[i]+low-high
  cosangles.append(Cosine_of_the_angle(*others))
  others[0]+=high
  others[1]+=high
  low-=others[0]
  high-=others[0]
  cosangles.append(Cosine_of_the_angle(low,high))
  low+=others[0]-others[1]
  high+=others[0]-others[1]
  cosangles.append(Cosine_of_the_angle(low,high))
  for cos in cosangles:
    if not cos==0:
      return False
  return True

p1= Point(0,2)
p2=Point(1,1)
p3=Point(3,3)
p4=Point(2,4)
ps=Shape(p1,p2,p3,p4)
isRectangle(ps)
